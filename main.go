package main

import (
	"fmt"
	"pink/infrastructure/configs"
)

func main() {
	fmt.Print(configs.GetConfig("development"))
}
