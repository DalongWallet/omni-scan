package rest

import "github.com/mitchellh/cli"

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	apiServer := NewHttpServer(20517)
	apiServer.Run()
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

const synopsis = "Interact with the omni-scan rest api service"
const help = `
Usage: omni-scan rest <subcommand> [options] [args]

  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.
  
`

