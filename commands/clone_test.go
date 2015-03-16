package commands

import (
	"github.com/bmizerany/assert"
	"github.com/jingweno/gh/git"
	"github.com/jingweno/gh/github"
	"os"
	"testing"
)

func TestTransformCloneArgs(t *testing.T) {
	os.Setenv("GH_PROTOCOL", "git")
	github.CreateTestConfigs("jingweno", "123")

	args := NewArgs([]string{"clone", "foo/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git://github.com/foo/gh.git", args.FirstParam())

	args = NewArgs([]string{"clone", "-p", "foo/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:foo/gh.git", args.FirstParam())

	args = NewArgs([]string{"clone", "jingweno/gh"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git://github.com/jingweno/gh.git", args.FirstParam())

	args = NewArgs([]string{"clone", "-p", "acl-services/devise-acl"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:acl-services/devise-acl.git", args.FirstParam())

	args = NewArgs([]string{"clone", "jekyll_and_hyde"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:jingweno/jekyll_and_hyde.git", args.FirstParam())

	args = NewArgs([]string{"clone", "-p", "jekyll_and_hyde"})
	transformCloneArgs(args)

	assert.Equal(t, 1, args.ParamsSize())
	assert.Equal(t, "git@github.com:jingweno/jekyll_and_hyde.git", args.FirstParam())

	args = NewArgs([]string{"clone", "git://github.com/jingweno/gh", "gh"})
	transformCloneArgs(args)

	assert.Equal(t, 2, args.ParamsSize())
	assert.Equal(t, "git://github.com/jingweno/gh", args.FirstParam())
	assert.Equal(t, "gh", args.GetParam(1))
}

func TestSaveHost(t *testing.T) {
	defer git.UnsetConfig(github.GhHostConfig)

	args := NewArgs([]string{"clone", "git://github.com/jingweno/gh"})
	saveHost(args)
	v, _ := git.Config(github.GhHostConfig)
	assert.Equal(t, "github.com", v)
	git.UnsetConfig(github.GhHostConfig)

	args = NewArgs([]string{"clone", "http://github.com/jingweno/gh"})
	saveHost(args)
	v, _ = git.Config(github.GhHostConfig)
	assert.Equal(t, "github.com", v)
	git.UnsetConfig(github.GhHostConfig)

	args = NewArgs([]string{"clone", "http://github.enterprise.com/jingweno/gh"})
	saveHost(args)
	v, _ = git.Config(github.GhHostConfig)
	assert.Equal(t, "github.enterprise.com", v)
	git.UnsetConfig(github.GhHostConfig)

	args = NewArgs([]string{"clone", "https://github.enterprise.com/jingweno/gh"})
	saveHost(args)
	v, _ = git.Config(github.GhHostConfig)
	assert.Equal(t, "github.enterprise.com", v)
	git.UnsetConfig(github.GhHostConfig)

	args = NewArgs([]string{"clone", "git@github.enterprise.com:jingweno/jekyll_and_hyde.git"})
	saveHost(args)
	v, _ = git.Config(github.GhHostConfig)
	assert.Equal(t, "github.enterprise.com", v)
	git.UnsetConfig(github.GhHostConfig)
}
