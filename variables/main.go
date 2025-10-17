package main

import "fmt"

func main() {
	var name string = "mbappe"
	fmt.Printf("hello my name is %s\n", name)

	myName := "ronaldo"
	myInt := 10
	myFloat := 23.455

	fmt.Printf("hello my name is %s my int is %d and my float is %f\n", myName, myInt, myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	myFriendsName = "sana"

	fmt.Printf("my friend name is %s and his bool is %t and his int is %d\n", myFriendsName, myBool, myOtherInt)
}
