package practice_question

import (
	"fmt"
)

var validPairs []string

func main() {
	determinePairs(3, 3, "")
	fmt.Println(validPairs)
}

func determinePairs(forwardNeeded, backwardNeeded int, currentString string) {
	if forwardNeeded == 0 && backwardNeeded == 0 {
		// fmt.Println(currentString, forwardNeeded, backwardNeeded)
		validPairs = append(validPairs, currentString)
		return
	}

	if forwardNeeded == 3 && backwardNeeded == 3 {
		determinePairs(forwardNeeded-1, backwardNeeded, currentString+"(")
	}

	if forwardNeeded > 0 && backwardNeeded > 0 {
		if backwardNeeded < forwardNeeded {
			determinePairs(forwardNeeded, backwardNeeded-1, currentString+")")
		}
		determinePairs(forwardNeeded-1, backwardNeeded, currentString+"(")
	}

	if backwardNeeded > 0 {
		determinePairs(forwardNeeded, backwardNeeded-1, currentString+")")
	}
}
