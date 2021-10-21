package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/singfield/rest-api-golang/internal/transport/http"
)

// contains things like the
// pointers to database connections
type  App struct{}

// startup of app
func ( app * App) Run() error {
	fmt.Println("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8000", handler.Router); err !=nil{
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}


func main(){
	app := App{}
	if err := app.Run(); err !=nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}