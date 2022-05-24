package football_world_cup_score_board

import (
	"errors"
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
	"time"
)

var (
	ErrGameAlreadyExist = errors.New("game already exist")
	ErrGameNotFound     = errors.New("game not found")
)

type CreateGameRequest struct {
	HomeTeamName string
	AwayTeamName string
}

func CreateGameRequestToGameEntity(request CreateGameRequest) entity.Game {
	now := time.Now()
	return entity.Game{
		StartedAt:     now,
		UpdatedAt:     now,
		Status:        entity.GameStatusActive,
		HomeTeamName:  request.HomeTeamName,
		AwayTeamName:  request.AwayTeamName,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
	}
}

type UpdateGameRequest struct {
	HomeTeamName  string
	HomeTeamScore int
	AwayTeamName  string
	AwayTeamScore int
}

type FinishGameRequest struct {
	HomeTeamName string
	AwayTeamName string
}

type SummaryResponse struct {
	Items []SummaryGameItem
}

type SummaryGameItem struct {
	HomeTeamName  string
	HomeTeamScore int
	AwayTeamName  string
	AwayTeamScore int
}
