package main

import (
	"omni-scan/api/rest"
)

func main() {
	rest.NewHttpServer(80)
}
