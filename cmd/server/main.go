package main

import (
	"fmt"
	"net/http"

	"github.com/singfield/rest-api-golang/internal/comment"
	"github.com/singfield/rest-api-golang/internal/database"
	transportHTTP "github.com/singfield/rest-api-golang/internal/transport/http"
)

// contains things like the
// pointers to database connections
type App struct{}

// startup of app
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err !=nil {
		return err
	}
	
	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8000", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}
