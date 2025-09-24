// 3. PostWithUser struct for feed responses
// 4. Validation for post creation

package models

import (
	"fmt"
	"strings"
	"time"

)


type Post struct {
	ID              int 	`json:"id" validate:"required"`
	UserID          int 	`json:"user_id" validate:"required"`
	Title           string	`json:"title" validate:"omitempty,max=100"`
	Description     string	`json:"description" validate:"omitempty,max=400"`
	ImageURL        string 	`json:"image_url" validate:"omitempty,url"`
	Recipe          Recipe 	`json:"recipe" validate:"required,dive"`
	LikesCount 	    int 	`json:"likes_count" validate:"required,min=0"`
	CommentsCount 	int 	`json:"comments_count" validate:"required,min=0"`
	CreatedAt 	    string 	`json:"created_at" validate:"required"`
	UpdatedAt 		string 	`json:"updated_at" validate:"required"`
}
type Ingredients struct {
	Name     string             `json:"name" validate:"required"`
	Quantity string             `json:"quantity" validate:"required"`
	Grams    int                `json:"grams" validate:"omitempty,min=0"`
}

type Recipe struct {
	Ingredients  []Ingredients `json:"ingredients" validate:"required,min=1,dive,required"`
	Instructions []string      `json:"instructions" validate:"required,min=1,dive,required"`
}

type PostWithUser struct {
	Post
	User User `json:"user" validate:"required,dive"`
}

func (p *Post) Validate() error {
	return validate.Struct(p)
}

func (p *Post) TimeFormat() string{
	t, err := time.Parse(time.RFC3339, p.CreatedAt)
	if err != nil {
		return ""
	}
	return t.Format("02 Jan 2006 15:04")
}