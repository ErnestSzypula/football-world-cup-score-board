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

func (s *InMemoryStorage) GetGameByTeams(homeTeam, awayTeam string) *entity.Game {
	game, exist := lo.Find(s.games, func(g entity.Game) bool {
		return g.Status == entity.GameStatusActive && g.HomeTeamName == homeTeam && g.AwayTeamName == awayTeam
	})

	if !exist {
		return nil
	}

	return &game
}
