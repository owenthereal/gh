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
	nameOwner := strings.Split(args.First(),"/")
	log.Println(nameOwner)

	if len(nameOwner) <= 2 && len(nameOwner) >= 1 {
		   gh := github.New()
		   if len(nameOwner) == 1 {
			nameOwner = append(nameOwner, nameOwner[0])
			nameOwner[0] = ""
		   }
		   url := gh.CloneURL(nameOwner[1], nameOwner[0], isSSH)
		   args.Remove(0)
		   args.Append(url)
	}
}

func parseClonePrivateFlag(args *Args) bool {
	if i := args.IndexOf("-p"); i == 0 {
		args.Remove(i)
		return true
	}

	return false
}
