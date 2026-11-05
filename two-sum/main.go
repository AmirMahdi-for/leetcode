package main

import "fmt"

func twoSum(arr []int, target int) []int {

	var result []int

	for i, v := range arr {
		for j, w := range arr {
			sum := w + v
			if sum == target && j > i {
				result = append(result, i, j)
			}
		}
	}

	return result
}

func main() {
	arr := []int{2, 6, 11, 7, 4}
	target := 9

	fmt.Println(twoSum(arr, target))

}
