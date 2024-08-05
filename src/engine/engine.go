package engine

import (
	"bufio"
	"fmt"
	"os"
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
		rdr := bufio.NewReader(os.Stdin)
		c, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("error!")
		}

		e.running = false
		fmt.Print(c)
	}

	return nil
}
