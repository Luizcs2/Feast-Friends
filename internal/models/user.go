// user.go contains the User struct that represents a user in the database
// it includes fields like ID, Email, Username, FullName, Bio, AvatarURL, FollowersCount, FollowingCount, PostsCount and CreatedAt
// it also includes methods to validate the struct fields using go-playground/validator package
// and methods to format the CreatedAt field to a more readable format
// and methods to display the user's full name or username if full name is not available
// and methods to return a public profile of the user with selected fields

package models

import (
	"fmt"
	"strings"
	"feast-friends-api/pkg/helpers"
	"time"
)



// the struct that represents a user in the database
type User struct {
	ID             int    `json:"id" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Username       string `json:"username" validate:"required,alphanum,min=3,max=20"`
	FullName       string `json:"full_name" validate:"required,min=2,max=50"`
	Bio            string `json:"bio" validate:"omitempty,max=200"`
	AvatarURL      string `json:"avatar_url" validate:"omitempty,url"`
	FollowersCount int    `json:"followers_count" validate:"required,min=0"`
	FollowingCount int    `json:"following_count" validate:"required,min=0"`
	PostsCount     int    `json:"posts_count" validate:"required,min=0"`
	CreatedAt      time.Time `json:"created_at" validate:"required"`
}


// function that displays the users full name if available otherwise username
func (u *User) DisplayName() string {
	if strings.TrimSpace(u.FullName) != "" {
		return u.FullName
	}
	return u.Username
}

// function that formats the CreatedAt field to a more readable format
func (u *User) TimeFormat() string {
	return helpers.FormatTime(u.CreatedAt)
}

// use this func to validate user struct fields
func (u *User) Validate() error {
	return helpers.ValidateStruct(u)
}


// function that returns a string representation of the user profile for development and debugging
func (u User) UserProfile() string {
	return fmt.Sprintf(
		"User  Name: %s Created: %s Posts: %v Followers: %v Following: %v",
		u.DisplayName(),
		u.TimeFormat(),
		u.PostsCount,
		u.FollowersCount,
		u.FollowingCount,
	)
}

// function that returns a public profile of the user with selected fields
func (u *User) PublicProfile() map[string]interface{} {
	return map[string]interface{}{
		"UserName":   u.DisplayName(),
		"Created":    u.TimeFormat(),
		"Posts":      u.PostsCount,
		"Followers":  u.FollowersCount,
		"Following":  u.FollowingCount,
		"bio":        u.Bio,
		"avatar_url": u.AvatarURL,
	}
}
