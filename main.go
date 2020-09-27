package main

import (
	"os"
	"pink/config"
)

func main() {
	// Example:
	// go run main.go development 3000
	env := os.Args[1]
	port := os.Args[2]
	config.RunApp(env, port)
}
