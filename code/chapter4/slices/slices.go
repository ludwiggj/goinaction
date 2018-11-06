package main

import "fmt"

func creatingSlices() {
	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	slice1 := make([]string, 5)
	fmt.Println(slice1)

	// Contains a length of 3 and has a capacity of 5 elements.
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)

	// Create a slice of integers.
	// Make the length larger than the capacity.

	// Compiler Error:
	// len larger than cap in make([]int)

	// slice3 := make([]int, 5, 3)

	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	slice4 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice4)

	// Create a slice of integers.
	// Contains a length and capacity of 3 elements.
	slice5 := []int{10, 20, 30}
	fmt.Println(slice5)

	// Create a slice of strings.
	// Initialize the 100th element with an empty string.
	slice6 := []string{99: "a"}
	fmt.Println(slice6)

	// Create an array of three integers.
	array := [3]int{10, 20, 30}
	fmt.Println(array)

	// Create a slice of integers with a length and capacity of three.
	slice7 := []int{10, 20, 30}
	fmt.Println(slice7)

	// Create a nil slice of integers.
	var nilSlice []int
	fmt.Println(nilSlice)

	// Use make to create an empty slice of integers.
	emptySlice1 := make([]int, 0)
	fmt.Println(emptySlice1)

	// Use a slice literal to create an empty slice of integers.
	emptySlice2 := []int{}
	fmt.Println(emptySlice2)
}

func assignSliceElement() {
	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	// Change the value of index 1.
	slice[1] = 25
	fmt.Println(slice)
}

func sliceOfSlice() {
	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]
	fmt.Println(newSlice)

	// For slice[i:j] with an underlying array of capacity k
	//   Length: j - i
	// Capacity: k - i

	// Change index 1 of newSlice.
	// Change index 2 of the original slice.
	newSlice[1] = 35

	fmt.Println(slice)
	fmt.Println(newSlice)

	// Change index 3 of newSlice.
	// This element does not exist for newSlice.
	newSlice[3] = 45

	// Runtime Exception:
	// panic: runtime error: index out of range
}

func appendElementToSlice() {
	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]
	fmt.Println(newSlice)

	// Allocate a new element from capacity.
	// Assign the value of 60 to the new element.
	newSlice = append(newSlice, 60)
	fmt.Println(slice)
	fmt.Println(newSlice)
}

func appendToIncreaseElementSize() {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}
	fmt.Println(slice)

	// Append a new value to the slice.
	// Assign the value of 50 to the new element.
	newSlice := append(slice, 50)
	fmt.Println(newSlice)
}

func threeIndexSlice() {
	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// Slice the third element and restrict the capacity.
	// Contains a length of 1 element and capacity of 2 elements.
	slice := source[2:3:4]

	// For slice[i:j:k] or [2:3:4]
	//   Length: j - i or 3 - 2 = 1
	// Capacity: k - i or 4 - 2 = 2
	fmt.Println(slice)

	// This slicing operation attempts to set the capacity to 4.
	// This is greater than what is available.
	invalidSlice := source[2:3:6]

	// Runtime Error:
	// panic: runtime error: slice bounds out of range
	fmt.Println(invalidSlice)
}

func sliceLengthCapacityTheSame() {
	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Println(source)

	// Slice the third element and restrict the capacity.
	// Contains a length and capacity of 1 element.
	slice := source[2:3:3]
	fmt.Println(slice)

	// Append a new string to the slice.
	slice = append(slice, "Kiwi")
	fmt.Println(source)
	fmt.Println(slice)
}

func addSlicesTogether() {
	// Create two slices each initialized with two integers.
	s1 := []int{1, 2}
	s2 := []int{3, 4}

	// Append the two slices together and display the results.
	fmt.Printf("%v\n", append(s1, s2...))
}

func sliceIteration() {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}

	// Iterate over each element and display each value.
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d Value-Addr: %X ElemAddr: %X\n", index, value, &value, &slice[index])
	}

	// Iterate over each element and display each value, discarding index
	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}

	// Iterate over each element starting at element 3.
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}

func multiDimensionalSlices() {
	slice := [][]int{{10}, {100, 200}}
	fmt.Println(slice)

	// Append the value of 20 to the first slice of integers.
	slice[0] = append(slice[0], 20)
	fmt.Println(slice)
}

func main() {
	// creatingSlices()
	// assignSliceElement()
	// sliceOfSlice()
	// appendElementToSlice()
	// appendToIncreaseElementSize()
	// threeIndexSlice()
	// sliceLengthCapacityTheSame()
	// addSlicesTogether()
	// sliceIteration()
	multiDimensionalSlices()
}
