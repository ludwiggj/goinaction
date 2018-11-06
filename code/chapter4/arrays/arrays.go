package main

import "fmt"

func main() {
	// Declare an integer array of five elements.
	var array1 [5]int
	fmt.Println(array1)

	// Declare an integer array of five elements.
	// Initialize each element with a specific value.
	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2)

	// Change the value at index 2.
	array2[2] = 35
	fmt.Println(array2)

	// Declare an integer array.
	// Initialize each element with a specific value.
	// Capacity is determined based on the number of values initialized.
	array3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array3)

	// Declare an integer array of five elements.
	// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	array4 := [5]int{1: 10, 2: 20}
	fmt.Println(array4)

	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	array5 := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*array5[0] = 10
	*array5[1] = 20
	fmt.Println(array5)

	// Declare a string array of five elements.
	var array6 [5]string
	fmt.Println(array6)

	// Declare a second string array of five elements.
	// Initialize the array with colors.
	array7 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(array7)

	// Copy the values from array2 into array1.
	array6 = array7
	fmt.Println(array6)

	// Declare a string array of four elements.
	// var array8 [4]string

	// Declare a second string array of five elements.
	// Initialize the array with colors.
	// array9 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// Copy the values from array2 into array1.
	// Compiler Error: cannot use array9 (type [5]string) as type [4]string in assignment
	// array8 = array9

	// Declare a string pointer array of three elements.
	var array10 [3]*string

	// panic: runtime error: invalid memory address or nil pointer dereference
	// *array10[0] = "Black"

	// cannot take address of "Black"
	// array10[0] = &("Black")

	black := "Black"
	array10[0] = &black

	fmt.Println(array10)

	// Declare a second string pointer array of three elements.
	// Initialize the array with string pointers.
	array11 := [3]*string{new(string), new(string), new(string)}

	// Add colors to each element
	*array11[0] = "Red"
	*array11[1] = "Blue"
	*array11[2] = "Green"

	// Copy the values from array11 into array10.
	array10 = array11

	for _, p := range array10 {
		fmt.Printf("%v\n", *p)
	}

	for _, p := range array11 {
		fmt.Printf("%v ", *p)
	}

	fmt.Println()

	// Declare a two dimensional integer array of four elements
	// by two elements.
	var array12 [4][2]int
	fmt.Println(array12)

	// Use an array literal to declare and initialize a two
	// dimensional integer array.
	array13 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array13)

	// Declare and initialize index 1 and 3 of the outer array.
	array14 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println(array14)

	// Declare and initialize individual elements of the outer
	// and inner array.
	array15 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array15)

	// Declare a two dimensional integer array of two elements.
	var array16 [2][2]int
	// Set integer values to each individual element.
	array16[0][0] = 10
	array16[0][1] = 20
	array16[1][0] = 30
	array16[1][1] = 40

	fmt.Println(array16)

	var array17 [2][2]int
	array17 = array16

	fmt.Println(array17)

	array18 := array16
	fmt.Println(array18)

	// Copy index 1 of array1 into a new array of the same type.
	var array19 [2]int = array18[1]
	fmt.Println(array19)

	// Copy the integer found in index 1 of the outer array
	// and index 0 of the interior array into a new variable of
	// type integer.
	var value int = array18[1][0]
	fmt.Println(value)
}
