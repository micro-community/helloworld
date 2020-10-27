package main

import (
	"github.com/micro-community/helloworld/handler"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	_ "github.com/micro-community/helloworld/profile"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Register handler
	srv.Handle(new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
