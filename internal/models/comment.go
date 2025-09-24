package models

import ()

type Comment struct {
	ID     int `json:"id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
	PostID int `json:"post_id" validate:"required"`
	Content string `json:"content" validate:"required min=0,max=250"`
	CreatedAt string `json:"created_at" validate:"required"`
}

type CommentWithUser struct {
	Comment
	User User `json:"user" validate:"required,dive"`
}

func (v *Comment) Validate() error {
	return validate.Struct(v)
}