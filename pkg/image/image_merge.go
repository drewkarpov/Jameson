package image

import (
	"bytes"
	"image"
	"image/draw"
)

type SecondGrid struct {
	ImageBytes []byte
	VoidZones  []VoidZone
}

type MergeImage struct {
	Grid         SecondGrid
	ImageCountDX int
	ImageCountDY int
}

func New(grid SecondGrid, imageCountDX, imageCountDY int) *MergeImage {
	mi := &MergeImage{
		Grid:         grid,
		ImageCountDX: imageCountDX,
		ImageCountDY: imageCountDY,
	}

	return mi
}

func (m *MergeImage) readGridImage() (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(m.Grid.ImageBytes))
	if err != nil {

	}
	return img, nil
}

func (m *MergeImage) mergeGrids(img image.Image) (*image.RGBA, error) {
	var canvas *image.RGBA

	canvasMaxPoint := image.Point{X: img.Bounds().Dx(), Y: img.Bounds().Dy()}
	canvasRect := image.Rectangle{Min: image.Point{}, Max: canvasMaxPoint}
	canvas = image.NewRGBA(canvasRect)

	draw.Draw(canvas, canvasRect, img, image.Point{}, draw.Src)

	for _, voidZone := range m.Grid.VoidZones {
		img, err := voidZone.GetVoidZoneRectangleImage()
		if err != nil {
			return nil, err
		}

		gridRect := canvasRect.Bounds().Add(image.Point{X: voidZone.OffsetX, Y: voidZone.OffsetY})
		draw.Draw(canvas, gridRect, img, image.Point{}, draw.Over)
	}

	return canvas, nil
}

func (m *MergeImage) Merge() (*image.RGBA, error) {
	img, err := m.readGridImage()
	if err != nil {
		return nil, err
	}
	return m.mergeGrids(img)
}

func GetImageWithVoidZones(imageBytes []byte, voidZones []VoidZone) []byte {
	grid := SecondGrid{
		ImageBytes: imageBytes,
		VoidZones:  voidZones,
	}
	rgba, err := New(grid, 1, 1).Merge()

	if err != nil {

	}
	return getBytesFromRGBAImage(rgba)
}
