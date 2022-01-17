package main

import "fmt"

// App - the struct which contai things l;ike pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error  {
	fmt.Println("Setting Up Our App")
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