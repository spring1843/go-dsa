package problems

import "testing"

func TestNewRehearsalEntry(t *testing.T) {
	input := `## Rehearsal

* [Bubble Sort](./bubble_sort_test.go), [Solution](./bubble_sort.go)
* [Reverse Array In-place](./reverse_inplace_test.go), [Solution](./reverse_inplace.go)
* [Add Two Numbers](./add_two_numbers_test.go), [Solution](./add_two_numbers.go)
* [Find Duplicate in Array](./find_duplicate_in_array_test.go), [Solution](./find_duplicate_in_array.go)
* [Zero Sum Triplets](./zero_sum_triplets_test.go), [Solution](./zero_sum_triplets.go)
* [Product of All Other Elements](./product_of_all_other_elements_test.go), [Solution](./product_of_all_other_elements.go)`

	entries, err := newRehearsalEntry(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(entries) != 6 {
		t.Fatalf("expected 5 entries, got %d", len(entries))
	}

	strOutput := stringRehearsalEntries(entries)
	if strOutput == "" {
		t.Fatal("expected non-empty string, got empty string")
	}
}
