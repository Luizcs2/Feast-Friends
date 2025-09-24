package main

import (
	"fmt"
	"time"
	"github.com/go-playground/validator/v10"

	"feast-friends-api/internal/models"
)

// Helper function to create a valid user for testing
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
		CreatedAt:      time.Now(),
	}
}

// Helper function to create a valid post for testing
func newValidPost() *models.Post {
	return &models.Post{
		ID:            1,
		UserID:        1,
		Title:         "Test Recipe Post",
		Description:   "A delicious test recipe",
		ImageURL:      "http://example.com/recipe.jpg",
		Recipe: models.Recipe{
			Ingredients: []models.Ingredients{
				{Name: "Flour", Quantity: "2 cups", Grams: 240},
				{Name: "Sugar", Quantity: "1 cup", Grams: 200},
			},
			Instructions: []string{
				"Mix ingredients together",
				"Bake for 30 minutes",
			},
		},
		LikesCount:    5,
		CommentsCount: 3,
		CreatedAt:     time.Now(),
	}
}

// Helper function to create a valid comment for testing
func newValidComment() *models.Comment {
	return &models.Comment{
		ID:        1,
		UserID:    1,
		PostID:    1,
		Content:   "Great recipe!",
		CreatedAt: time.Now(),
	}
}

// Helper function to create a valid event for testing
func newValidEvent() *models.Event {
	return &models.Event{
		ID:               1,
		CreatorID:        1,
		Title:            "Food Festival",
		Description:      "A great food event for everyone",
		Location:         "Central Park",
		EventDate:        time.Now().Add(24 * time.Hour), // Future date
		MaxAttendees:     50,
		CurrentAttendees: 10,
		ImageURL:         "http://example.com/event.jpg",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

// Helper function to create a valid conversation for testing
func newValidConversation() *models.Conversation {
	return &models.Conversation{
		ID:            1,
		User1ID:       1,
		LastMessageAt: time.Now(),
		CreatedAt:     time.Now(),
	}
}

// Helper function to create a valid message for testing
func newValidMessage() *models.Message {
	return &models.Message{
		ID:             1,
		ConversationID: 1,
		SenderID:       1,
		Content:        "Hello there!",
		ReadAt:         time.Now(),
		MessageType:    "text",
		CreatedAt:      time.Now(),
	}
}

func main() {
	fmt.Println("=== Running Complete Models Test Suite ===")
	fmt.Println("NOTE: Some validation tests might fail due to model validation tag issues.")
	fmt.Println("Check the console output for specific fixes needed in your models.")
	validate := validator.New()

	// --- USER MODEL TESTS ---
	fmt.Println("\n--- Testing User Model ---")
	testUserModel(validate)

	// --- POST MODEL TESTS ---
	fmt.Println("\n--- Testing Post Model ---")
	testPostModel(validate)

	// --- COMMENT MODEL TESTS ---
	fmt.Println("\n--- Testing Comment Model ---")
	testCommentModel(validate)

	// --- EVENT MODEL TESTS ---
	fmt.Println("\n--- Testing Event Model ---")
	testEventModel(validate)

	// --- CONVERSATION MODEL TESTS ---
	fmt.Println("\n--- Testing Conversation Model ---")
	testConversationModel(validate)

	// --- MESSAGE MODEL TESTS ---
	fmt.Println("\n--- Testing Message Model ---")
	testMessageModel(validate)

	fmt.Println("\n=== Test Suite Complete ===")
}

func testUserModel(validate *validator.Validate) {
	// Valid user test
	validUser := newValidUser()
	err := validate.Struct(validUser)
	if err != nil {
		fmt.Printf("FAIL: Valid user failed validation: %v\n", err)
	} else {
		fmt.Println("PASS: Valid user passed validation")
	}

	// Invalid email test
	invalidUser := newValidUser()
	invalidUser.Email = "not-an-email"
	err = validate.Struct(invalidUser)
	if err == nil {
		fmt.Println("FAIL: Invalid email should fail validation")
	} else {
		fmt.Println("PASS: Invalid email failed validation as expected")
	}

	// DisplayName test with full name
	if got := validUser.DisplayName(); got != "Test User" {
		fmt.Printf("FAIL: Expected 'Test User', got '%s'\n", got)
	} else {
		fmt.Println("PASS: DisplayName with full name works correctly")
	}

	// DisplayName test without full name
	validUser.FullName = "  "
	if got := validUser.DisplayName(); got != "testuser" {
		fmt.Printf("FAIL: Expected 'testuser', got '%s'\n", got)
	} else {
		fmt.Println("PASS: DisplayName falls back to username correctly")
	}

	// TimeFormat test
	if timeStr := validUser.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: TimeFormat works: %s\n", timeStr)
	}

	// PublicProfile test
	profile := validUser.PublicProfile()
	if profile["UserName"] != "testuser" {
		fmt.Printf("FAIL: Expected 'testuser' in public profile, got '%v'\n", profile["UserName"])
	} else {
		fmt.Println("PASS: PublicProfile UserName is correct")
	}
}

func testPostModel(validate *validator.Validate) {
	// Valid post test
	validPost := newValidPost()
	err := validate.Struct(validPost)
	if err != nil {
		fmt.Printf("FAIL: Valid post failed validation: %v\n", err)
		// Continue with other tests even if validation fails due to model issues
	} else {
		fmt.Println("PASS: Valid post passed validation")
	}

	// Test individual recipe validation (since Post validation might fail due to model issues)
	validRecipe := &models.Recipe{
		Ingredients: []models.Ingredients{
			{Name: "Flour", Quantity: "2 cups", Grams: 240},
		},
		Instructions: []string{"Mix well"},
	}
	err = validate.Struct(validRecipe)
	if err != nil {
		fmt.Printf("FAIL: Valid recipe failed validation: %v\n", err)
	} else {
		fmt.Println("PASS: Valid recipe passed validation")
	}

	// Invalid recipe test - empty ingredients
	invalidRecipe := &models.Recipe{
		Ingredients:  []models.Ingredients{},
		Instructions: []string{"Mix well"},
	}
	err = validate.Struct(invalidRecipe)
	if err == nil {
		fmt.Println("FAIL: Recipe with empty ingredients should fail validation")
	} else {
		fmt.Println("PASS: Recipe with empty ingredients failed validation as expected")
	}

	// TimeFormat test
	if timeStr := validPost.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: Post TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: Post TimeFormat works: %s\n", timeStr)
	}

	// Test recipe structure
	if len(validPost.Recipe.Ingredients) != 2 {
		fmt.Printf("FAIL: Expected 2 ingredients, got %d\n", len(validPost.Recipe.Ingredients))
	} else {
		fmt.Println("PASS: Recipe ingredients count is correct")
	}
}

func testCommentModel(validate *validator.Validate) {
	// Valid comment test
	validComment := newValidComment()
	err := validate.Struct(validComment)
	if err != nil {
		fmt.Printf("FAIL: Valid comment failed validation: %v\n", err)
	} else {
		fmt.Println("PASS: Valid comment passed validation")
	}

	// Invalid comment test - content too long
	invalidComment := newValidComment()
	invalidComment.Content = string(make([]byte, 251)) // More than 250 chars
	for i := range invalidComment.Content {
		invalidComment.Content = invalidComment.Content[:i] + "a" + invalidComment.Content[i+1:]
	}
	err = validate.Struct(invalidComment)
	if err == nil {
		fmt.Println("FAIL: Comment with content > 250 chars should fail validation")
	} else {
		fmt.Println("PASS: Comment with long content failed validation as expected")
	}

	// TimeFormat test
	if timeStr := validComment.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: Comment TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: Comment TimeFormat works: %s\n", timeStr)
	}
}

func testEventModel(validate *validator.Validate) {
	// Valid event test
	validEvent := newValidEvent()
	err := validate.Struct(validEvent)
	if err != nil {
		fmt.Printf("FAIL: Valid event failed validation: %v\n", err)
	} else {
		fmt.Println("PASS: Valid event passed validation")
	}

	// Invalid event test - title too short
	invalidEvent := newValidEvent()
	invalidEvent.Title = "Hi" // Less than 3 characters
	err = validate.Struct(invalidEvent)
	if err == nil {
		fmt.Println("FAIL: Event with short title should fail validation")
	} else {
		fmt.Println("PASS: Event with short title failed validation as expected")
	}

	// TimeFormat test
	if timeStr := validEvent.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: Event TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: Event TimeFormat works: %s\n", timeStr)
	}

	// Test EventRSVP
	eventRSVP := &models.EventRSVP{
		Event:     *validEvent,
		User:      *newValidUser(),
		CreatedAt: time.Now(),
		Statues:   "going", // Note: There's a typo in your model - should be "Status"
	}

	if timeStr := eventRSVP.RSVPTimeFormat(); timeStr == "" {
		fmt.Println("FAIL: EventRSVP RSVPTimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: EventRSVP RSVPTimeFormat works: %s\n", timeStr)
	}
}

func testConversationModel(validate *validator.Validate) {
	// Valid conversation test
	validConversation := newValidConversation()
	err := validate.Struct(validConversation)
	if err != nil {
		fmt.Printf("FAIL: Valid conversation failed validation: %v\n", err)
	} else {
		fmt.Println("PASS: Valid conversation passed validation")
	}

	// TimeFormat test
	if timeStr := validConversation.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: Conversation TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: Conversation TimeFormat works: %s\n", timeStr)
	}

	// LastMessageTimeFormat test
	if timeStr := validConversation.LastMessageTimeFormat(); timeStr == "" {
		fmt.Println("FAIL: Conversation LastMessageTimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: Conversation LastMessageTimeFormat works: %s\n", timeStr)
	}

	// Test ConversationWithUser
	conversationWithUser := &models.ConversationWithUser{
		Conversation: *validConversation,
		OtherUser:    *newValidUser(),
		LastMessage:  newValidMessage(),
	}

	if timeStr := conversationWithUser.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: ConversationWithUser TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: ConversationWithUser TimeFormat works: %s\n", timeStr)
	}
}

func testMessageModel(validate *validator.Validate) {
	// Valid message test
	validMessage := newValidMessage()
	err := validate.Struct(validMessage)
	if err != nil {
		fmt.Printf("FAIL: Valid message failed validation: %v\n", err)
	} else {
		fmt.Println("PASS: Valid message passed validation")
	}

	// Invalid message test - content too long
	invalidMessage := newValidMessage()
	invalidMessage.Content = string(make([]byte, 101)) // More than 100 chars
	for i := range invalidMessage.Content {
		invalidMessage.Content = invalidMessage.Content[:i] + "a" + invalidMessage.Content[i+1:]
	}
	err = validate.Struct(invalidMessage)
	if err == nil {
		fmt.Println("FAIL: Message with content > 100 chars should fail validation")
	} else {
		fmt.Println("PASS: Message with long content failed validation as expected")
	}

	// Invalid message type test
	invalidMessage2 := newValidMessage()
	invalidMessage2.MessageType = "invalid_type"
	err = validate.Struct(invalidMessage2)
	if err == nil {
		fmt.Println("FAIL: Message with invalid type should fail validation")
	} else {
		fmt.Println("PASS: Message with invalid type failed validation as expected")
	}

	// TimeFormat test
	if timeStr := validMessage.TimeFormat(); timeStr == "" {
		fmt.Println("FAIL: Message TimeFormat returned empty string")
	} else {
		fmt.Printf("PASS: Message TimeFormat works: %s\n", timeStr)
	}
}