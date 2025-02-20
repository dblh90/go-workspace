package main

import "fmt"

type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (ds *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", ds.Name, "got data change:", data)
}
