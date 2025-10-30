package insertionsort

import "fmt"

func ExampleMoveLeft() {
	a1 := []int{1, 3, 4, 5, 2}
	a2 := []int{1, 5, 3, 4, 2}

	MoveLeft(a1, 4)
	fmt.Println(a1)

	MoveLeft(a2, 4)
	fmt.Println(a2)

	// Output:
	// [1 2 3 4 5]
	// [1 2 5 3 4]
}

func ExampleInsertionSort() {
	a1 := []int{1, 5, 4, 7, 42, 23, 38, -3, 8, 0}
	InsertionSort(a1)
	fmt.Println(a1)

	// Output:
	// [-3 0 1 4 5 7 8 23 38 42]
}
