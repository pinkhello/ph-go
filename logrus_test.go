package phgo

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAppName(t *testing.T) {
	assert.Equal(t, appName, "-")
	SetAppName("hydra")
	assert.Equal(t, appName, "hydra")
}

func TestFileLog(t *testing.T) {
	path, err := ioutil.TempDir("./", "logs")
	defer os.RemoveAll(path)
	assert.Equal(t, err, nil)
	SetAppName("hydra")
	logger := &Logger{LogPath: path, LogType: "track"}
	logger.Init(nil)
	logger.Log.Warnf("%s", "233")
}
