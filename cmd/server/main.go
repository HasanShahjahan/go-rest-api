package main

import "fmt"

//App - the struct which contains things like pointers
//to database connections
type App struct {
}

//Run - sets up application
func (a *App) Run() error {
	fmt.Println("Setting up APP")
	return nil
}

func main() {
	fmt.Println("GO REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
