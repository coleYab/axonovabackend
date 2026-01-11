package entity

import "time"

type Event struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Picture      string    `json:"picture"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	StartTime    string    `json:"startTime"`
	DurationMin  int       `json:"minDuration"`
	Price        int       `json:"price"`
	MaxAttendees int       `json:"maxAttendees"`
	IsOnline     bool      `json:"isOnline"`
	Platform     string    `json:"platform"`
	MeetingLink  string    `json:"meetingLink"`
	Tags         []string  `json:"tags"`
}
