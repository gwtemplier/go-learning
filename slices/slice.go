package main

import (
	"fmt"
)

func main() {

	//course := make([]string, 5, 10)
	courses := []string{"Docker & Kubernetes: The Big Picture",
		"Getting Started with Docker",
		"Getting Started with Kubernetes"}

	fmt.Printf("Length of slice is %d and capacity is %d\n",
		len(courses), cap(courses))

	for _, i := range courses {
		fmt.Println(i)
	}

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(mySlice[4])

	mySlice[1] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]
	// Start index is included and finish index is excluded...
	sliceOfSliceFromImplicitZero := mySlice[0:5]
	fmt.Println(sliceOfSlice, sliceOfSliceFromImplicitZero)

	mySliceAppend := make([]int, 1, 4)
	fmt.Printf("Length starts out as %d with a capacity of %d\n",
		len(mySliceAppend), cap(mySliceAppend))

	for i := 1; i < 17; i++ {
		mySliceAppend = append(mySliceAppend, i)
		fmt.Printf("Slice length is %d but capacity is %d\n",
			len(mySliceAppend), cap(mySliceAppend))
	}

	mySliceMisc := []int{1, 2, 3, 4, 5}
	fmt.Println(mySliceMisc)

	for _, i := range mySliceMisc {
		fmt.Println(i)
	}

	newSlice := []int{10, 20, 30}
	mySliceMisc = append(mySliceMisc, newSlice...)
	fmt.Printf("mySlice NOW contains %d and has a new length of %d and capacity of %d\n",
		mySliceMisc, len(mySliceMisc), cap(mySliceMisc))
}
