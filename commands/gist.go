package commands

var cmdGist = &Command{
	Run:   gist,
	Usage: "gist [-p] [-d] <FILE>",
	Short: "Create a gist on GitHub",
	Long: `Creates a gist on GitHub.
- It requires one or multiple files as the arguments.
- Use the flag -d to add description.
- Use the flag -p to create a private gist.
`,
}

func init() {
	CmdRunner.Use(cmdGist)
}

func gist(cmd *Command, args *Args) {
}
