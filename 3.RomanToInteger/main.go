package main

import "fmt"

func romanToInt(romanNum string) int {
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0

	for i, char := range romanNum {
		current := romanValues[string(char)]

		next := 0
		if i+1 < len(romanNum) {
			next = romanValues[string(romanNum[i+1])]
		}

		if current < next {
			sum -= current
		} else {
			sum += current
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}