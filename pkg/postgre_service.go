package pkg

import (
	"fmt"
	"image"
)

type PostgreImageService struct {
}

func (ps PostgreImageService) GetReference(path string) image.Image {
	wrapper := ImageWrapper{}
	wrapper.SetReference("./images/ref1.png")
	wrapper.SetCandidate("./images/ref2.png")

	img1 := wrapper.Reference.Body
	img2 := wrapper.Candidate.Body

	wrapper.CheckImagesProperties()

	resultImg, percentage := GetImageDifference(img1, img2)
	wrapper.MustSaveImage(resultImg, "./images/result.png")

	fmt.Printf("Image difference: %f%%\n", percentage)

	return wrapper.Reference.Body
}
