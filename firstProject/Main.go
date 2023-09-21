package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	toyota := Car{
		Name:   "To",
		Year:   123,
		Domain: "AAA",
		Engine: nil,
	}

	hp, err := toyota.EngineEv()

	if err != nil {
		log.Println("Error: " + err.Error())
		os.Exit(1)
	}

	fmt.Println(hp)

}

type Car struct {
	Name   string
	Year   uint8
	Domain string

	Engine *Engine
}

func (c Car) EngineEv() (uint16, error) {
	if c.Engine == nil {
		return 0, errors.New("No engine")
	}

	return c.Engine.HorsePower, nil
}

type Engine struct {
	Valves     uint8
	HorsePower uint16
}
