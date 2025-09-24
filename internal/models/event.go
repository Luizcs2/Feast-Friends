package models

import (
	"feast-friends-api/pkg/helpers"
	"time"
)

// Event represents an event created by a user.
// It contains all the necessary details about the event.
type Event struct {
	ID               int       `json:"id" validate:"required"`                        
	CreatorID        int       `json:"creator_id" validate:"required"`                
	Title            string    `json:"title" validate:"required,min=3,max=20"`        
	Description      string    `json:"description" validate:"required,max=500"`       
	Location         string    `json:"location" validate:"required"`                  
	EventDate        time.Time `json:"event_date" validate:"required,future"`         
	MaxAttendees     int       `json:"max_attendees" validate:"required,min=1"`       // Maximum number of attendees allowed
	CurrentAttendees int       `json:"current_attendees"`                             // Current number of attendees
	ImageURL         string    `json:"image_url"`                                     // Optional image URL for the event
	CreatedAt        time.Time `json:"created_at"`                                    // Timestamp when the event was created
	UpdatedAt        time.Time `json:"updated_at"`                                    // Timestamp when the event was last updated
}

// EventRSVP represents a user's RSVP to an event.
// It embeds the Event struct and adds RSVP-specific fields.
type EventRSVP struct {
	Event
	User      User      `json:"user" validate:"required,dive"`                        // The user who RSVP'd to the event
	CreatedAt time.Time `json:"created_at"`                                           // Timestamp when the RSVP was created
	Statues   string    `json:"statues" validate:"required,oneof=going maybe cancelled"` // Status of the RSVP (going, maybe, cancelled)
}

// Validate validates the Event struct fields using the helpers package.
func (x *Event) Validate() error {
	return helpers.ValidateStruct(x)
}

// TimeFormat returns the formatted creation time of the event.
func (x *Event) TimeFormat() string {
	return helpers.FormatTime(x.CreatedAt)
}	