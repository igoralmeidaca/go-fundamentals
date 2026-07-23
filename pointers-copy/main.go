package main

import "fmt"

func main() {
	name := "Pedro"

	person1 := Person{
		Name: &name,
		Age:  26,
	}

	person2 := person1

	fmt.Println("Person name 1:", *person1.Name)
	fmt.Println("Person name 2:", *person2.Name)

	person2.Name = toPointer("Davi")

	fmt.Println("----------------")
	fmt.Println("Person name 1:", *person1.Name)
	fmt.Println("Person name 2:", *person2.Name)

	person3 := Person{
		Name: toPointer("Leonardo"),
		Age:  26,
	}
	person4 := deepCopy(person3)

	*person4.Name = "Lucas"
	fmt.Println("----------------")
	fmt.Println("Person name 3:", *person3.Name)
	fmt.Println("Person name 4:", *person4.Name)

	people1 := []Person{person1, person2, person3, person4}
	people2 := deepCopyList(people1)

	fmt.Println("People list 1:", people1)
	fmt.Println("People list 2:", people2)
}

func toPointer(s string) *string {
	return &s
}

func deepCopy(source Person) Person {
	var target Person

	target.Age = source.Age
	target.Name = toPointer(*source.Name)

	return target
}

func deepCopyList(source []Person) []Person {
	target := make([]Person, len(source))
	for i, person := range source {
		target[i] = deepCopy(person)
	}

	return target
}

type Person struct {
	Name *string
	Age  int
}
