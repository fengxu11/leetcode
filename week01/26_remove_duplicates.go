package week01

import "fmt"

func removeDuplicates(num []int) int {

	n := 0
	for i := 0; i < len(num); i++ {

		if i == 0 || num[i] != num[i-1] {
			num[n] = num[i]
			n++
		}

		fmt.Println("num: ", num, "i: ", i, "i-1: ", i-1)
	}

	return n
}
