package practice_question

import (
	"fmt"
)

func main() {

	divisors := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	input := 1735
	var builder string

	for i := len(divisors) - 1; i >= 0; i-- {
		if input >= divisors[i] {
			builder = builder + "1"
			fmt.Println("input", divisors[i])
			input = input - divisors[i]
		} else {
			builder = builder + "0"
		}
	}

	fmt.Println(builder)
}
