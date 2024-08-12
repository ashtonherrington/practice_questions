package practice_question

import "fmt"

func main() {
	input := "abcdeffg"

	testValue := "fef"
	testByteHash := NewByteHash(5)
	for _, tv := range testValue {
		testByteHash.Add(byte(tv))
	}

	var containsPermutation bool
	for i := 0; i < len(input)-1; i++ {
		for j := len(input); j > i; j-- {
			compareByteHash := NewByteHash(5)
			for _, k := range input[i:j] {
				compareByteHash.Add(byte(k))
			}
			if testByteHash.EqualContents(compareByteHash) {
				containsPermutation = true
			}
		}
	}
	fmt.Println(input, " contains permutation of ", testValue, containsPermutation)
}
