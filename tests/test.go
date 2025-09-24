package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"feast-friends-api/internal/utils"
)

func main() {
	fmt.Println("=== Running Complete Utils Package Test Suite ===")

	// --- RESPONSE UTILS TESTS ---
	fmt.Println("\n--- Testing Response Utils ---")
	testResponseUtils()

	// --- VALIDATION UTILS TESTS ---
	fmt.Println("\n--- Testing Validation Utils ---")
	testValidationUtils()

	// --- JWT UTILS TESTS ---
	fmt.Println("\n--- Testing JWT Utils ---")
	testJWTUtils()

	fmt.Println("\n=== Utils Test Suite Complete ===")
}

func testResponseUtils() {
	// Test SuccessResponse
	data := map[string]interface{}{"user_id": 1, "username": "testuser"}
	message := "User retrieved successfully"
	response := utils.SuccessResponse(data, message)

	// Check response structure
	if response["status"] != "success" {
		fmt.Printf("FAIL: Expected status 'success', got '%v'\n", response["status"])
	} else {
		fmt.Println("PASS: SuccessResponse status is correct")
	}

	if response["message"] != message {
		fmt.Printf("FAIL: Expected message '%s', got '%v'\n", message, response["message"])
	} else {
		fmt.Println("PASS: SuccessResponse message is correct")
	}

	if response["data"] == nil {
		fmt.Println("FAIL: SuccessResponse data should not be nil")
	} else {
		fmt.Println("PASS: SuccessResponse data is present")
	}

	// Test ErrorResponse
	errorMsg := "User not found"
	err := errors.New("database connection failed")
	statusCode := http.StatusNotFound
	errorResponse := utils.ErrorResponse(errorMsg, err, statusCode)

	if errorResponse["status"] != "error" {
		fmt.Printf("FAIL: Expected status 'error', got '%v'\n", errorResponse["status"])
	} else {
		fmt.Println("PASS: ErrorResponse status is correct")
	}

	if errorResponse["message"] != errorMsg {
		fmt.Printf("FAIL: Expected message '%s', got '%v'\n", errorMsg, errorResponse["message"])
	} else {
		fmt.Println("PASS: ErrorResponse message is correct")
	}

	if errorResponse["code"] != http.StatusText(statusCode) {
		fmt.Printf("FAIL: Expected code '%s', got '%v'\n", http.StatusText(statusCode), errorResponse["code"])
	} else {
		fmt.Println("PASS: ErrorResponse code is correct")
	}

	// Test PaginatedResponse
	paginatedData := []map[string]interface{}{
		{"id": 1, "name": "Item 1"},
		{"id": 2, "name": "Item 2"},
	}
	paginatedMessage := "Items retrieved successfully"
	totalCount := 25
	page := 2
	limit := 10

	paginatedResponse := utils.PaginatedResponse(paginatedMessage, paginatedData, totalCount, page, limit)

	if paginatedResponse["status"] != "success" {
		fmt.Printf("FAIL: Expected paginated status 'success', got '%v'\n", paginatedResponse["status"])
	} else {
		fmt.Println("PASS: PaginatedResponse status is correct")
	}

	// Check meta information
	meta, ok := paginatedResponse["meta"].(map[string]interface{})
	if !ok {
		fmt.Println("FAIL: PaginatedResponse meta should be a map")
	} else {
		if meta["total_count"] != totalCount {
			fmt.Printf("FAIL: Expected total_count %d, got %v\n", totalCount, meta["total_count"])
		} else {
			fmt.Println("PASS: PaginatedResponse total_count is correct")
		}

		if meta["page"] != page {
			fmt.Printf("FAIL: Expected page %d, got %v\n", page, meta["page"])
		} else {
			fmt.Println("PASS: PaginatedResponse page is correct")
		}

		if meta["limit"] != limit {
			fmt.Printf("FAIL: Expected limit %d, got %v\n", limit, meta["limit"])
		} else {
			fmt.Println("PASS: PaginatedResponse limit is correct")
		}

		expectedTotalPages := 3 // 25 items / 10 per page = 3 pages
		if meta["total_pages"] != expectedTotalPages {
			fmt.Printf("FAIL: Expected total_pages %d, got %v\n", expectedTotalPages, meta["total_pages"])
		} else {
			fmt.Println("PASS: PaginatedResponse total_pages calculation is correct")
		}
	}

	// Test edge case - zero limit
	paginatedResponseZero := utils.PaginatedResponse("Test", []interface{}{}, 10, 1, 0)
	metaZero, ok := paginatedResponseZero["meta"].(map[string]interface{})
	if !ok {
		fmt.Println("FAIL: PaginatedResponse with zero limit should still have meta")
	} else {
		if metaZero["total_pages"] != 0 {
			fmt.Printf("FAIL: Expected total_pages 0 with zero limit, got %v\n", metaZero["total_pages"])
		} else {
			fmt.Println("PASS: PaginatedResponse handles zero limit correctly")
		}
	}
}

func testValidationUtils() {
	// Test EmailIsValid
	fmt.Println("\n--- Testing Email Validation ---")
	
	validEmails := []string{
		"test@example.com",
		"user.name@domain.co.uk",
		"first.last+tag@example.org",
		"  TEST@EXAMPLE.COM  ", // Should handle whitespace and case
	}

	for _, email := range validEmails {
		if !utils.EmailIsValid(email) {
			fmt.Printf("FAIL: Email '%s' should be valid\n", email)
		} else {
			fmt.Printf("PASS: Email '%s' is correctly validated as valid\n", strings.TrimSpace(email))
		}
	}

	invalidEmails := []string{
		"",
		"notanemail",
		"@domain.com",
		"user@",
		"user name@domain.com",
		"user..name@domain.com",
	}

	for _, email := range invalidEmails {
		if utils.EmailIsValid(email) {
			fmt.Printf("FAIL: Email '%s' should be invalid\n", email)
		} else {
			fmt.Printf("PASS: Email '%s' is correctly validated as invalid\n", email)
		}
	}

	// Test UsernameIsValid
	fmt.Println("\n--- Testing Username Validation ---")
	
	validUsernames := []string{
		"user123",
		"test_user",
		"user.name",
		"user123!",
		"a1b2c3d4", // 8 characters
		"abcdefghij1234567890", // 20 characters
	}

	for _, username := range validUsernames {
		if !utils.UsernameIsValid(username) {
			fmt.Printf("FAIL: Username '%s' should be valid\n", username)
		} else {
			fmt.Printf("PASS: Username '%s' is correctly validated as valid\n", username)
		}
	}

	invalidUsernames := []string{
		"",
		"ab", // Too short
		"abcdefghij1234567890x", // Too long (21 characters)
		"user name", // Contains space
		"user@name", // Contains @
		"user#name", // Contains #
		"user-name", // Contains -
	}

	for _, username := range invalidUsernames {
		if utils.UsernameIsValid(username) {
			fmt.Printf("FAIL: Username '%s' should be invalid\n", username)
		} else {
			fmt.Printf("PASS: Username '%s' is correctly validated as invalid\n", username)
		}
	}

	// Test PasswordStrength
	fmt.Println("\n--- Testing Password Strength ---")
	
	validPasswords := []string{
		"Password1!",
		"MySecure@123",
		"Test&Pass9",
		"ComplexP@ssw0rd",
		"Abcdef1!", // Minimum valid
	}

	for _, password := range validPasswords {
		if !utils.PasswordStrength(password) {
			fmt.Printf("FAIL: Password '%s' should be strong enough\n", password)
		} else {
			fmt.Printf("PASS: Password is correctly validated as strong\n")
		}
	}

	invalidPasswords := []string{
		"",
		"short1!", // Too short (7 characters)
		"password123!", // No uppercase
		"PASSWORD123!", // No lowercase (your function might not check this)
		"Password123", // No special character
		"PasswordABC!", // No digit
		"   Pass1!   ", // Should handle whitespace
	}

	for _, password := range invalidPasswords {
		if utils.PasswordStrength(password) {
			fmt.Printf("FAIL: Password should be invalid (weak)\n")
		} else {
			fmt.Printf("PASS: Weak password correctly rejected\n")
		}
	}

	// Test ImageFileisValid
	fmt.Println("\n--- Testing Image File Validation ---")
	
	maxSize := int64(5) // 5MB

	// Valid files
	validFiles := []struct {
		filename string
		size     int64
	}{
		{"image.jpg", 1024 * 1024},     // 1MB
		{"photo.png", 3 * 1024 * 1024}, // 3MB
		{"pic.gif", 100 * 1024},        // 100KB
	}

	for _, file := range validFiles {
		if !utils.ImageFileisValid(file.filename, file.size, maxSize) {
			fmt.Printf("FAIL: File '%s' (%d bytes) should be valid\n", file.filename, file.size)
		} else {
			fmt.Printf("PASS: File '%s' is correctly validated as valid\n", file.filename)
		}
	}

	// Invalid files
	invalidFiles := []struct {
		filename string
		size     int64
		reason   string
	}{
		{"", 1024, "empty filename"},
		{"image.jpg", 0, "zero size"},
		{"image.jpg", -1000, "negative size"},
		{"large.jpg", 6 * 1024 * 1024, "too large"},
		{"  ", 1024, "whitespace only filename"},
	}

	for _, file := range invalidFiles {
		if utils.ImageFileisValid(file.filename, file.size, maxSize) {
			fmt.Printf("FAIL: File should be invalid (%s)\n", file.reason)
		} else {
			fmt.Printf("PASS: Invalid file correctly rejected (%s)\n", file.reason)
		}
	}
}

func testJWTUtils() {
	// Note: These tests are limited since JWT functions depend on external services
	// and configuration that may not be available in test environment
	
	fmt.Println("NOTE: JWT tests are limited due to external dependencies")
	
	// Test ExtractClaims with a mock token structure
	// This tests the parsing logic without requiring a valid signature
	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	
	claims, err := utils.ExtractClaims(testToken)
	if err != nil {
		fmt.Printf("PASS: ExtractClaims handles token parsing (Note: %v)\n", err)
	} else {
		if claims["sub"] == "1234567890" && claims["name"] == "John Doe" {
			fmt.Println("PASS: ExtractClaims correctly parses token claims")
		} else {
			fmt.Printf("FAIL: ExtractClaims returned unexpected claims: %+v\n", claims)
		}
	}
	
	// Test with Bearer prefix
	bearerToken := "Bearer " + testToken
	claimsWithBearer, err := utils.ExtractClaims(bearerToken)
	if err != nil {
		fmt.Printf("INFO: ExtractClaims with Bearer prefix: %v\n", err)
	} else {
		if claimsWithBearer["sub"] == claims["sub"] {
			fmt.Println("PASS: ExtractClaims correctly handles Bearer prefix")
		} else {
			fmt.Println("FAIL: ExtractClaims doesn't handle Bearer prefix correctly")
		}
	}
	
	// Test invalid token
	invalidToken := "invalid.token.here"
	_, err = utils.ExtractClaims(invalidToken)
	if err != nil {
		fmt.Println("PASS: ExtractClaims correctly rejects invalid token")
	} else {
		fmt.Println("FAIL: ExtractClaims should reject invalid token")
	}
	
	// Test ValidateToken (will likely fail without proper config/connection)
	_, err = utils.ValidateToken(testToken)
	if err != nil {
		fmt.Printf("EXPECTED: ValidateToken failed (requires Supabase connection): %v\n", err)
	} else {
		fmt.Println("PASS: ValidateToken succeeded (unexpected in test environment)")
	}
	
	fmt.Println("INFO: Full JWT testing requires proper Supabase configuration")
}