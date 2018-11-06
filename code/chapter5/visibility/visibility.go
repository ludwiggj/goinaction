package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/visibility/counters"
	"github.com/goinaction/code/chapter5/visibility/entities"
	"github.com/goinaction/code/chapter5/visibility/entitiez"
)

func doCounters() {
	// Compilation error: cannot refer to unexported name counters.alertCounter
	// undefined: counters.alertCounter
	// counter := counters.alertCounter(10)

	counter := counters.AlertCounter(10)
	fmt.Printf("Counter: %d\n", counter)

	// Create a variable of the unexported type using the exported New function from the package counters

	// the short variable declaration operator is capable of inferring the type and creating a variable
	// of the unexported type
	anotherCounter := counters.New(100)
	fmt.Printf("Counter: %d\n", anotherCounter)
}

func doEntities() {
	// Compilation error
	// unknown field 'email' in struct literal of type entities.User (but does have entities.email)

	// Can comment it out, but NOT REALLY SURE WHAT IT'S DONE
	u := entities.User{
		Name: "Bill",
		//email: "bill@email.com",
	}

	fmt.Println(u)
}

// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.

func doEntitiez() {
	// Create a value of type Admin from the entities package.
	a := entitiez.Admin{
		Rights: 1234,
	}

	// Set the exported fields from the unexported
	// inner type.
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}

func main() {
	doCounters()
	doEntities()
	doEntitiez()
}
