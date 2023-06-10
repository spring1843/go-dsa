package graph

import (
	"errors"
	"testing"
)

func TestEmployeeHeadCount(t *testing.T) {
	tests := []struct {
		employeeData string
		expectedErr  error
		headCounts   map[int]int
	}{
		{"A", errInvalidInteger, map[int]int{1: -1}},
		{"1", nil, map[int]int{1: 1}},
		{"1,2\n3,4,5,6\n4\n5\n6", nil, map[int]int{1: 2}},
		{"1,2\n3,4,5,6\n4\n5\n6", nil, map[int]int{3: 4}},
		{"1,2\n3,4,5,6\n4\n5\n6,7\n7", nil, map[int]int{3: 5}},
		{"1,4,5\n5,7\n4\n7", nil, map[int]int{7: 1, 4: 1, 5: 2, 1: 4}},
		{"23,29,27\n25\n26,25\n27\n28,32\n29,30\n30\n31\n32", nil, map[int]int{23: 4, 29: 2, 32: 1, 30: 1}},
		{"23,29,27,28\n25\n26,25\n27\n28,32\n29,30\n30\n31\n32", nil, map[int]int{23: 6, 29: 2, 32: 1, 30: 1}},
	}

	for i, test := range tests {
		headCounter, err := NewHeadCount(test.employeeData)
		if !errors.Is(err, test.expectedErr) {
			t.Errorf("Failed test case #%d. Unexpected error. Want %s got %s", i, test.expectedErr, err)
		}

		for employeeID, headCount := range test.headCounts {
			got := headCounter.HeadCount(employeeID)
			if got != headCount {
				t.Errorf("Failed test case #%d employee ID %d.  Want %d got %d", i, employeeID, headCount, got)
			}
		}
	}
}
