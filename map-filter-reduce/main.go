package main

import "fmt"

func main() {
	list := mySlice{1, 2, 3, 4, 5, 6}

	result := list.Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 10
	}).Reduce(func(i1, i2 int) int {
		return i1 + i2
	}, 0)

	fmt.Println(result)
}

type mySlice []int

func (m mySlice) Filter(cond func(int) bool) mySlice {
	var result mySlice

	for _, number := range m {
		if cond(number) {
			result = append(result, number)
		}
	}

	return result
}

func (m mySlice) Map(trans func(int) int) mySlice {
	var result mySlice

	for _, number := range m {
		result = append(result, trans(number))
	}

	return result
}

func (m mySlice) Reduce(acc func(int, int) int, init int) int {
	result := init

	for _, number := range m {
		result = acc(result, number)
	}

	return result
}
