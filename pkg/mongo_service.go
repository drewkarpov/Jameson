package pkg

import (
	"Jameson/config"
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
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

func (ms MongoImageService) CreateProject(project Project) interface{} {
	project.ID = GetNewId()
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := ms.ProjectsCollection.InsertOne(ctx, project)
	if err != nil {
		log.Fatal(err)
	}
	return result.InsertedID
}

func (ms MongoImageService) GetProjects() interface{} {
	var projects []Project
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, _ := ms.ProjectsCollection.Find(ctx, bson.M{})

	if cursor != nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var project Project
			cursor.Decode(&project)
			projects = append(projects, project)
		}
	}
	return projects
}

func (ms MongoImageService) GetContainers() interface{} {
	return ms.getTestContainersByFilter(bson.M{})
}
func (ms MongoImageService) GetContainerByName(name string) (*TestContainer, bool) {
	return ms.getContainerByDocument(bson.M{"name": name})
}

func (ms MongoImageService) GetContainerById(containerId string) (*TestContainer, bool) {
	return ms.getContainerByDocument(bson.M{"id": containerId})
}

func (ms MongoImageService) ApproveReferenceForContainer(containerId string) {
	ms.updateTestContainer(bson.M{"id": containerId}, bson.M{"$set": bson.M{"approved": true}})
}

func (ms MongoImageService) WritingTestResultToContainer(containerId string, test Test) {
	ms.updateTestContainer(bson.M{"id": containerId}, bson.M{"$push": bson.M{"tests": test}})
}

func (ms MongoImageService) CreateNewTestContainer(testContainer TestContainer) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	testContainer.ID = GetNewId()
	testContainer.Tests = []Test{}
	_, err := ms.ContainersCollection.InsertOne(ctx, testContainer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
func (ms MongoImageService) updateTestContainer(filter, changes bson.M) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	_, err := ms.ContainersCollection.UpdateOne(ctx, filter, changes)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (ms MongoImageService) getContainerByDocument(document bson.M) (*TestContainer, bool) {
	var containers = ms.getTestContainersByFilter(document)
	if containers == nil {
		return nil, false
	}
	return &containers[0], true
}

func (ms MongoImageService) getTestContainersByFilter(filter bson.M) []TestContainer {
	var containers []TestContainer
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, _ := ms.ContainersCollection.Find(ctx, filter)

	if cursor != nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var container TestContainer
			cursor.Decode(&container)
			containers = append(containers, container)
		}
	}
	return containers
}

func (ms MongoImageService) UploadImage(data []byte) string {
	filename := GetNewId()
	bucket, err := gridfs.NewBucket(
		ms.Database,
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename + ".png",
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)

	return filename
}
func (ms MongoImageService) DownloadImage(fileName string) []byte {

	db := ms.Database
	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File size to download: %v\n", dStream)
	return buf.Bytes()
}
