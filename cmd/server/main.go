package main

import (
	"fmt"
	"net/http"

	"github.com/manujo-varghese/go-rest-api/internal/article"
	"github.com/manujo-varghese/go-rest-api/internal/database"
	transportHTTP "github.com/manujo-varghese/go-rest-api/internal/transport/http"
)

// App - the struct which contai things l;ike pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error  {
	fmt.Println("Setting Up Our App")

	var err error
	db, err = database.NewDatabase()
	if err != nil{
		return err
	}
	err= database.MigrateDB(db)
	if err != nil{
		return err
	}

	articleService := article.NewService(db)

	handler := transportHTTP.NewHandler(articleService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil{
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
	
}

func main()  {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
		
	}

	
}