// response.go this file contains helper functions to format JSON responses for the API
// it includes functions to format success, error and paginated responses
// it uses standard http package for status codes and custom logger package for logging

package utils

import (
	"feast-friends-api/pkg/logger"
	"math"
	"net/http"
)

// this helper function formats a successful JSON response
//	it logs the success message and returns a map with status, message and data
func SuccessResponse(data interface{}, message string) map[string]interface{} {
	logger.Info("success response : %v" , message)

	return map[string]interface{}{
		"status" : "success",
		"message" : message,
		"data" : data,
	}

}

//this func formats an error JSON response
// it logs the error and returns a map with status, message and code
func ErrorResponse(message string, err error, statusCode int) map[string]interface{} {
	logger.Error("error response : %v" , err.Error())

	return map[string]interface{}{
		"status" : "error",
		"message" : message,
		"code" : http.StatusText(statusCode),
	}
}

//this func formats a successful JSON response without message
// it is used for paginated responses
// includes meta info like total count, page, limit and total pages
func PaginatedResponse(message string, data interface{}, totalCount, page, limit int) map[string]interface{}{
	logger.Info("paginated response : page %d limit %d total_count %d", page, limit, totalCount)

	totalpages := 0 
	if limit > 0 {
		totalpages = int(math.Ceil(float64(totalCount) / float64(limit)))
	}

	return map[string]interface{}{
		"status" : "success",
		"message" : message,
		"data" : data,
		"meta" : map[string]interface{}{
			"total_count" : totalCount,
			"page" : page,
			"limit" : limit,
			"total_pages" : totalpages,
		},
	}
}

