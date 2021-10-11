package main

import "fmt"

type OperationFunc func(int, int) int

func main() {
	/* result := add(100, 200)
	fmt.Println("Result : ", result)
	result = subtract(100, 50)
	fmt.Println("Result : ", result) */
	loggerAdd := logger(add)
	loggerSubtract := logger(subtract)
	result := loggerAdd(100, 200)
	fmt.Println("Result : ", result)
	result = loggerSubtract(100, 50)
	fmt.Println("Result : ", result)
}

func logger(fn OperationFunc) OperationFunc {
	return func(x int, y int) int {
		fmt.Println("function invoked with ", x, y)
		result := fn(x, y)
		fmt.Println("function invocation complete with result ", result)
		return result
	}
}

func add(x, y int) int {
	result := x + y
	return result
}

func subtract(x, y int) int {
	result := x - y
	return result
}
