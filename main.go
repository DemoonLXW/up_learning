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
	app.Run("127.0.0.1:8080")
}
