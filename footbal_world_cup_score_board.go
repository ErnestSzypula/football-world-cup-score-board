package football_world_cup_score_board

import (
	"errors"
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
	"time"
)

type CreateGame struct {
	HomeTeam string
	AwayTeam string
}

var (
	ErrGameAlreadyExist = errors.New("game already exist")
)

func CreateGameToGameEntity(game CreateGame) entity.Game {
	return entity.Game{
		StartedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Status:        entity.GameStatusActive,
		HomeTeamName:  game.HomeTeam,
		AwayTeamName:  game.AwayTeam,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
	}
}
