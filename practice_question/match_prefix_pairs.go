package practice_question

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"abs", "app", "be", "apple", "bee", "better", "bet", "absolute"}

	storage := make(map[string][]string)

	for _, v := range input {
		storage[v] = nil
		for _, eclipsingV := range input {

			if v == eclipsingV {
				continue
			}

			if strings.HasPrefix(eclipsingV, v) {
				storage[v] = append(storage[v], eclipsingV)
			}
		}
	}

	fmt.Println(storage)
}
