package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"

	// NOTE: You will need to replace "your_module_name" with your Go project's actual module name
	// (the one from your go.mod file). This is so this main program can find and import
	// the 'models' package where your User struct lives.
	"feast-friends-api/internal/models"
)

// Helper function to create a valid user for testing.
// It now returns a *models.User from the imported package.
func newValidUser() *models.User {
	return &models.User{
		ID:             1,
		Email:          "test@example.com",
		Username:       "testuser",
		FullName:       "Test User",
		Bio:            "This is a test bio.",
		AvatarURL:      "http://example.com/avatar.png",
		FollowersCount: 10,
		FollowingCount: 5,
		PostsCount:     20,
		CreatedAt:      "2025-09-23T15:00:00Z", // RFC3339 format
	}
}

// main is the entry point for the program.
// All the checks from the original test file have been moved here.
func main() {
	fmt.Println("--- Running User Model Checks ---")

	// --- Check User Validation ---
	fmt.Println("\n--- Testing User Validation ---")
	// IMPORTANT: We create the validator here in main.
	// We can't call the user.Validate() method directly because it relies on a private
	// variable in the 'models' package. Instead, we perform the validation from here.
	validate := validator.New()

	// Scenario 1: A valid user should pass validation.
	validUser := newValidUser()
	err := validate.Struct(validUser)
	if err != nil {
		fmt.Printf("FAIL: Expected a valid user to pass validation, but got error: %v\n", err)
	} else {
		fmt.Println("PASS: Valid user passed validation.")
	}

	// Scenario 2: An invalid user should fail validation.
	invalidUser := newValidUser()
	invalidUser.Email = "not-an-email" // Make the email invalid
	err = validate.Struct(invalidUser)
	if err == nil {
		fmt.Println("FAIL: Expected an invalid user to fail validation, but it passed.")
	} else {
		fmt.Println("PASS: Invalid user failed validation as expected.")
	}

	// --- Check DisplayName() ---
	fmt.Println("\n--- Testing DisplayName ---")
	userWithFullName := newValidUser()
	expectedName := "Test User"
	if got := userWithFullName.DisplayName(); got != expectedName {
		fmt.Printf("FAIL: Expected DisplayName to be '%s', but got '%s'\n", expectedName, got)
	} else {
		fmt.Printf("PASS: DisplayName with full name is correct ('%s').\n", got)
	}

	userWithoutFullName := newValidUser()
	userWithoutFullName.FullName = "  " // Empty string with spaces
	expectedUsername := "testuser"
	if got := userWithoutFullName.DisplayName(); got != expectedUsername {
		fmt.Printf("FAIL: Expected DisplayName to be '%s', but got '%s'\n", expectedUsername, got)
	} else {
		fmt.Printf("PASS: DisplayName without full name falls back to username ('%s').\n", got)
	}

	// --- Check TimeFormat() ---
	fmt.Println("\n--- Testing TimeFormat ---")
	userForTime := newValidUser()
	expectedTime := "23 Sep 2025 15:00"
	if got := userForTime.TimeFormat(); got != expectedTime {
		fmt.Printf("FAIL: Expected TimeFormat to be '%s', but got '%s'\n", expectedTime, got)
	} else {
		fmt.Printf("PASS: TimeFormat with valid date is correct ('%s').\n", got)
	}

	userForTime.CreatedAt = "invalid-date"
	if got := userForTime.TimeFormat(); got != "invalid-date" {
		fmt.Printf("FAIL: Expected TimeFormat to return the original invalid string, but got '%s'\n", got)
	} else {
		fmt.Println("PASS: TimeFormat with invalid date returned original string as expected.")
	}

	// --- Check PublicProfile() ---
	fmt.Println("\n--- Testing PublicProfile ---")
	userForProfile := newValidUser()
	profile := userForProfile.PublicProfile()

	if profile["UserName"] != "Test User" {
		fmt.Printf("FAIL: Expected UserName in public profile to be 'Test User', got '%v'\n", profile["UserName"])
	} else {
		fmt.Println("PASS: Public profile 'UserName' is correct.")
	}
	if profile["Posts"] != 20 {
		fmt.Printf("FAIL: Expected Posts in public profile to be 20, got '%v'\n", profile["Posts"])
	} else {
		fmt.Println("PASS: Public profile 'Posts' is correct.")
	}

	if _, exists := profile["id"]; exists {
		fmt.Println("FAIL: Public profile should not contain the user ID.")
	} else {
		fmt.Println("PASS: Public profile does not contain sensitive 'id' field.")
	}
}

