package main

import "fmt"

func main() {
	sliceArrayDifferences()
	insecureSlicesCopy()
	secureSlicesCopy()
	deepCopy()
}

func deepCopy() {
	fmt.Println("Deep Copy")
	matriz1 := [][]int{
		{1, 2},
		{3, 4},
	}
	matriz2 := make([][]int, len(matriz1))

	for i, row := range matriz1 {
		matriz2[i] = make([]int, len(row))
		copy(matriz2[i], row)
	}

	fmt.Println("Matriz 1:", matriz1)
	fmt.Println("Matriz 2:", matriz2)

	matriz2[0][0] = 100
	fmt.Println("--------")
	fmt.Println("Matriz 1:", matriz1)
	fmt.Println("Matriz 2:", matriz2)
}

func secureSlicesCopy() {
	fmt.Println("Secure Slices Copy")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice1[0] = 100
	fmt.Println("--------")
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
}

func insecureSlicesCopy() {
	fmt.Println("Insecure Slices Copy")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice1[0] = 100
	fmt.Println("--------")
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice2 = append(slice2, 6)
	slice2[0] = 99
	fmt.Println("--------")
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
}

func sliceArrayDifferences() {
	fmt.Println("Slice Array Differences")
	slice := []int{1, 2, 3, 4, 5}  // dynamic capacity
	array := [5]int{1, 2, 3, 4, 5} // fixed capacity

	fmt.Println("Array size:", len(array))
	fmt.Println("Array capacity:", cap(array))
	fmt.Println("Slice size:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	fmt.Println("--------")

	slice = append(slice, 6)

	fmt.Println("Slice size:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))
}
