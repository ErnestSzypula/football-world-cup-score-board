package storage

import (
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
	"github.com/samber/lo"
	"sort"
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

func (s *InMemoryStorage) UpdateGame(index int, game entity.Game) {
	s.games[index] = game
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

func (s *InMemoryStorage) GetActiveGamesSortedByScore() []entity.Game {
	results := lo.Filter(s.games, func(g entity.Game, _ int) bool {
		return g.Status == entity.GameStatusActive
	})

	sort.Sort(byScore(results))

	return results
}

type byScore []entity.Game

func (b byScore) Len() int      { return len(b) }
func (b byScore) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byScore) Less(i, j int) bool {
	leftScore := b[i].HomeTeamScore + b[i].AwayTeamScore
	rightScore := b[j].HomeTeamScore + b[j].AwayTeamScore

	if leftScore == rightScore {
		return b[i].UpdatedAt.After(b[j].UpdatedAt)
	}

	return leftScore > rightScore
}
