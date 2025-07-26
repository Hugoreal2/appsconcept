package service

import (
	"strconv"

	"github.com/Hugoreal2/appsconcept/internal/stats"
)

// FizzBuzzService handles the business logic for fizzbuzz operations
type FizzBuzzService struct {
	statsService *stats.Service
}

// NewFizzBuzzService creates a new fizzbuzz service
func NewFizzBuzzService(statsService *stats.Service) *FizzBuzzService {
	return &FizzBuzzService{
		statsService: statsService,
	}
}

// GenerateFizzBuzz generates the fizzbuzz sequence based on the provided parameters
func (s *FizzBuzzService) GenerateFizzBuzz(int1, int2, limit int, str1, str2 string) []string {
	// Record the request for statistics
	s.statsService.RecordRequest(int1, int2, limit, str1, str2)

	result := make([]string, 0, limit)

	for i := 1; i <= limit; i++ {
		value := s.processNumber(i, int1, int2, str1, str2)
		result = append(result, value)
	}

	return result
}

// processNumber processes a single number according to fizzbuzz rules
func (s *FizzBuzzService) processNumber(num, int1, int2 int, str1, str2 string) string {
	isMultipleOfInt1 := num%int1 == 0
	isMultipleOfInt2 := num%int2 == 0

	switch {
	case isMultipleOfInt1 && isMultipleOfInt2:
		return str1 + str2
	case isMultipleOfInt1:
		return str1
	case isMultipleOfInt2:
		return str2
	default:
		return strconv.Itoa(num)
	}
}
