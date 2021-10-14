package ph_go

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEbv(t *testing.T) {
	value := GetEnv("k", "1")
	assert.Equal(t, value, "1")

	err := os.Setenv("k2", "2")
	assert.Equal(t, err, nil)
	value = GetEnv("k2", "")
	assert.Equal(t, value, "2")
}
