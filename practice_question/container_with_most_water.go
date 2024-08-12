package practice_question

import "fmt"

func addNumbers(a uint32, b uint32) uint32 {
	return (a + b)
}

func main() {
	// index position is X axis, value is Y axis
	graph := []int32{3, 9, 4, 8, 2, 6, 1}

	//extract entire thing into unit test, use table driven test approach
	var largest *XCooridinateSet
	for forwardIndex, value := range graph {
		for backwardIndex := len(graph) - 1; backwardIndex > forwardIndex; backwardIndex-- {
			width := backwardIndex - forwardIndex
			// make testable method that determines heigh with 3 unit tests of greater lesser equal
			height := value
			if graph[backwardIndex] < height {
				height = graph[backwardIndex]
			}

			result := int32(width) * height
			if largest == nil {
				largest = &XCooridinateSet{
					small: int32(forwardIndex),
					large: int32(backwardIndex),
					holds: result,
				}
			}
			if largest != nil {
				if result > largest.holds {
					largest = &XCooridinateSet{
						small: int32(forwardIndex),
						large: int32(backwardIndex),
						holds: result,
					}
				}
			}
		}
	}

	if largest == nil {
		fmt.Errorf("somehow, some way, this remained null! Was the input slice empty?")
	}

	fmt.Println(largest.small, largest.large, largest.holds)
}

type XCooridinateSet struct {
	small, large, holds int32
}
