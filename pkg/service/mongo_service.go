//nolint:govet
package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/drewkarpov/Jameson/pkg/image"
	"log"
	"time"

	"github.com/drewkarpov/Jameson/config"
	mdl "github.com/drewkarpov/Jameson/pkg/model"
	"github.com/drewkarpov/Jameson/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoImageService struct {
	Database             *mongo.Database
	ProjectsCollection   *mongo.Collection
	ContainersCollection *mongo.Collection
}

func InitMongoService(config config.Config) MongoImageService {
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	var cred = options.Credential{Username: config.Database.Username, Password: config.Database.Password}
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+config.Database.Host+":27017").SetAuth(cred))
	database := client.Database(config.Database.DbName)
	project := database.Collection("projects")
	containers := database.Collection("containers")
	return MongoImageService{Database: database, ProjectsCollection: project, ContainersCollection: containers}
}

func (ms MongoImageService) CreateProject(project *mdl.Project) (*mdl.Project, error) {
	project.ID = utils.GetNewId()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err := ms.ProjectsCollection.InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (ms MongoImageService) GetProjects() ([]mdl.Project, error) {
	var projects []mdl.Project
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, _ := ms.ProjectsCollection.Find(ctx, bson.M{})

	var err error
	if cursor != nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var project mdl.Project
			err = cursor.Decode(&project)
			if err != nil {
				return nil, err
			}
			projects = append(projects, project)
		}
	}
	return projects, err
}

func (ms MongoImageService) GetContainers() []mdl.TestContainer {
	return ms.getTestContainersByFilter(bson.M{})
}
func (ms MongoImageService) GetContainerByName(name string) (*mdl.TestContainer, bool) {
	return ms.getContainerByDocument(bson.M{"name": name})
}

func (ms MongoImageService) GetContainerById(containerId string) (*mdl.TestContainer, bool) {
	return ms.getContainerByDocument(bson.M{"id": containerId})
}

func (ms MongoImageService) GetContainerByTestId(testId string) (*mdl.TestContainer, bool) {
	return ms.getContainerByDocument(bson.M{"tests.id": testId})
}

func (ms MongoImageService) ApproveReferenceForContainer(containerId string) (bool, error) {
	_, isExists := ms.GetContainerById(containerId)
	if !isExists {
		return false, nil
	}
	return ms.updateTestContainer(bson.M{"id": containerId}, bson.M{"$set": bson.M{"approved": true}})
}

func (ms MongoImageService) SetNewReferenceForContainer(containerId string, reference mdl.Reference) (bool, error) {
	_, isExists := ms.GetContainerById(containerId)
	if !isExists {
		return false, nil
	}
	return ms.updateTestContainer(bson.M{"id": containerId}, bson.M{"$set": bson.M{"reference_id": reference.ID}})
}

func (ms MongoImageService) AddVoidZonesForReference(containerId string, zones []image.VoidZone) error {
	_, isExists := ms.GetContainerById(containerId)
	if !isExists {
		return errors.New(fmt.Sprintf("cannot find container with id %s", containerId))
	}
	isSuccess, err := ms.updateTestContainer(bson.M{"id": containerId}, bson.M{"$set": bson.M{"void_zones": zones}})
	if err != nil || !isSuccess {
		return err
	}
	return nil
}

func (ms MongoImageService) WritingTestResultToContainer(candidate, result []byte, percentage float64, containerId, referenceId string) (*mdl.Test, error) {
	candidateId, err := ms.UploadImage(candidate)
	if err != nil {
		return nil, err
	}
	resultId, err := ms.UploadImage(result)
	if err != nil {
		return nil, err
	}

	_, isExists := ms.GetContainerById(containerId)
	if !isExists {
		return nil, err
	}

	test := mdl.Test{ID: utils.GetNewId(), CandidateId: *candidateId, ReferenceId: referenceId,
		Result: mdl.TestResult{ID: *resultId, Percentage: percentage}}

	isSuccess, err := ms.updateTestContainer(bson.M{"id": containerId}, bson.M{"$push": bson.M{"tests": test}})
	if err != nil || !isSuccess {
		return nil, err
	}
	return &test, nil
}

func (ms MongoImageService) CreateNewTestContainer(testContainer mdl.TestContainer) (*mdl.TestContainer, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	testContainer.ID = utils.GetNewId()
	testContainer.Tests = []mdl.Test{}
	_, err := ms.ContainersCollection.InsertOne(ctx, testContainer)
	if err != nil {
		return nil, err
	}
	return &testContainer, nil
}

func (ms MongoImageService) DeleteContainerById(containerId string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := ms.ContainersCollection.DeleteOne(ctx, bson.M{"id": containerId})
	if err != nil {
		return false, err
	}
	switch result.DeletedCount {
	case 1:
		return true, err
	default:
		return false, err
	}
}
func (ms MongoImageService) updateTestContainer(filter, changes bson.M) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err := ms.ContainersCollection.UpdateOne(ctx, filter, changes)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ms MongoImageService) getContainerByDocument(document bson.M) (*mdl.TestContainer, bool) {
	var containers = ms.getTestContainersByFilter(document)
	if containers == nil {
		return nil, false
	}
	return &containers[0], true
}

func (ms MongoImageService) getTestContainersByFilter(filter bson.M) []mdl.TestContainer {
	var containers []mdl.TestContainer
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, _ := ms.ContainersCollection.Find(ctx, filter)

	if cursor != nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var container mdl.TestContainer
			err := cursor.Decode(&container)
			if err != nil {
				continue
			}
			containers = append(containers, container)
		}
	}
	return containers
}

func (ms MongoImageService) UploadImage(data []byte) (*string, error) {
	filename := utils.GetNewId()
	bucket, err := gridfs.NewBucket(
		ms.Database,
	)
	if err != nil {
		return nil, err
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename + ".png",
	)
	if err != nil {
		return nil, err
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		return nil, err
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)

	return &filename, nil
}
func (ms MongoImageService) DownloadImage(fileName string) ([]byte, error) {
	db := ms.Database
	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)

	if err != nil {
		return nil, err
	}
	bucket, _ := gridfs.NewBucket(db)

	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)

	if err != nil {
		return nil, err
	}
	fmt.Printf("File size to download: %v\n", dStream)
	return buf.Bytes(), nil
}
