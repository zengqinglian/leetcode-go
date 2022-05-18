package maximalNetworkRank

import "testing"

func TestFunc(t *testing.T) {
	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}
	actual := maximalNetworkRank(8, roads)
	expected := 5
	if actual != expected {
		t.Error("failed")
	}
}
