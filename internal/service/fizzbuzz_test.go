package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzService_GenerateFizzBuzz(t *testing.T) {
	service := NewFizzBuzzService()

	tests := []struct {
		name     string
		int1     int
		int2     int
		limit    int
		str1     string
		str2     string
		expected []string
	}{
		{
			name:     "Classic fizzbuzz",
			int1:     3,
			int2:     5,
			limit:    15,
			str1:     "fizz",
			str2:     "buzz",
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
		},
		{
			name:     "Custom parameters",
			int1:     2,
			int2:     3,
			limit:    6,
			str1:     "foo",
			str2:     "bar",
			expected: []string{"1", "foo", "bar", "foo", "5", "foobar"},
		},
		{
			name:     "Single element",
			int1:     2,
			int2:     3,
			limit:    1,
			str1:     "a",
			str2:     "b",
			expected: []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.GenerateFizzBuzz(tt.int1, tt.int2, tt.limit, tt.str1, tt.str2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFizzBuzzService_processNumber(t *testing.T) {
	service := NewFizzBuzzService()

	tests := []struct {
		name     string
		num      int
		int1     int
		int2     int
		str1     string
		str2     string
		expected string
	}{
		{"Regular number", 1, 3, 5, "fizz", "buzz", "1"},
		{"Multiple of int1", 3, 3, 5, "fizz", "buzz", "fizz"},
		{"Multiple of int2", 5, 3, 5, "fizz", "buzz", "buzz"},
		{"Multiple of both", 15, 3, 5, "fizz", "buzz", "fizzbuzz"},
		{"Zero case", 6, 2, 3, "foo", "bar", "foobar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.processNumber(tt.num, tt.int1, tt.int2, tt.str1, tt.str2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
