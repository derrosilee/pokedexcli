package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"Hello World", []string{"hello", "world"}},
		{"  leading spaces", []string{"leading", "spaces"}},
		{"trailing spaces  ", []string{"trailing", "spaces"}},
		{"  extra   spaces   ", []string{"extra", "spaces"}},
		{"CAPS LOCK", []string{"caps", "lock"}},
		{"mixed CAPS and small letters", []string{"mixed", "caps", "and", "small", "letters"}},
		{"", []string{}}, //empty string
	}
	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := cleanInput(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
