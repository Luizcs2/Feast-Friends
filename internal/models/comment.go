package models
import(
	"feast-friends-api/pkg/helpers"
)

// Comment represents a comment made by a user on a post
type Comment struct {
	ID        int    `json:"id" validate:"required"`      // Unique identifier for the comment
	UserID    int    `json:"user_id" validate:"required"` // ID of the user who made the comment
	PostID    int    `json:"post_id" validate:"required"` // ID of the post being commented on
	Content   string `json:"content" validate:"required min=0,max=250"` // The comment text, limited to 250 characters
	CreatedAt string `json:"created_at" validate:"required"`           // Timestamp when the comment was created
}

// CommentWithUser is a composite structure that embeds a Comment
// and includes the associated User information
type CommentWithUser struct {
	Comment           // Embedded Comment struct
	User    User `json:"user" validate:"required,dive"` // The user who made the comment
}

// Validate checks if the Comment struct fields meet the validation rules
// defined in the struct tags using the validator package
func (x *Comment) Validate() error {
	return helpers.ValidateStruct(x)
}

//timeFormat wil convert the created at timestamp to a human readable formtat 
//it will return a date in the format "02 Jan 2002 15:04"
func (x *Comment) TimeFormat() string {
	return helpers.FormatTime(x.CreatedAt)
}
