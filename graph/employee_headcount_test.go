package graph

import (
	"testing"
)

/*
TestEmployeeHeadCount tests solution(s) with the following signature and problem description:

	HeadCount(employeeID int) int

Head count of a person in an organization is 1 + the number of all their reports
(direct and indirect). Given a list of employee IDs and their direct reports in each line like
`1,4,5`, `5,7`, `4`, `7`. Where 1 has two direct reports (4 and 5); 5 has one report (7);
4 and 5 have no reports; Return the head count of a given employee ID. For example head
count of 7 is 1, and headcount of 1 is 4.
*/
func TestEmployeeHeadCount(t *testing.T) {
	tests := []struct {
		employeeData string
		headCounts   map[int]int
	}{
		{"A", map[int]int{1: -1}},
		{"1", map[int]int{1: 1}},
		{"A,B\nC,D,E,F\nG\nH\nI", map[int]int{1: -1}},
		{"1,2\n3,4,5,6\n4\n5\n6", map[int]int{1: 2}},
		{"1,2\n3,4,5,6\n4\n5\n6", map[int]int{3: 4}},
		{"1,2\n3,4,5,6\n4\n5\n6,7\n7", map[int]int{3: 5}},
		{"1,4,5\n5,7\n4\n7", map[int]int{7: 1, 4: 1, 5: 2, 1: 4}},
		{"23,29,27\n25\n26,25\n27\n28,32\n29,30\n30\n31\n32", map[int]int{23: 4, 29: 2, 32: 1, 30: 1}},
		{"23,29,27,28\n25\n26,25\n27\n28,32\n29,30\n30\n31\n32", map[int]int{23: 6, 29: 2, 32: 1, 30: 1}},
	}

	for i, test := range tests {
		for employeeID, headCount := range test.headCounts {
			got := HeadCount(test.employeeData, employeeID)
			if got != headCount {
				t.Errorf("Failed test case #%d employee ID %d.  Want %d got %d", i, employeeID, headCount, got)
			}
		}
	}
}
