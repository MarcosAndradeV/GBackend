package main

import (
	"fmt"
	"gbackend/internal/config"
)

func main() {
	ac := config.NewAppConfig(".env.json")
	fmt.Printf("ac: %v\n", ac)
}
