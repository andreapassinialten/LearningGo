package main

import (
	"Event/eventManger"
	"fmt"
)

func main() {
	manager := eventManger.GetInstance()

	var myEvent eventManger.Event
	manager.Add(&myEvent)

	var element Element
	myEvent.Subscribers = append(myEvent.Subscribers, element)

	manager.Invoke(&myEvent, "Carlo")

}

type Element struct {
}

func (elem Element) Call(in string) {
	fmt.Println(in)
}
