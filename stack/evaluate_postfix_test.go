package stack

import "testing"

func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		postfix   []string
		expectErr bool
		outcome   float64
	}{
		{[]string{""}, true, -1},
		{[]string{"+"}, true, -1},
		{[]string{"A", "B", "+"}, true, -1},
		{[]string{"1", "2", "+"}, false, 3},
		{[]string{"1", "2", "3", "+", "*"}, false, 5},
		{[]string{"1", "2", "3", "+", "+"}, false, 6},
		{[]string{"1", "3", "2", "-", "+"}, false, 2},
		{[]string{"9", "3", "/"}, false, 3},
		{[]string{"3", "9", "3", "/", "-"}, false, 0},
		{[]string{"1", "2", "3", "4", "5", "*", "+", "+", "*"}, false, 25},
		{[]string{"1", "2", "3", "+", "4", "5", "*", "+", "*"}, false, 25},
	}

	for i, test := range tests {
		got, err := EvaluatePostfixExpression(test.postfix)
		if err != nil && !test.expectErr {
			t.Fatalf("Did not expect an error. Error : %s", err)
		}

		if got != test.outcome {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.outcome, got)
		}
	}
}
