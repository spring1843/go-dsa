package strings

import (
	"strings"
	"testing"
)

func TestInMemoryDictionary(t *testing.T) {
	tests := []struct {
		input, allOutputs string
	}{
		{"EXISTS A\nSET A 1\nGET A", "false 1"},
		{"EXISTS A\nSET A 1\nGET A\nEXISTS A\nUNSET A\nGET A\nEXISTS A", "false 1 true nil false"},
		{"GET A\nBEGIN\nSET A 1\nGET A\nROLLBACK\nGET A", "nil 1 nil"},
		{"GET A\nBEGIN\nSET A 1\nGET A\nCOMMIT\nGET A", "nil 1 1"},
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
