package week01

import "fmt"

func moveZero(num []int) {

	n := 0
	for i := 0; i < len(num); i++ {

		if num[i] != 0 {
			num[n] = num[i]
			n++
		}
	}

	fmt.Println(num)
	fmt.Println("n: ", n)

	for n < len(num) {
		num[n] = 0
		n++
	}
}
