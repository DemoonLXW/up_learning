package main

import (
	"fmt"

	"github.com/DemoonLXW/up_learning/injection"
)

func main() {
	app, err := injection.ProvideApplication()
	if err != nil {
		fmt.Println("Application Setup Error...")
		return
	}
	app.Run(":8080")
}
