package pkg

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type MongoImageService struct {
	Database           *mongo.Database
	ProjectsCollection *mongo.Collection
}

func InitMongoService() MongoImageService {
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	var cred = options.Credential{Username: "mongoadmin", Password: "mongoadmin"}
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://89.223.26.208:27017").SetAuth(cred))
	database := client.Database("jameson")
	project := database.Collection("projects")
	return MongoImageService{Database: database, ProjectsCollection: project}
}

func (ms MongoImageService) CreateProject(name string) interface{} {
	project := Project{Name: name}
	project.ID = primitive.NewObjectID()
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
			var person Project
			cursor.Decode(&person)
			projects = append(projects, person)
		}
	}
	return projects
}

func (ms MongoImageService) AddTestContainerToProject(id string, testContainer TestContainer) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	testContainer.ID = primitive.NewObjectID()
	objID, errr := primitive.ObjectIDFromHex(id)
	if errr != nil {
		panic(errr)
	}
	_, err := ms.ProjectsCollection.UpdateOne(
		ctx,
		bson.M{"id": objID},
		bson.M{"$push": bson.M{"containers": testContainer}},
	)
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
		filename,
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
