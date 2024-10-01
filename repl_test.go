package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello unit test",
			expected: []string{
				"hello",
				"unit",
				"test",
			},
		},
	}

	for _, cs := range cases {
		res := cleanInput(cs.input)
		if len(res) != len(cs.expected) {
			t.Errorf("len mismatch, actual [%v] - expected [%v]", len(res), len(cs.expected))
			continue
		}

		for i := range res {
			actualWord := res[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("content mismatch: actual [%v] - expected [%v]", actualWord, expectedWord)
			}
		}
	}

}
