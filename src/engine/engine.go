package engine

import (
	"fmt"
	"log"
)

type Engine interface {
	Run() error
}

type engine struct {
	running bool
}

func NewEngine() Engine {
	return &engine{}
}

func (e engine) isRunning() bool {
	return e.running
}

func (e *engine) Run() error {
	e.running = true

	for e.isRunning() {
		var c string
		fmt.Scanln(&c)

		if c == "quit()" {
			e.running = false
		}

		log.Println(c)
		// fmt.Println(c)
	}

	return nil
}
