package practice_question

import (
	"fmt"
	"strings"
)

func main() {

	input := "abcabcbb"

	var longestPerStart []string
	for i, _ := range input {
		var perIndexBuilder string
		for j := i; j < len(input); j++ {
			if strings.Contains(perIndexBuilder, string(input[j])) {
				longestPerStart = append(longestPerStart, perIndexBuilder)
				break
			} else {
				perIndexBuilder = perIndexBuilder + string(input[j])
				if j == len(input)-1 {
					longestPerStart = append(longestPerStart, perIndexBuilder)
				}
			}
		}
	}

	var longest string
	for _, candidate := range longestPerStart {
		if len(candidate) > len(longest) {
			longest = candidate
		}
	}

	fmt.Println(longest)
}
