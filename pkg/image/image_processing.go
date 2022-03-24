package image

import (
	"image"
	"image/color"
	"image/draw"
)

func GetImageDifference(imgBuff1, imgBuff2 []byte) ([]byte, float64, error) {
	img1, _ := getImageFromBytes(imgBuff1)
	img2, _ := getImageFromBytes(imgBuff2)

	if img1 == img2 {
		println()
	}

	b := img1.Bounds()

	resultImg := image.NewRGBA(image.Rect(
		b.Min.X,
		b.Min.Y,
		b.Max.X,
		b.Max.Y,
	))
	draw.Draw(resultImg, resultImg.Bounds(), img2, image.Point{0, 0}, draw.Src)

	accumError := int64(0)

	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()

			diff := int64(SqDiffUInt32(r1, r2))
			diff += int64(SqDiffUInt32(g1, g2))
			diff += int64(SqDiffUInt32(b1, b2))
			diff += int64(SqDiffUInt32(a1, a2))

			if diff > 0 {
				accumError += diff
				resultImg.Set(
					b.Min.X+x,
					b.Min.Y+y,
					color.RGBA{R: 255, A: 255})
			}
		}
	}

	nPixels := (b.Max.X - b.Min.X) * (b.Max.Y - b.Min.Y)
	percentage := float64(accumError/10) / (float64(nPixels) * 0xffff * 3)
	return getBytesFromRGBAImage(resultImg), percentage, nil
}

func SqDiffUInt32(x, y uint32) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}
