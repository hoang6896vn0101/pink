package main

import (
	"os"
	entry_point "pink/entry_point"
)

func main() {
	// Example:
	// go run main.go development 3000
	env := os.Args[1]
	port := os.Args[2]
	entry_point.RunApp(env, port)
}
