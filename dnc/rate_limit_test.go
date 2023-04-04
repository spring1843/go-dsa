package dnc

import (
	"reflect"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	var tests = []struct {
		limitPerSecond  int
		firstCallTimes  int
		sleep           int
		secondCallTimes int
		want            []bool
	}{
		{0, 0, 0, 0, []bool{}},
		{0, 1, 0, 0, []bool{false}},
		{0, 1, 0, 1, []bool{false, false}},
		{1, 2, 0, 2, []bool{true, false, false, false}},
		{2, 5, 0, 5, []bool{true, true, false, false, false, false, false, false, false, false}},
		{3, 5, 1, 5, []bool{true, true, true, false, false, true, true, true, false, false}},
		{10, 100, 1, 100, []bool{true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}},
	}

	for i, test := range tests {
		rateLimitEvents = make([]int64, 0)
		got := make([]bool, 0)
		for i := 0; i < test.firstCallTimes; i++ {
			got = append(got, IsAllowed(test.limitPerSecond))
		}
		time.Sleep(time.Duration(test.sleep) * time.Second)
		for i := 0; i < test.secondCallTimes; i++ {
			got = append(got, IsAllowed(test.limitPerSecond))
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.want, got)
		}
	}
}
