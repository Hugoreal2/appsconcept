package handler

import (
	"net/http"

	"github.com/Hugoreal2/appsconcept/internal/service"

	"github.com/gin-gonic/gin"
)

// FizzBuzzHandler handles HTTP requests for fizzbuzz operations
type FizzBuzzHandler struct {
	service      *service.FizzBuzzService
	statsService *service.StatsService
}

// NewFizzBuzzHandler creates a new fizzbuzz handler
func NewFizzBuzzHandler(service *service.FizzBuzzService, statsService *service.StatsService) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		service:      service,
		statsService: statsService,
	}
}

// FizzBuzzRequest represents the request parameters for fizzbuzz
type FizzBuzzRequest struct {
	Int1  int    `form:"int1" binding:"required,min=1"`
	Int2  int    `form:"int2" binding:"required,min=1"`
	Limit int    `form:"limit" binding:"required,min=1,max=10000"`
	Str1  string `form:"str1" binding:"required"`
	Str2  string `form:"str2" binding:"required"`
}

// FizzBuzzResponse represents the response for fizzbuzz
type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// StatsResponse represents the response for the most frequent request statistics
type StatsResponse struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Count int    `json:"count"`
}

// NoStatsResponse represents the response when no requests have been recorded
type NoStatsResponse struct {
	Message string `json:"message"`
}

// FizzBuzz handles the fizzbuzz endpoint
// @Summary Generate FizzBuzz sequence
// @Description Generate a customizable FizzBuzz sequence based on provided parameters
// @Tags fizzbuzz
// @Accept json
// @Produce json
// @Param int1 query int true "First integer for replacement logic" minimum(1)
// @Param int2 query int true "Second integer for replacement logic" minimum(1)
// @Param limit query int true "Upper limit for the sequence" minimum(1) maximum(10000)
// @Param str1 query string true "String to replace multiples of int1"
// @Param str2 query string true "String to replace multiples of int2"
// @Success 200 {object} FizzBuzzResponse
// @Failure 400 {object} ErrorResponse
// @Router /fizzbuzz [get]
func (h *FizzBuzzHandler) FizzBuzz(c *gin.Context) {
	var req FizzBuzzRequest

	// Parse and validate query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid parameters",
			Message: err.Error(),
		})
		return
	}

	// Additional validation
	if req.Int1 == req.Int2 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid parameters",
			Message: "int1 and int2 must be different",
		})
		return
	}

	// Record the request for statistics (handler responsibility)
	h.statsService.RecordRequest(req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)

	// Generate fizzbuzz sequence
	result := h.service.GenerateFizzBuzz(req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)

	c.JSON(http.StatusOK, FizzBuzzResponse{
		Result: result,
	})
}

// GetStats handles the statistics endpoint
// @Summary Get request statistics
// @Description Returns the most frequently requested parameters and their count
// @Tags stats
// @Accept json
// @Produce json
// @Success 200 {object} StatsResponse "Most frequent request statistics"
// @Success 404 {object} NoStatsResponse "No requests recorded yet"
// @Router /stats [get]
func (h *FizzBuzzHandler) GetStats(c *gin.Context) {
	stats := h.statsService.GetMostFrequentRequest()

	if stats == nil {
		c.JSON(http.StatusNotFound, NoStatsResponse{
			Message: "No requests recorded yet",
		})
		return
	}

	c.JSON(http.StatusOK, StatsResponse{
		Int1:  stats.Int1,
		Int2:  stats.Int2,
		Limit: stats.Limit,
		Str1:  stats.Str1,
		Str2:  stats.Str2,
		Count: stats.Count,
	})
}
