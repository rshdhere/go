package main

import (
	"fmt"
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

	// plants = slices.Delete(plants, 0, 1)

	fmt.Println(plants)

	for i := 0; i < len(plants); i++ {
		fmt.Printf("this is my plant %s\n", plants[i])
	}

	for j := range animals {
		fmt.Printf("this is my animal %s\n", animals[j])
	}

	for index, value := range animals {
		fmt.Printf("this is my index %d and this is my %s\n", index, value)
	}

	for value := range 11 {
		fmt.Println(value)
	}

	// while loops arent there in go
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}
