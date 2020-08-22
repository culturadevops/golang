package main

/*
Ejemplo map
En este ejemplo intento mostrar
trabajar estructuras con maps y maps con estructuras
*/
import (
	"fmt"
)

type Stats struct {
	cnt      int
	category map[string]Events
}

type Events struct {
	cnt   int
	event map[string]*Event
}

type Event struct {
	value int64
}

func main() {

	stats := new(Stats)
	stats.cnt = 33
	stats.category = make(map[string]Events)
	e, f := stats.category["aa"]
	if !f {
		e = Events{}
	}
	e.cnt = 66

	e.event = make(map[string]*Event)
	stats.category["aa"] = e
	stats.category["aa"].event["bb"] = &Event{}
	stats.category["aa"].event["bb"].value = 99

	fmt.Println(stats)
	fmt.Println(stats.cnt, stats.category["aa"].event["bb"].value)
}
