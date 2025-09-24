package models

import (
	"feast-friends-api/pkg/helpers"
	"time"
)

// Post represents a social media post containing recipe information
type Post struct {
	ID            int    	`json:"id" validate:"required"`              // Unique identifier for the post
	UserID        int    	`json:"user_id" validate:"required"`        // ID of the user who created the post
	Title         string 	`json:"title" validate:"omitempty,max=100"` // Optional title of the post, max 100 chars
	Description   string 	`json:"description" validate:"omitempty,max=400"` // Optional description, max 400 chars
	ImageURL      string 	`json:"image_url" validate:"omitempty,url"`      // Optional URL to post's image
	Recipe        Recipe 	`json:"recipe" validate:"required,dive"`          // Recipe details, required
	LikesCount    int    	`json:"likes_count" validate:"required,min=0"`   // Number of likes, must be non-negative
	CommentsCount int    	`json:"comments_count" validate:"required,min=0"` // Number of comments, must be non-negative
	CreatedAt     time.Time `json:"created_at" validate:"required"`          // Timestamp of post creation in RFC3339 format
}

// Ingredients represents a single ingredient in a recipe
type Ingredients struct {
	Name    	 string		`json:"name" validate:"required"`     // Name of the ingredient
	Quantity	 string		`json:"quantity" validate:"required"` // Amount needed (e.g., "2 cups")
	Grams    	int    		`json:"grams" validate:"omitempty,min=0"` // Optional weight in grams
}	

// Recipe contains the full recipe information
type Recipe struct {
	Ingredients  []Ingredients `json:"ingredients" validate:"required,min=1,dive,required"`  // List of ingredients, at least one required
	Instructions []string      `json:"instructions" validate:"required,min=1,dive,required"` // List of steps, at least one required
}

// PostWithUser extends Post to include user information
type PostWithUser struct {
	Post                // Embedded Post struct
	User User `json:"user" validate:"required,dive"` // User who created the post
}

// Validate performs validation on the Post struct using the validator package
func (p *Post) Validate() error {
	return helpers.ValidateStruct(p)
}

// TimeFormat converts the CreatedAt timestamp to a human-readable format
// Returns date in format "02 Jan 2006 15:04"
func (p *Post) TimeFormat() string {
	return helpers.FormatTime(p.CreatedAt)
}
