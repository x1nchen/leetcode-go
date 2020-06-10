package p009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var cases = []struct {
		x     int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{10001, true},
		{1001, true},
	}

	for _, tt := range cases {
		result := isPalindrome(tt.x)
		assert.Equal(t, result, tt.expected)
	}
}
