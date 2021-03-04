package main

import "fmt"

type Event struct {
	id   int
	desc string
}

func (e *Event) showInfo() {
	fmt.Println(e == nil)
	fmt.Println(e.id, e.desc)
}

func (e Event) showInfoNew() {
	fmt.Println(e.id, e.desc)
}

func main() {
	var e = Event{
		id:   1,
		desc: "first event"}
	e.showInfo()
	Event.showInfo(e)
	e.showInfoNew()
	Event.showInfoNew(e)
}
