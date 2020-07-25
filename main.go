package main

import (
	"fmt"

	"github.com/DE-labtory/zulu/server"
)

func main() {
	s := server.New()
	if err := s.Run(":8080"); err != nil {
		panic(fmt.Sprintf("failed to run server: %s", err.Error()))
	}
}
