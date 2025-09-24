package models

import (
	"feast-friends-api/pkg/helpers"
	"time"
)

// Conversation represents a chat between two users.
type Conversation struct {
	ID            int       `json:"id" validate:"required"`
	User1ID       int       `json:"user1_id" validate:"required"`
	LastMessageAt time.Time `json:"last_message_at" validate:"omitempty"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
}

// Message represents a single message in a conversation.
type Message struct {
	ID             int       `json:"id" validate:"required"`
	ConversationID int       `json:"conversation_id" validate:"required"` // Use int for foreign key
	SenderID       int       `json:"sender_id" validate:"required"`       // Fixed typo
	Content        string    `json:"content" validate:"required,min=1,max=100"` // Fixed JSON tag
	ReadAt         time.Time `json:"read_at" validate:"required"`
	MessageType    string    `json:"message_type" validate:"required,oneof=text image video"` // Use string values
	CreatedAt      time.Time `json:"created_at" validate:"required"`
}

// ConversationWithUser combines a conversation with the other user and last message.
type ConversationWithUser struct {
	Conversation
	OtherUser   User     `json:"other_user" validate:"required,dive"`
	LastMessage *Message `json:"last_message,omitempty"`
}

// Validate validates the Message struct fields.
func (x *Message) Validate() error {
	return helpers.ValidateStruct(x)
}

// TimeFormat returns the formatted creation time of the message.
func (x *Message) TimeFormat() string {
	return helpers.FormatTime(x.CreatedAt)
}

// TimeFormat returns the formatted creation time of the conversation.
func (c *Conversation) TimeFormat() string {
	return helpers.FormatTime(c.CreatedAt)
}

// LastMessageTimeFormat returns the formatted time of the last message in the conversation.
func (c *Conversation) LastMessageTimeFormat() string {
	return helpers.FormatTime(c.LastMessageAt)
}

// TimeFormat returns the formatted creation time for ConversationWithUser.
func (c *ConversationWithUser) TimeFormat() string {
	return helpers.FormatTime(c.CreatedAt)
}