package commands

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestTransformCloneArgs(t *testing.T) {
	args := NewArgs([]string{"jingweno/gh"})
	
	transformCloneArgs(args)
	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git://github.com/jingweno/gh.git", args.First())

	args = NewArgs([]string{"-p", "jingweno/gh"})

	transformCloneArgs(args)
	assert.Equal(t, 1, args.Size())
	assert.Equal(t, "git@github.com:jingweno/gh.git", args.First())
}
