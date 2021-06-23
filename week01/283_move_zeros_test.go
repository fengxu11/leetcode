package week01

import (
	"fmt"
	"testing"
)

func TestMoveZerosFirst(t *testing.T) {

	num := []int{0, 1, 0, 3, 12}
	moveZero(num)
	fmt.Println(num)
}
