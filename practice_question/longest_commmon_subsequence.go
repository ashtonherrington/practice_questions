package practice_question

import (
	"fmt"
)

func main() {

	input1 := "abccba"
	input2 := "abddba"
	// input1 :=  "zfadeg"
	// input2 := "cdfsdg"
	// input1 := "abd"
	// input2 := "badc"

	var currentSharedSet []byte
	var startingIndex *int
	for i := 0; i < len(input1); i++ {
		for j := i; j < len(input2); j++ {
			if input1[i] == input2[j] {
				if startingIndex == nil {
					startingIndex = &i
				}
				currentSharedSet = append(currentSharedSet, input1[i])
				break
			}

		}
	}
	fmt.Println(string(currentSharedSet))
}
