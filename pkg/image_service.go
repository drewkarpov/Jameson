package pkg

import "image"

type ImageService interface {
	GetReference(path string) image.Image
}
