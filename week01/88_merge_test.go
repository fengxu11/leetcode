package week01

import (
	"fmt"
	"testing"
)

func TestMergeFirst(t *testing.T) {

	fmt.Println(-1 < 0)

	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}

	merge(num1, 3, num2, 3)
}

func TestMergeSecond(t *testing.T) {
	num1 := []int{1}
	num2 := []int{}
	merge(num1, 1, num2, 0)
}
