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

	// Or use promotion
	ad.notify()

	sendNotification(&ad.user)

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)
}
