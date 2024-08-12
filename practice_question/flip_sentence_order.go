package practice_question

import (
	"fmt"
	"strings"
)

func main() {
	input := "I have a lovely bunch of coconuts"
	words := strings.Split(input, " ")

	var builder string
	for i := len(words) - 1; i >= 0; i-- {
		builder = builder + words[i]
		if i != 0 {
			builder = builder + " "
		}
	}

	fmt.Println(builder)
}
