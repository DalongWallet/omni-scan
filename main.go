package main

import (
	"github.com/DalongWallet/omni-scan/api/rest"
	"github.com/DalongWallet/omni-scan/conf"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func init() {
	conf.SetUp()
}

func main() {
	c := cli.NewCLI("omni-scan", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"RestApi":  RunRestApi,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}


func RunRestApi() (cli.Command, error) {
	return rest.New(), nil
}
