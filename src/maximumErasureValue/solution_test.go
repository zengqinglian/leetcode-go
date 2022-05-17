package maximumErasureValue

import (
	"testing"
)

func TestFunc(t *testing.T) {
	nums := []int{4, 2, 4, 5, 6}
	actual := maximumUniqueSubarray(nums)
	expected := 17
	if actual != expected {
		t.Error("failed")
	}
}
