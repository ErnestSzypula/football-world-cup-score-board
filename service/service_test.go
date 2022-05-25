package service

import (
	football_world_cup_score_board "github.com/ErnestSzypula/football-world-cup-score-board"
	"github.com/ErnestSzypula/football-world-cup-score-board/entity"
	"github.com/ErnestSzypula/football-world-cup-score-board/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService(t *testing.T) {
	db := storage.NewInMemoryStorage()
	service := NewService(db)

	// test create game
	createGameRequest1 := football_world_cup_score_board.CreateGameRequest{
		HomeTeamName: "Poland",
		AwayTeamName: "Germany",
	}

	err := service.CreateGame(createGameRequest1)
	require.NoError(t, err)

	err = service.CreateGame(createGameRequest1)
	require.ErrorIs(t, err, football_world_cup_score_board.ErrGameAlreadyExist)

	game1, index := db.GetActiveGameByTeams(createGameRequest1.HomeTeamName, createGameRequest1.AwayTeamName)
	require.NotNil(t, game1)
	require.Equal(t, 0, index)

	assert.Equal(t, entity.GameStatusActive, game1.Status)
	assert.Equal(t, createGameRequest1.HomeTeamName, game1.HomeTeamName)
	assert.Equal(t, 0, game1.HomeTeamScore)
	assert.Equal(t, createGameRequest1.AwayTeamName, game1.AwayTeamName)
	assert.Equal(t, 0, game1.AwayTeamScore)

	createGameRequest2 := football_world_cup_score_board.CreateGameRequest{
		HomeTeamName: "Slovakia",
		AwayTeamName: "Austria",
	}

	err = service.CreateGame(createGameRequest2)
	require.NoError(t, err)

	game2, _ := db.GetActiveGameByTeams(createGameRequest2.HomeTeamName, createGameRequest2.AwayTeamName)
	require.NotNil(t, game2)

	assert.Equal(t, entity.GameStatusActive, game2.Status)
	assert.Equal(t, createGameRequest2.HomeTeamName, game2.HomeTeamName)
	assert.Equal(t, 0, game2.HomeTeamScore)
	assert.Equal(t, createGameRequest2.AwayTeamName, game2.AwayTeamName)
	assert.Equal(t, 0, game2.AwayTeamScore)

	// test summary game
	summary1 := service.Summary()

	require.Len(t, summary1.Items, 2)
	assert.Equal(t, createGameRequest2.HomeTeamName, summary1.Items[0].HomeTeamName)
	assert.Equal(t, 0, summary1.Items[0].HomeTeamScore)
	assert.Equal(t, createGameRequest2.AwayTeamName, summary1.Items[0].AwayTeamName)
	assert.Equal(t, 0, summary1.Items[0].AwayTeamScore)

	assert.Equal(t, createGameRequest1.HomeTeamName, summary1.Items[1].HomeTeamName)
	assert.Equal(t, 0, summary1.Items[1].HomeTeamScore)
	assert.Equal(t, createGameRequest1.AwayTeamName, summary1.Items[1].AwayTeamName)
	assert.Equal(t, 0, summary1.Items[1].AwayTeamScore)

	// update game
	updateGameRequest1 := football_world_cup_score_board.UpdateGameRequest{
		HomeTeamName:  createGameRequest1.HomeTeamName,
		HomeTeamScore: 1,
		AwayTeamName:  createGameRequest1.AwayTeamName,
		AwayTeamScore: 2,
	}

	err = service.UpdateGame(updateGameRequest1)
	require.NoError(t, err)

	summary2 := service.Summary()

	require.Len(t, summary2.Items, 2)
	assert.Equal(t, updateGameRequest1.HomeTeamName, summary2.Items[0].HomeTeamName)
	assert.Equal(t, updateGameRequest1.HomeTeamScore, summary2.Items[0].HomeTeamScore)
	assert.Equal(t, updateGameRequest1.AwayTeamName, summary2.Items[0].AwayTeamName)
	assert.Equal(t, updateGameRequest1.AwayTeamScore, summary2.Items[0].AwayTeamScore)

	assert.Equal(t, createGameRequest2.HomeTeamName, summary2.Items[1].HomeTeamName)
	assert.Equal(t, 0, summary2.Items[1].HomeTeamScore)
	assert.Equal(t, createGameRequest2.AwayTeamName, summary2.Items[1].AwayTeamName)
	assert.Equal(t, 0, summary2.Items[1].AwayTeamScore)

	// update game fail path
	updateGameRequest2 := football_world_cup_score_board.UpdateGameRequest{
		HomeTeamName:  "bad1",
		HomeTeamScore: 1,
		AwayTeamName:  "bad2",
		AwayTeamScore: 2,
	}

	err = service.UpdateGame(updateGameRequest2)
	require.ErrorIs(t, err, football_world_cup_score_board.ErrGameNotFound)

	// finish game test
	finishGameRequest1 := football_world_cup_score_board.FinishGameRequest{
		HomeTeamName: createGameRequest1.HomeTeamName,
		AwayTeamName: createGameRequest1.AwayTeamName,
	}

	err = service.FinishGame(finishGameRequest1)
	require.NoError(t, err)

	summary3 := service.Summary()
	require.Len(t, summary3.Items, 1)

	assert.Equal(t, createGameRequest2.HomeTeamName, summary3.Items[0].HomeTeamName)
	assert.Equal(t, 0, summary3.Items[0].HomeTeamScore)
	assert.Equal(t, createGameRequest2.AwayTeamName, summary3.Items[0].AwayTeamName)
	assert.Equal(t, 0, summary3.Items[0].AwayTeamScore)

	// finish game fail path
	finishGameRequest2 := football_world_cup_score_board.FinishGameRequest{
		HomeTeamName: "bad1",
		AwayTeamName: "bad2",
	}

	err = service.FinishGame(finishGameRequest2)
	require.ErrorIs(t, err, football_world_cup_score_board.ErrGameNotFound)
}
