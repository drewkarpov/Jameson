package pkg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/png"
	"log"
	"os"
)

type ImageWrapper struct {
	Reference ImageContainer
	Candidate ImageContainer
	Result    ImageContainer
}

type ImageContainer struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Body  image.Image        `json:"body" bson:"body.pix"`
	Error error              `json:"error" bson:"error"`
}

func (i *ImageWrapper) SetReference(filename string) {
	i.Reference = loadImage(filename)
	if i.Reference.Error != nil {
		panic(i.Reference.Error)
	}
}

func (i *ImageWrapper) SetCandidate(filename string) {
	i.Candidate = loadImage(filename)
	if i.Candidate.Error != nil {
		panic(i.Candidate.Error)
	}
}

func (i *ImageWrapper) CheckImagesProperties() {
	if i.Reference.Body.ColorModel() != i.Candidate.Body.ColorModel() {
		log.Fatal("different color models")
	}

	b := i.Reference.Body.Bounds()
	if !b.Eq(i.Candidate.Body.Bounds()) {
		log.Fatal("different image sizes")
	}
}

func (i *ImageWrapper) MustSaveImage(img image.Image, output string) {
	f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(f, img)
}

func loadImage(filename string) ImageContainer {
	f := mustOpen(filename)
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return ImageContainer{Body: img, Error: err}
}

func mustOpen(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return f
}
