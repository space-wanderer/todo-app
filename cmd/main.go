package main

import (
	"log"

	"github.com/space-wanderer/todo-app"
	"github.com/space-wanderer/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while eunning http server: %s", err.Error())
	}
}
