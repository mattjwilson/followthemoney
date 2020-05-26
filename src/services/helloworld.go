package main

import "fmt"

func main() {
	defer after()
	fmt.Println("Hello world")
}

func after() {
	// stuff

	mybool := true
	pointerBool(&mybool)

	fmt.Println(mybool)
}
