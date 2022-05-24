package storage

import (
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
	"github.com/samber/lo"
)

type InMemoryStorage struct {
	games []entity.Game
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		games: []entity.Game{},
	}
}

func (s *InMemoryStorage) AddGame(game entity.Game) {
	s.games = append(s.games, game)
}

func (s *InMemoryStorage) GetActiveGameByTeams(homeTeam, awayTeam string) (*entity.Game, int) {
	game, index, exist := lo.FindIndexOf(s.games, func(g entity.Game) bool {
		return g.Status == entity.GameStatusActive && g.HomeTeamName == homeTeam && g.AwayTeamName == awayTeam
	})

	if !exist {
		return nil, -1
	}

	return &game, index
}
