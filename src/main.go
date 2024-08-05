package main

import (
	"fmt"

	"github.com/game-craft/src/engine"
)

func main() {
	engine := engine.NewEngine()
	fmt.Println("running!")
	err := engine.Run()
	if err != nil {
		fmt.Println("error!")
	}
	fmt.Print("bye")
}
