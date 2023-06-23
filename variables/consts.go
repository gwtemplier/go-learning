package main

import (
	"fmt"
	"os"
)

func main() {
	const c = 186000
	fmt.Println("The speed of the light is", c, "miles per second")

	fmt.Println(os.Environ())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
