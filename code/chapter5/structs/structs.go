package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user // In this case, person is of type user. This is different to type embedding
	level  string
}

func structAssignment() {
	// Any time a variable is created and initialized to its zero value, it’s idiomatic to use
	// the keyword var. Reserve the use of the keyword var as a way to indicate that a variable
	// is being set to its zero value.
	var bill user
	fmt.Println(bill)

	// If the variable will be initialized to something other than its zero value, then use the
	// short variable declaration operator with a struct literal.

	// Fields can be in any order
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(lisa)

	// Declare a variable of type user.
	// Fields must be in same order as struct declaration
	burt := user{"Burt", "burt@email.com", 120, true}
	fmt.Println(burt)

	// Declare a variable of type admin
	// Using struct literals to create values for fields
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)
}

func typeAliases() {
	type Duration int64

	// var dur Duration

	// Compiler error - values of type int64 can’t be used as values of type Duration
	// dur = int64(1000)

	// This is the way to do it
	dur := Duration(int64(1000))
	fmt.Println(dur)
}

func main() {
	// structAssignment()
	typeAliases()
}
