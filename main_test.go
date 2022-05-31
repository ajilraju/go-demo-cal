package main

import "testing"

func TestAdd(t *testing.T) {
	if add(2, 2) != 4 {
		t.Error("Expected 2 + 2 is 4")
	}
}

func TestMulti(t *testing.T) {
	if multi(2, 2) != 4 {
		t.Error("Expected 2 * 2 is 4")
	}
}

func TestTableAdd(t *testing.T) {
	var tests = []struct {
		f_input  int
		s_intput int
		expected int
	}{
		{2, 2, 4},
		{2, 1, 3},
		{1000, 2, 1002},
		{2, 200, 202},
	}

	for _, test := range tests {
		if output := add(test.f_input, test.s_intput); output != test.expected {
			t.Errorf("Test Failed: %d + %d expected: %d, received: %d", test.f_input, test.s_intput, test.expected, output)
		}
	}
}

func TestTableMulti(t *testing.T) {
	var tests = []struct {
		f_input  int
		s_intput int
		expected int
	}{
		{2, 2, 4},
		{2, 1, 2},
		{1000, 2, 2000},
		{2, 200, 400},
	}

	for _, test := range tests {
		if output := multi(test.f_input, test.s_intput); output != test.expected {
			t.Errorf("Test Failed: %d * %d expected: %d, received: %d", test.f_input, test.s_intput, test.expected, output)
		}
	}
}
