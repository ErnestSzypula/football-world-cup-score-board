package service

import (
	football_world_cup_score_board "github.com/ErnestSzypula/football-world-cup-score-board"
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
)

type Storage interface {
	AddGame(game entity.Game)
	GetGameByTeams(homeTeam, awayTeam string) *entity.Game
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateGame(request football_world_cup_score_board.CreateGame) error {
	gameByTeams := s.storage.GetGameByTeams(request.HomeTeam, request.AwayTeam)

	if gameByTeams != nil {
		return football_world_cup_score_board.ErrGameAlreadyExist
	}

	gameEntity := football_world_cup_score_board.CreateGameToGameEntity(request)

	s.storage.AddGame(gameEntity)

	return nil
}
