package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

/*
TestGenerateParentheses tests solution(s) with the following signature and problem description:

	GenerateParentheses(n int) []string

Given an integer n produce all valid variations of arranging
n pair of parentheses. e.g. for `2` it should return `()()` and `(())`.
*/
func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		n                int
		validParenthesis []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
		{5, []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"}},
	}

	for i, test := range tests {
		got := GenerateParentheses(test.n)
		if len(got) > 0 {
			sort.Strings(got)
		}
		if !reflect.DeepEqual(test.validParenthesis, got) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.validParenthesis, got)
		}
	}
}
