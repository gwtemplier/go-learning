package main

import (
	"fmt"
	"time"
)

func main() {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			println("boummm!!!!")
			break
		}
		println(timer)
		time.Sleep(1 * time.Second)
	}

	//range for
	coursesInProg := []string{
		"Docker & Kubernetes: The Big Picture",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive"}
	coursesCompleted := []string{
		"Docker & Kubernetes: The Big Picture",
		"Docker Deep Dive"}

	for _, i := range coursesInProg {
		println(i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nHey we found a clash.", i, "is in both lists.")
			}
		}
	}

	// Key word "continu"
	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	// key word break and break breakpoint
	//breakpoint:
	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			break // break breakpoint
		}

		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

}
