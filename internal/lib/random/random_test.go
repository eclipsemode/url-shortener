package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"size = 1", 1},
		{"size = 5", 5},
		{"size = 10", 10},
		{"size = 30", 30},
		{"size = 100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size)
			str2 := NewRandomString(tt.size + 1)

			assert.Len(t, str1, tt.size)
			assert.Len(t, str2, tt.size+1)

			// Check that two generated strings are different
			// This is not an absolute guarantee that the function works correctly,
			// but this is a good heuristic for a simple random generator.
			assert.NotEqual(t, str1, str2, "Should be different")
		})
	}
}
