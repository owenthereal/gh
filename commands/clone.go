package commands

import (
	"log"
	"strings"
	"github.com/jingweno/gh/github"
)

var cmdClone = &Command{
	Run:          clone,
	GitExtension: true,
	Usage:        "clone [-p] [USER/]REPOSITORY",
	Short:        "Clone a GitHub repository",
}

/**
  $ gh clone jingweno/gh
  > git clone git://github.com/jingweno/gh.git 

  $ gh clone -p jingweno/gh
  > git clone git@github.com:jingweno/gh.git

  $ gh clone gh
  > git clone git@github.com:YOUR_LOGIN/gh.git
**/

func clone(command *Command, args *Args) {
	isSSH := parseClonePrivateFlag(args)
	ownerName := strings.Split(args.First(),"/")

	if len(ownerName) <= 2 && len(ownerName) >= 1 {
		   if len(ownerName) == 1 {
			isSSH = true
			ownerName = append(ownerName, ownerName[0])
			ownerName[0] = ""
		   }
		   owner, name := ownerName[0], ownerName[1]
		   url := cloneUrl(name, owner, isSSH)
		   args.Remove(0)
		   args.Append(url)
	}
}

func cloneUrl(name, owner string, isSSH bool) string {
	gh := github.NewBlank()
	return gh.CloneURL(name, owner, isSSH)
}

func parseClonePrivateFlag(args *Args) bool {
	if i := args.IndexOf("-p"); i == 0 {
		args.Remove(i)
		return true
	}

	return false
}
