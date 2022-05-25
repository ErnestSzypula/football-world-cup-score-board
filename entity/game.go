package entity

import "time"

type GameStatus string

const (
	GameStatusActive   GameStatus = "ACTIVE"
	GameStatusFinished GameStatus = "FINISHED"
)

type Game struct {
	StartedAt    time.Time
	UpdatedAt    time.Time
	FinishedAt   *time.Time
	Status       GameStatus
	HomeTeamName string
	AwayTeamName string
	// At production solution I will consider store Score as separate entity with saving subsequent changes
	HomeTeamScore int
	AwayTeamScore int
}
