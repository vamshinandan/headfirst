package main

import "fmt"

// Switch Define type switch for on off operations
type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

// Toggleable Stores the operations on type switch
type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	s.toggle()
	s.toggle()
	// cannot assign a value to interface because it implements the pointer
	// value of switch type
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
