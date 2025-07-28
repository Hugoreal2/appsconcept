package service

import (
	"fmt"
	"sync"
)

// RequestStats represents the statistics for a fizzbuzz request
type RequestStats struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Count int    `json:"count"`
}

// StatsService handles request statistics
type StatsService struct {
	mu    sync.RWMutex
	stats map[string]*RequestStats
}

// NewStatsService creates a new statistics service
func NewStatsService() *StatsService {
	return &StatsService{
		stats: make(map[string]*RequestStats),
	}
}

// RecordRequest records a fizzbuzz request for statistics
// It increments the count for the request parameters or creates a new entry if it doesn't exist.
// It's thread-safe, using a mutex to protect access to the stats map.
func (s *StatsService) RecordRequest(int1, int2, limit int, str1, str2 string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := s.generateKey(int1, int2, limit, str1, str2)

	if existing, exists := s.stats[key]; exists {
		existing.Count++
	} else {
		s.stats[key] = &RequestStats{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
			Count: 1,
		}
	}
}

// GetMostFrequentRequest returns the most frequently requested parameters
func (s *StatsService) GetMostFrequentRequest() *RequestStats {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var mostFrequent *RequestStats
	maxCount := 0

	for _, stat := range s.stats {
		if stat.Count > maxCount {
			maxCount = stat.Count
			mostFrequent = stat
		}
	}

	return mostFrequent
}

// generateKey creates a unique key for the request parameters
func (s *StatsService) generateKey(int1, int2, limit int, str1, str2 string) string {
	return fmt.Sprintf("%d:%d:%d:%s:%s", int1, int2, limit, str1, str2)
}
