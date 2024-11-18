package main

import (
	"fmt"
	"main/main/internal/api"
)

func main() {
	config, err := api.LoadConfig()
	if err != nil {
		panic(err)
	}

	engine := config.NewAPI()

	if err = engine.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Alive and Ready to serve traffic!")
}
