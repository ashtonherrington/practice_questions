package practice_question

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	result := decodeString("ab2[c]3[de4[f]]")
	//abccdeffffdeffffdeffff
	fmt.Println(result)
}

func decodeString(currentString string) string {
	if strings.Contains(currentString, "[") {

		var depth int
		var startIndex, endIndex *int
		var copies []byte
		var numberDone bool
		for i := 0; i < len(currentString); i++ {
			if isDigit(currentString[i]) {
				if !numberDone {
					copies = append(copies, currentString[i])
				}
			}

			if currentString[i] == '[' {
				numberDone = true
				depth++
				if startIndex == nil {
					indexAddress := i + 1
					startIndex = &indexAddress
				}
			}
			if currentString[i] == ']' {
				depth--
				if depth == 0 && endIndex == nil {
					indexAddress := i
					endIndex = &indexAddress
				}
			}
		}

		intCopies, _ := strconv.Atoi(string(copies))
		var builder string
		for i := 0; i < intCopies; i++ {
			builder += currentString[*startIndex:*endIndex]
		}
		if len(currentString) > *endIndex {
			return currentString[:*startIndex-2] + decodeString(builder+currentString[*endIndex+1:])
		} else {
			return currentString[:*startIndex-2] + decodeString(builder)
		}
	}

	return currentString
}
func isDigit(value byte) bool {
	if value >= '0' && value <= '9' {
		return true
	}
	return false
}
