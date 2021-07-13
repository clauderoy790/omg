package main

import "testing"

func TestServers(t *testing.T) {
	tests := []Pair{
		{[]int{5, 3, 1}, 2},
		{[]int{5, 4, 1, 2}, 3},
		{[]int{3, 2, 1}, 4},
		{[]int{2, 3}, 1},
		{[]int{1}, 2},
		{[]int{}, 1},
	}

	for _, pair := range tests {
		ret := nextServerNumber(pair.key)
		if ret != pair.value {
			t.Errorf("next server number should be %v but we returned %v ", pair.value, ret)
		}
	}
}

type Pair KeyValue

type KeyValue struct {
	key   []int
	value int
}
