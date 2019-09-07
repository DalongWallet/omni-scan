package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"omni-scan/api/rest"
)

type program struct{

}

func main() {
	rest.NewHttpServer(80)
}

func (program) Init(env svc.Environment) error {
	fmt.Println("init...")
	return nil
}
