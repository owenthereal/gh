package git

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestGitRemotes(t *testing.T) {
	output := []string{
		"origin	https://github.com/jingweno/gh (fetch)",
		"origin	https://github.com/jingweno/gh (push)"}

	remotes, err := gitRemotes(output)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, remotes)
}

func TestGitRemotesEnterprise(t *testing.T) {
	output := []string{
		"origin	https://github.enterprise.com/jingweno/gh (fetch)",
		"origin	https://github.enterprise.com/jingweno/gh (push)"}

	remotes, err := gitRemotes(output)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, remotes)
}
