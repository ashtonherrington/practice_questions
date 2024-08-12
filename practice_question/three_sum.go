package practice_question

import (
	"fmt"
	"sort"
)

func main() {
	input := []int32{-1, 0, 1, 2, -1, -4}

	var foundSets [][]int
	// instead of ending at the final index, end 2 before it as we are looking only for triplets
	for i1 := 0; i1 < len(input)-3; i1++ {
		for i2 := i1 + 1; i2 < len(input)-2; i2++ {
			for i3 := i2 + 1; i3 < len(input)-1; i3++ {
				if input[i1]+input[i2]+input[i3] == 0 {
					set := []int{int(input[i1]), int(input[i2]), int(input[i3])}
					// Sort before comparison and insertion
					sort.Ints(set)
					// works in IDE, not in hackerrank IDE
					//    if !slices.Contains(foundSets, set) {
					//        foundSets = append(foundSets, set)
					//    }
					found, err := contains(foundSets, set)
					if err != nil {
						// something went terrible wrong, log to inform
						fmt.Errorf("something bad happened: %w", err)
						return
					}
					if !found {
						foundSets = append(foundSets, set)
					}
				}
			}
		}
	}

	fmt.Println(foundSets)
}

func contains(s [][]int, e []int) (bool, error) {
	if len(s) == 0 {
		return false, nil // set contains no values yet
	}
	for _, a := range s {
		if len(a) != 3 || len(e) != 3 {
			return false, fmt.Errorf("slices should have 3 elements at this point, in set length: %d: insertion set length: %d", len(a), len(e))
		}

		areEqual := testEq(a, e)
		if areEqual {
			return true, nil
		}
	}
	return false, nil
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
