package main

import (
	"github.com/mitchellh/cli"
	"log"
	"omni-scan/api/rest"
	"omni-scan/scan"
	"os"
)

func main() {
	c := cli.NewCLI("omni-scan", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"ScanData":   ScanData,
		"RestApi":    RunRestApi,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}


func ScanData() (cli.Command, error) {
 	return scan.New(), nil
}

func RunRestApi() (cli.Command, error) {
	return rest.New(), nil
}