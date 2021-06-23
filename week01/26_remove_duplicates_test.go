package week01

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesFirst(t *testing.T) {

	num := []int{1, 1, 2}
	result := removeDuplicates(num)
	fmt.Printf("result: %v", result)
}

func TestRemoveDuplicatesSecond(t *testing.T) {

	num := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(num)
	fmt.Printf("result: %v", result)
}
