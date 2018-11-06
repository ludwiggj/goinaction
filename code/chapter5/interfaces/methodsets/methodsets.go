package main

import "fmt"

// notifier is an interface that defined notification type behavior
type notifier interface {
	notify()
}

type greeter interface {
	greet()
}

type user struct {
	name  string
	email string
}

// Methods Receivers  | Values
// -----------------------------------------------
// (t T)              | T and *T
// (t *T)             | *T

// If you implement an interface using a pointer receiver, then only pointers of that type implement the interface.
// If you implement an interface using a value receiver, then both values and pointers of that type implement the
// interface.

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

// greet implements a method with a value receiver.
func (u user) greet() {
	fmt.Printf("Hello to user %s\n", u.name)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

func sayHello(g greeter) {
	g.greet()
}

// duration is a type with a base type of int.
type duration int

// format pretty-prints the duration value.
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	// Following doesn't work
	// cannot use u (type user) as type notifier in argument to sendNotification:
	// user does not implement notifier (notify method has pointer receiver)

	// sendNotification(u)

	// BUT, as per methods, can call method directly on value
	// Behund the scenes, go is doing (&u).notify()
	u.notify()

	// Though you can't invoke it directly in this way

	// Compilation error: invalid indirect of u (type user)
	// (*u).notify()

	// Can do it by declaring a new variable
	uRef := &u
	uRef.notify()

	// Whereas this does
	sendNotification(&u)

	sayHello(u)

	sayHello(&u)

	// cannot call pointer method on duration(42)
	// cannot take the address of duration(42)
	// duration(42).pretty()

	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
