package main

import (
	"fmt"
	"net/http"

	"github.com/NeerChayaphon/go-rest-api/internal/database"
	"github.com/NeerChayaphon/go-rest-api/internal/todo"
	transportHTTP "github.com/NeerChayaphon/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	database.MigrateDB(db)
	if err != nil {
		return err
	}

	todoService := todo.NewService(db)

	handler := transportHTTP.NewHandler(todoService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
