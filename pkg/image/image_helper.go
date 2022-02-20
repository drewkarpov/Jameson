package image

import (
	"bytes"
	"image"
	"image/png"
	"log"
)

func getImageFromBytes(bts []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(bts))
	if err != nil {
		log.Fatalln(err)
	}
	return img, err
}

func getBytesFromRGBAImage(img *image.RGBA) []byte {
	buff := new(bytes.Buffer)
	err := png.Encode(buff, img)

	if err != nil {
		log.Fatalln(err)
	}
	return buff.Bytes()
}
