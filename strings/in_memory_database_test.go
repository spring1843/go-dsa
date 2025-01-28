package strings

import (
	"strings"
	"testing"
)

/*
TestInMemoryDictionary tests solution(s) with the following signature and problem description:

	func RunDBCommand(cmd string) string

Write an in memory database that stores string key-value pairs and supports SET, GET, EXISTS,
and UNSET commands. It should also allow transactions with BEGIN, COMMIT and ROLLBACK commands.

For example:

	GET A // outputs nil
	BEGIN
	SET A 1
	GET A // outputs 1
	COMMIT
	GET A // outputs 1

At the first GET A, nil is returned because it has never been set. The second and third
GET A will output 1 because the value of A was set as 1.

BEGIN, ROLLBACK and COMMIT are referred to as transactions in databases. A transaction is
started by BEGIN. The commands that are followed by a BEGIN are completely ignored if a ROLLBACK
command is given and actually applied only when COMMIT command is given.

For example:

	BEGIN
	SET A 1
	COMMIT
	BEGIN
	SET A 2
	ROLLBACK
	GET A // outputs 1

The output is 1 because SET A 2 was never committed so the value of A remains 1.
*/
func TestInMemoryDictionary(t *testing.T) {
	tests := []struct {
		input, allOutputs string
	}{
		{"EXISTS A\nSET A 1\nGET A", "false 1"},
		{"EXISTS A\nSET A 1\nGET A\nEXISTS A\nUNSET A\nGET A\nEXISTS A", "false 1 true nil false"},
		{"GET A\nBEGIN\nSET A 1\nGET A\nROLLBACK\nGET A", "nil 1 nil"},
		{"GET A\nBEGIN\nSET A 1\nGET A\nCOMMIT\nGET A", "nil 1 1"},
		{"BEGIN\nSET A 1\nCOMMIT\nBEGIN\nSET A 2\nROLLBACK\nGET A", "1"},
		{"SET A 1\nGET A\nBEGIN\nSET A 2\nGET A\nBEGIN\nUNSET A\nGET A\nCOMMIT\nROLLBACK\nGET A", "1 2 nil 1"},
		{"SET A 2\nGET A\nBEGIN\nSET A 1\nGET A\nCOMMIT\nGET A\nBEGIN\nSET A 2\nGET A\nROLLBACK\nGET A", "2 1 1 2 1"},
	}

	for i, test := range tests {
		allOutputs := ""
		dbs = make([]map[string]string, 0)
		dbs = append(dbs, make(map[string]string))
		for _, cmd := range strings.Split(test.input, "\n") {
			output := RunDBCommand(cmd)
			if output != "" {
				allOutputs += " " + output
			}
		}
		allOutputs = allOutputs[1:]

		if allOutputs != test.allOutputs {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.allOutputs, allOutputs)
		}
	}
}
