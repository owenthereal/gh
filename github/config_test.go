package github

import (
	"github.com/bmizerany/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveGitHubConfig(t *testing.T) {
	config := NewConfig("jingweno", "123")
	file := "./test_support/test"
	defer os.RemoveAll(filepath.Dir(file))

	err := saveTo(file, &config)
	assert.Equal(t, nil, err)

	config, err = loadFrom(file)
	assert.Equal(t, nil, err)
	assert.Equal(t, "jingweno", config.User)
	assert.Equal(t, "123", config.Token)
	assert.Equal(t, "https://api.github.com", config.Url)
}

func TestSaveGitHubEnterpriseConfig(t *testing.T) {
	testGitHubEnterpriseUrl(t, "github.corporate.com", "https://github.corporate.com")
}

func TestSaveGitHubEnterpriseHttpConfig(t *testing.T) {
	testGitHubEnterpriseUrl(t, "http://github.corporate.com", "http://github.corporate.com")
}

func testGitHubEnterpriseUrl(t *testing.T, savedUrl, expectedUrl string) {
	config := NewConfigWithUrl("foo", "456", savedUrl)
	file := "./test_support/test"
	err := saveTo(file, &config)
	assert.Equal(t, nil, err)

	config, err = loadFrom(file)
	assert.Equal(t, "foo", config.User)
	assert.Equal(t, "456", config.Token)
	assert.Equal(t, expectedUrl, config.Url)
}
