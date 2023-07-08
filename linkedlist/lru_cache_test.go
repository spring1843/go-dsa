package linkedlist

import "testing"

/*
TestLRU tests solution(s) with the following signature and problem description:

	func get(key int) int
	func put(key int, value int) {

Implement a least recently used cache with integer keys and values, where the
least recently used evicted upon insertion when cache is at full capacity.
*/
func TestLRU(t *testing.T) {
	tests := []struct {
		capacity int
		puts     []int // n(i) element is key, n(i+1) element is value
		gets     []int // n(i) element is the key to get, n(i+1) element is the expected value
	}{
		{1, []int{}, []int{1, -1}},
		{1, []int{2, 1, 1, 1, 2, 3, 4, 1}, []int{2, -1, 4, 1}},
		{2, []int{2, 1, 1, 1, 2, 3, 4, 1}, []int{1, -1, 2, 3, 4, 1}},
		{3, []int{2, 1, 1, 1, 2, 3, 4, 1}, []int{1, 1, 2, 3, 4, 1}},
	}

	for i, test := range tests {
		lru := newLRU(test.capacity)

		for j := 0; j < len(test.puts); j += 2 {
			lru.put(test.puts[j], test.puts[j+1])
		}

		for j := 0; j < len(test.gets); j += 2 {
			got := lru.get(test.gets[j])
			want := test.gets[j+1]
			if got != want {
				t.Fatalf("Failed test #%d on get #%d, want %d got %d", i, j, want, got)
			}
		}
	}
}
