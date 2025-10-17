package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {
	myPerson := NewPerson("trump", 76)
	myPerson.changeName("shazeb")

	a := 7
	b := &a
	*b = 9

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(&a)
	fmt.Println(b)

	mySlice := []int{
		1, 2, 3,
	}

	for values := range mySlice {
		fmt.Printf("the value is %d\n", mySlice[values])
	}

	fmt.Printf("this is my details %+v\n", myPerson)
}
