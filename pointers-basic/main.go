package main

import "fmt"

func main() {
	a := 10
	b := &a

	fmt.Println("Variable a:", a)
	fmt.Println("Variable b:", *b)

	newFunction()
}

func newFunction() {
	p := NewPerson("Igor", 37)
	fmt.Println(p.Name, p.Age, p.Phone())
	p.UpdateAge(38)
	fmt.Println(p.Name, p.Age, p.Phone())
	p.UpdatePhone("123-999")
	fmt.Println(p.Name, p.Age, p.Phone())
}

type Person struct {
	Name  string
	Age   int
	phone *string
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) Phone() string {
	if p.phone == nil {
		return ""
	}

	return *p.phone
}

func (p *Person) UpdateAge(age int) {
	p.Age = age
}

func (p *Person) UpdatePhone(phone string) {
	p.phone = &phone
}
