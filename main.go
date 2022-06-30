package main

import (
	"fmt"
	"github.com/xxarupakaxx/deepl-prompt/utils"
)

const (
	version = "1.0.0"
)

func main() {
	p := utils.NewPrompt()

	fmt.Printf("deepl-prompt %s\n", version)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	p.Run()
}
