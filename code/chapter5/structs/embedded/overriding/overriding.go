package main

import (
	"fmt"
)

// a user
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // Embedded Type
	level string
}

// New bit!
// notify implements a method that can be called via a value of type admin.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

// Next step - let's add an interface
type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Access inner type directly
	ad.user.notify()

	// Promotion no longer used
	ad.notify()

	sendNotification(&ad.user)

	// Promotion no longer used
	sendNotification(&ad)
}
