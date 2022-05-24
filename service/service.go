package service

import (
	football_world_cup_score_board "github.com/ErnestSzypula/football-world-cup-score-board"
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
	"github.com/ErnestSzypula/football-world-cup-score-board/pkg/utils"
	"time"
)

type Storage interface {
	AddGame(game entity.Game)
	GetActiveGameByTeams(homeTeam, awayTeam string) (*entity.Game, int)
	UpdateGame(index int, game entity.Game)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateGame(request football_world_cup_score_board.CreateGameRequest) error {
	gameByTeams, _ := s.storage.GetActiveGameByTeams(request.HomeTeamName, request.AwayTeamName)

	if gameByTeams != nil {
		return football_world_cup_score_board.ErrGameAlreadyExist
	}

	gameEntity := football_world_cup_score_board.CreateGameRequestToGameEntity(request)

	s.storage.AddGame(gameEntity)

	return nil
}

func (s *Service) UpdateGame(request football_world_cup_score_board.UpdateGameRequest) error {
	gameEntity, index := s.storage.GetActiveGameByTeams(request.HomeTeamName, request.AwayTeamName)

	if gameEntity == nil {
		return football_world_cup_score_board.ErrGameNotFound
	}

	gameEntity.HomeTeamScore = request.HomeTeamScore
	gameEntity.AwayTeamScore = request.AwayTeamScore
	gameEntity.UpdatedAt = time.Now()

	s.storage.UpdateGame(index, *gameEntity)

	return nil
}

func (s *Service) FinishGame(request football_world_cup_score_board.FinishGameRequest) error {
	gameEntity, index := s.storage.GetActiveGameByTeams(request.HomeTeamName, request.AwayTeamName)

	if gameEntity == nil {
		return football_world_cup_score_board.ErrGameNotFound
	}

	now := time.Now()
	gameEntity.Status = entity.GameStatusFinished
	gameEntity.UpdatedAt = now
	gameEntity.FinishedAt = utils.Pointer(now)

	s.storage.UpdateGame(index, *gameEntity)

	return nil
}
