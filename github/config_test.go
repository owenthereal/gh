package github

import (
	"github.com/bmizerany/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveConfig(t *testing.T) {
	config := NewConfig("jingweno", "123")
	file := "./test_support/test"
	defer os.RemoveAll(filepath.Dir(file))

	err := saveTo(file, &config)
	assert.Equal(t, nil, err)

	config, err = loadFrom(file)
	assert.Equal(t, nil, err)
	assert.Equal(t, "jingweno", config.User)
	assert.Equal(t, "123", config.Token)
	assert.Equal(t, "", config.Host)
	assert.Equal(t, "github.com", config.FetchHost())

	newConfig := NewConfigWithUrl("foo", "456", "github.corporate.com")
	err = saveTo(file, &newConfig)
	assert.Equal(t, nil, err)

	config, err = loadFrom(file)
	assert.Equal(t, "foo", config.User)
	assert.Equal(t, "456", config.Token)
	assert.Equal(t, "github.corporate.com", config.Host)
}
