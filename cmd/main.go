package main

import (
	"log"

	"github.com/donipit99/todo-app"
	"github.com/donipit99/todo-app/pkg/handler"
	"github.com/donipit99/todo-app/pkg/repository"
	"github.com/donipit99/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("0.0.0.0:8000", handlers.InitRoutes()); err != nil { // error occured while running http server: listen tcp: address ::8000: too many colons in address

		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

// 2025/03/14 02:13:53 error occured while running http server:
// listen tcp: address 8000: missing port in address exit status 1

// [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// - using env:   export GIN_MODE=release
//  - using code:  gin.SetMode(gin.ReleaseMode) //
