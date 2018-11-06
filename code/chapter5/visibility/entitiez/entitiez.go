package entitiez

type user struct {
	Name  string
	Email string
}

// Admin defines an admin in the program.
type Admin struct {
	user   // The embedded type is unexported.
	Rights int
}
