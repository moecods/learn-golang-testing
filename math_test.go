package tutorialgolangtesting

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// Table Driven Test

func TestAddNew(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{10, 20, 30},
		{-1, -1, -2},
	}

	for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
        }
    }
}

func TestAddSubtests(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		expected int
	}{
        {"Positive numbers", 2, 3, 5},
        {"Negative numbers", -2, -3, -5},
        {"Mixed numbers", -2, 3, 1},
        {"Zeros", 0, 0, 0},
	}

	for _, tt := range tests {
        tt := tt // capture range variable

		// create subtest for each case
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}


func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}