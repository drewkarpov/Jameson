package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewId(t *testing.T) {
	str := GetNewId()

	assert.NotEmpty(t, str, "id should be is not empty")
	assert.Equal(t, 32, len(str), "id should have uuid len of symbols")
}
