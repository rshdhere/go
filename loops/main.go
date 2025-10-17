package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := [2]string{
		"dog",
		"cat",
	}

	// can be assigned the below way too
	// animals[0] = "cat"
	// animals[1] = "dog"

	fmt.Println(animals)

	// this one is called a slice
	plants := []string{
		"cactus",
		"money",
	}

	plants = append(plants, "rose")

	plants = slices.Delete(plants, 0, 1)

	fmt.Println(plants)
}
