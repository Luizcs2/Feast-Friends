
// TODO - Luiz
// 1. Standard JSON response helper functions:
//    - SuccessResponse(data, message)
//    - ErrorResponse(error, statusCode)
//    - PaginatedResponse(data, totalCount, page, limit)
// 2. Consistent error message formatting
// 3. HTTP status code constants
// 4. Response pagination metadata

package utils

import (
	"feast-friends-api/pkg/logger"
)


func SuccessResponse(data interface{}, message string) map[string]interface{} {
	logger.Info("success response : %v" , message)

	return map[string]interface{}{
		"status" : "success",
		"message" : message,
		"data" : data,
	}

}

func ErrorResponse(err error, statusCode int) map[string]interface{} {
	logger.Error("error response : %v" , err.Error())

	return map[string]interface{}{
		"status" : "error",
		"message" : err.Error(),
		"code" : statusCode,
	}
}

func PaginatedResponse(data interface{}, totalCount, page, limit int) map[string]interface{}{
	logger.Info("paginated response : page %d limit %d total_count %d", page, limit, totalCount)
	return map[string]interface{}{
		"status" : "success",
		"data" : data,
		"meta" : map[string]interface{}{
			"total_count" : totalCount,
			"page" : page,
			"limit" : limit,
		},
	}
}

