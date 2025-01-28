package strings

import "testing"

/*
TestReverseVowels tests solution(s) with the following signature and problem description:

	func ReverseVowels(str string) (string, error)

Given a string, return the same string while reversing the vowels {"a", "e", "i", "o", "u"}
appear in it.

For example given "coat", return "caot", because the vowels are "o" and "a" and their positions
are reversed.
*/
func TestReverseVowels(t *testing.T) {
	tests := []struct {
		word     string
		reversed string
	}{
		{"umbrella", "ambrellu"},
		{"coat", "caot"},
		{"eventually", "avunteelly"},
		{"sooner rather than later", "seanar rethar then lotor"},
	}

	for i, test := range tests {
		got, err := ReverseVowels(test.word)
		if err != nil {
			t.Fatalf("Failed test case #%d. Unexpected Error. Error: %s", i, err)
		}
		if got != test.reversed {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.reversed, got)
		}
	}
}
