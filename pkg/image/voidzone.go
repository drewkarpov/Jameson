package image

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

type VoidZone struct {
	Width   int `json:"width"`
	Height  int `json:"height"`
	OffsetX int `json:"offset_x"`
	OffsetY int `json:"offset_y"`
}

func (v *VoidZone) GetVoidZoneRectangleImage() (image.Image, error) {
	upLeft := image.Point{}
	lowRight := image.Point{X: v.Width, Y: v.Height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	for x := 0; x < v.Width; x++ {
		for y := 0; y < v.Height; y++ {
			img.Set(x, y, cyan)
		}
	}
	buff := bytes.NewBuffer([]byte{})
	err := png.Encode(buff, img)

	if err != nil {

	}
	return getImageFromBytes(buff.Bytes())
}
