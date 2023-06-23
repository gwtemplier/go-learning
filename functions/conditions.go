package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//dddLengthMins := 275 //Docker Deep Dive Course
	//cawLengthMins := 30  //Containers on AWS Wavelength course

	if dddLengthMins, cawLengthMins := 275, 30; dddLengthMins > cawLengthMins {
		println("Docker is longuer than AWS")
		if dddLengthMins > 120 {
			println("too long :)")
		}
	} else if dddLengthMins == cawLengthMins {
		println("Both course are the same length")
	} else {
		println("AWS is longuer than Docker")
	}

	switch "foo" {
	case "kubernetes":
		fmt.Println("Case 1. kubernetes with lower-case \"k\".")
	case "Kubernetes":
		fmt.Println("Case 2. Kubernetes with a capital \"K\".")
		fallthrough
	case "K8s":
		fmt.Println("Case 3. Kubernetes short name \"Kates\".")
		fallthrough
	case "Istio":
		fmt.Println("Case 4. Service mesh is the future.")
		fallthrough
	default:
		fmt.Println("Hit the default")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got the following even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got the following odd number -", tmpNum)
	}

	_, err := os.Open("./convert.go")

	if err != nil {
		fmt.Println("This is the error code:", err)
	}

	fmt.Println("This is the error code:", err)
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
