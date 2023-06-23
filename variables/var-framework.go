package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var ()

func main() {
	name, course := "Nigel Poulton", "Getting Started with kubernetes"
	module := "4"
	clip := 2
	//courseComplete := false

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Name is of type", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	fmt.Println("Memory adress of the variable module", &module)
	var ptr *string = &course
	fmt.Println("Pointing course variable at adress,", ptr, "which holds this value,", *ptr)

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)
	fmt.Println("\nHi", name, "your current course is", course)

	name = os.Getenv("USER")

	fmt.Println("User logged", name)
}

func updateCourse(course *string) string {
	*course = "Getting started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
