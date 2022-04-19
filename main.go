package main

import (
	"github.com/mel-rob/echo-overview/router"
)

//main function
func main() {
	// create a new echo instance
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
