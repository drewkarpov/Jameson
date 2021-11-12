package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestSqDiffUInt32(t *testing.T) {

	tests := []struct {
		name          string
		X             uint32
		Y             uint32
		expectedValue uint64
	}{
		{name: "positive", X: 11, Y: 1, expectedValue: 100},
		{name: "empty value for diff", X: 11, Y: 11, expectedValue: 0},
		{name: "not empty for diff, y > x", X: 0, Y: 11, expectedValue: 121},
		{name: "not empty for diff, x > y", X: 11, Y: 0, expectedValue: 121},
		{name: "x and y - zero values", X: 0, Y: 0, expectedValue: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := SqDiffUInt32(test.X, test.Y)
			assert.Equal(t, test.expectedValue, v)
		})
	}
}

func TestGetImageDifference(t *testing.T) {

	tests := []struct {
		name               string
		candidate          string
		reference          string
		expectedPercentage float64
		err                error
		resultBuffLength   int
	}{
		{
			name:               "images have diff",
			candidate:          "./testdata/ref1.png",
			reference:          "./testdata/ref2.png",
			expectedPercentage: 11365.422755233396,
			resultBuffLength:   11601,
		},
		{
			name:               "images no have diff",
			candidate:          "./testdata/ref1.png",
			reference:          "./testdata/ref1.png",
			expectedPercentage: 0,
			resultBuffLength:   11550,
		},
		{
			name:               "images difference by bounds",
			candidate:          "./testdata/ref1.png",
			reference:          "./testdata/ref3.png",
			expectedPercentage: 0,
			resultBuffLength:   0,
			err:                errors.New("img1 bounds is not equal img 2 bounds"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			img1Buff := loadImageBytes(test.candidate)
			img2Buff := loadImageBytes(test.reference)

			diff, actualPercentage, err := GetImageDifference(img1Buff, img2Buff)

			assert.Equal(t, test.resultBuffLength, len(diff), " length of bytes from images should be equal")
			assert.Equal(t, test.expectedPercentage, actualPercentage, "actualPercentage check")
			assert.Equal(t, test.err, err, "check error")
		})
	}
}

func loadImageBytes(filename string) []byte {
	buff, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot open file with path %s \nerror: %v", filename, err)
		return []byte{}
	}
	return buff
}
