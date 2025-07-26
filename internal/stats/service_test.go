package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsService_RecordRequest(t *testing.T) {
	service := NewService()

	// Record the same request multiple times
	service.RecordRequest(3, 5, 100, "fizz", "buzz")
	service.RecordRequest(3, 5, 100, "fizz", "buzz")
	service.RecordRequest(2, 7, 50, "foo", "bar")

	mostFrequent := service.GetMostFrequentRequest()
	assert.NotNil(t, mostFrequent)
	assert.Equal(t, 3, mostFrequent.Int1)
	assert.Equal(t, 5, mostFrequent.Int2)
	assert.Equal(t, 100, mostFrequent.Limit)
	assert.Equal(t, "fizz", mostFrequent.Str1)
	assert.Equal(t, "buzz", mostFrequent.Str2)
	assert.Equal(t, 2, mostFrequent.Count)
}

func TestStatsService_GetMostFrequentRequest_Empty(t *testing.T) {
	service := NewService()

	mostFrequent := service.GetMostFrequentRequest()
	assert.Nil(t, mostFrequent)
}

func TestStatsService_GenerateKey(t *testing.T) {
	service := NewService()

	key1 := service.generateKey(3, 5, 100, "fizz", "buzz")
	key2 := service.generateKey(3, 5, 100, "fizz", "buzz")
	key3 := service.generateKey(2, 7, 50, "foo", "bar")

	assert.Equal(t, key1, key2)
	assert.NotEqual(t, key1, key3)
}
