package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "nigel poulton"
	course := "getting started with kubernetes"

	fmt.Println(converter(author, course))

	bestFinish := championshipFinishes(12, 5, 6, 4, 3, 3)

	println(bestFinish)
}

func converter(author, course string) (a, c string) {

	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}

func championshipFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
