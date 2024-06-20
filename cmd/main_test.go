package main

import (
	"sandbox2"
	"testing"
)

func TestShiftPath(t *testing.T) {
	suites := []struct {
		name string
		path string
		want []string
	}{
		{name: "3-level route", path: "/a/bc/def", want: []string{"a", "/bc/def"}},
		{name: "2-level route", path: "/bc/def", want: []string{"bc", "/def"}},
		{name: "1-level route", path: "/def", want: []string{"def", "/"}},
		{name: "root route", path: "/", want: []string{"", "/"}},
	}

	for _, suite := range suites {
		t.Run(suite.name, func(t *testing.T) {
			head, tail := sandbox2.ShiftPath(suite.path)

			if head != suite.want[0] {
				t.Errorf("incorrect parsed head, want: %v, got: %v", suite.want[0], head)
			}
			if tail != suite.want[1] {
				t.Errorf("incorrect parsed tail, want: %v, got: %v", suite.want[1], tail)
			}
		})
	}
}
