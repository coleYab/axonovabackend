package dto

import "time"

type CreateEventDTO struct {
	Title        string    `json:"title" binding:"required"`
	Picture      string    `json:"picture" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Date         time.Time `json:"date" binding:"required"`
	StartTime    string    `json:"startTime" binding:"required"`
	DurationMin  int       `json:"minDuration" binding:"required"`
	Price        int       `json:"price" binding:"required,min=1"`
	MaxAttendees int       `json:"maxAttendees" binding:"required"`
	IsOnline     bool      `json:"isOnline" binding:"required"`
	Platform     string    `json:"platform" binding:"required,oneof=zoom google_meet"`
	MeetingLink  string    `json:"meetingLink" binding:"required,url"`
	Tags         []string  `json:"tags"`
}

type BookEventDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`

	EventID string `json:"eventId" binding:"required"`

	Quantity int `json:"quantity" binding:"required,min=1"`

	Phone              string `json:"phone" binding:"required"`
	SendReminderEmails bool   `json:"sendReminderEmails"`
}
