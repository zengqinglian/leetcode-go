package sorttheJumbledNumbers

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	mapping := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	res := sortJumbled(mapping, nums)
	fmt.Println(res)
}
