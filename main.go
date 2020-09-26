package main

import (
	"fmt"
	configs "pink/infa/configs"
)

func main() {
	fmt.Print(configs.GetConfig("development"))
}
