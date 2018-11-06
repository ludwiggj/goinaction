package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver

// method operates against a copy of the value used to make the method call
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
	u.name = u.name + " (notified)"
	fmt.Printf("Notified: %v\n", u)
}

// changeEmail implements a method with a pointer receiver

// In this case the value used to make the call is shared with the method
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	fmt.Printf("Bill: %v\n", bill)

	bill.notify()
	fmt.Printf("Bill: %v\n", bill)

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.

	// To support the method call, Go adjusts the pointer value to comply with the
	// method’s receiver. You can imagine that Go is performing the following operation:

	// (*lisa).notify()

	// Once again, notify is operating against a copy, but this time a copy of
	// the value that the lisa pointer points to.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.

	// Go adjusts the value to comply with the method’s receiver to support the call
	// (&bill).notify()
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	fmt.Printf("Lisa: %v\n", lisa)

	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()

	fmt.Printf("Lisa: %v\n", lisa)
}
