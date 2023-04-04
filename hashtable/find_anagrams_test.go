package hashtable

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		words    []string
		anagrams [][]string
	}{
		{[]string{"foo", "bar", "baz"}, [][]string{}},
		{[]string{"foo", "cat", "tac", "act"}, [][]string{{"cat", "tac", "act"}}},
		{[]string{"car", "cap", "arc", "tac"}, [][]string{{"car", "arc"}}},
	}

	for i, test := range tests {
		got := FindAnagrams(test.words)
		if !reflect.DeepEqual(got, test.anagrams) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.anagrams, got)
		}
	}
}
