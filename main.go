package main

import (
	"github.com/labstack/echo"
)

//main function
func main() {
	// create a new echo instance
	e := echo.New()
	e.Logger.Fatal(e.Start(":8000"))
}
