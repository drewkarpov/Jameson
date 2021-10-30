package pkg

import (
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

func InitMongoService() MongoImageService {
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	var cred = options.Credential{Username: "mongoadmin", Password: "mongoadmin"}
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://89.223.26.208:27017").SetAuth(cred))
	database := client.Database("jameson")
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
	var containers []TestContainer
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	cursor, _ := ms.ContainersCollection.Find(ctx, bson.M{})

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

func (ms MongoImageService) GetTestContainer() interface{} {
	var test TestContainer
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	ms.ProjectsCollection.FindOne(ctx, bson.M{"containers.name": "some_test"}).Decode(test)

	return test
}

func (ms MongoImageService) WritingTestContainer(testContainer TestContainer) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	testContainer.ID = GetNewId()
	_, err := ms.ContainersCollection.InsertOne(ctx, testContainer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (ms MongoImageService) UploadImage(data []byte, filename string) {

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
