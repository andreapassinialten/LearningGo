package events

import (
	"fmt"
	"sync"
)

type Action interface {
	Call(in string)
}

type Event struct {
	subscribers []Action
}

type EventManager struct {
	eventNumber int
	events      map[int](*Event)
}

var lock = &sync.Mutex{}

var singleInstance *EventManager

func getInstance() *EventManager {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &EventManager{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func (e *EventManager) Add(event *Event) int {
	(*e).events[(*e).eventNumber] = event

	key := (*e).eventNumber

	(*e).eventNumber += 1
	return key
}

func (e *EventManager) RemoveFromEvent(event *Event) {
	for key, value := range (*e).events { // Order not specified
		if value == event {
			delete((*e).events, key)
		}
	}
}

func (e *EventManager) RemoveFromKey(key int) {
	delete((*e).events, key)
}

func (e *EventManager) Invoke(event *Event, in string) {
	for key, value := range (*e).events { // Order not specified
		if value == event {
			for i := 0; i < len((*e).events[key].subscribers); i++ {
				(*e).events[key].subscribers[i].Call(in)
			}
		}
	}
}
