package main

import (
	"basics/imports"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	newTicket := imports.Ticket{
		ID:    123,
		Event: "yada-yada",
	}

	newTicket.PrintEvent()

	fmt.Println(newTicket)
}
