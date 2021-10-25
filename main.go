package main

import (
	"Jameson/pkg"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const mainDirectory = "./images"

func main() {
	if _, err := os.Stat(mainDirectory); os.IsNotExist(err) {
		err := os.Mkdir(mainDirectory, os.ModeDir)
		fmt.Errorf("cannot create folder\n error : %s", err)
	}

	mongos := pkg.MongoImageService{}.Init()

	data, err := ioutil.ReadFile("./images/ref1.png")
	if err != nil {
		log.Fatal(err)
	}
	conn := mongos.Database
	bucket, err := gridfs.NewBucket(
		conn,
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	uploadStream, err := bucket.OpenUploadStream("ref1.png")
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

	DownloadFile(*mongos.Database, "ref1.png")

	wrapper := pkg.ImageWrapper{}
	wrapper.SetReference("./images/ref1.png")
	wrapper.SetCandidate("./images/ref2.png")

	shutdown := make(chan error, 1)

	router := gin.Default()
	imgHandler := pkg.ImageHandler{pkg.PostgreImageService{}}
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.StaticFS("./images", http.Dir("images"))

	api := router.Group("/api/v1")
	{
		api.GET("/path/:path", imgHandler.GetOriginImage)
		api.POST("/upload", imgHandler.Upload)
	}

	err = http.ListenAndServe(":3333", router)
	shutdown <- err

}

func DownloadFile(db mongo.Database, fileName string) {

	fsFiles := db.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	// you can print out the results
	fmt.Println(results)

	bucket, _ := gridfs.NewBucket(
		&db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File size to download: %v\n", dStream)
	ioutil.WriteFile(fileName, buf.Bytes(), 0600)

}
