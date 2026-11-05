package main

import "fmt"

func twoSum(arr []int, target int) []int {

	var result []int

	for i := 0; i <= len(arr); i++ {
		j := i + 1
		if j < len(arr) {
			a := arr[i]
			b := arr[i+1]
			sum := a + b
			if sum == target {
				result = append(result, a, b)
			}
		}
	}

	return result
}

func main() {
	arr := []int{2, 6, 11, 7, 3}
	target := 9

	fmt.Println(twoSum(arr, target))
}
