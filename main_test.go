package main

import (
	"fmt"
	"testing"
)

func TestGetCalibrationString(t *testing.T) {
	// Define a slice of anonymous structs for test cases
	testCases := []struct {
		name     string // Name of the subtest
		input    string // Input to the test
		expected int    // Expected result
	}{
		{"Example Case", "a1b2c3", 13},
		{"Example Case 2", "111111", 11},
		{"Example Case 3", "1aaaaaa2", 12},
		{"Example Case 3", "2aaaaaa1", 21},
		{"Example Case 4", "2a34#9a1", 21},
		{"Example Case 5", "aa92bb", 92},
		{"Example Case 6", "aa9cc2bb", 92},

		{"t Case 1", "1abc2", 12},
		{"t Case 2", "pqr3stu8vwx", 38},
		{"t Case 3", "a1b2c3d4e5f", 15},
		{"t Case 4", "treb7uchet", 77},
	}

	for _, tc := range testCases {
		// Use t.Run to execute a subtest for each case
		t.Run(tc.name, func(t *testing.T) {
			got := getCalibrationString(tc.input)
			if got != tc.expected {
				t.Errorf("getCalibrationString(%q) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}

func TestDay1(t *testing.T) {
	s := findCalibrationString()
	fmt.Println("get your star with this:", s)
}
