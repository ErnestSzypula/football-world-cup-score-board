package service

import (
	"bufio"
	"errors"
	"fmt"
	football_world_cup_score_board "github.com/ErnestSzypula/football-world-cup-score-board"
	"os"
	"strconv"
	"strings"
)

type ActionToken string

const (
	actionTokenStartGame  = "start"
	actionTokenFinishGame = "finish"
	actionTokenUpdateGame = "update"
	actionTokenSummary    = "summary"
	actionTokenHelp       = "help"
	actionTokenExit       = "exit"
)

var helpInfo = "Type \"help\"  for more information."

type Api struct {
	service *Service
}

func NewApi(service *Service) *Api {
	return &Api{
		service: service,
	}
}

func (a *Api) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Football World Cup Score Board")
	fmt.Println(helpInfo)

	for {
		fmt.Printf(">> ")

		scanner.Scan()
		tokens := strings.Split(strings.Trim(strings.ToLower(scanner.Text()), " "), " ")
		if len(tokens) > 0 {
			switch tokens[0] {
			case actionTokenStartGame:
				a.StartGame(tokens[1:])
			case actionTokenUpdateGame:
				a.UpdateGame(tokens[1:])
			case actionTokenFinishGame:
				a.FinishGame(tokens[1:])
			case actionTokenSummary:
				a.Summary()
			case actionTokenHelp:
				a.Help()
			case actionTokenExit:
				return
			default:
				fmt.Println("unsupported command")
				fmt.Println(helpInfo)
			}
		}
	}
}

func (a *Api) StartGame(tokens []string) {
	if len(tokens) != 2 {
		fmt.Printf("ERROR: wrong command args, expect home team and away team, got %v\n", tokens)
		return
	}

	request := football_world_cup_score_board.CreateGameRequest{
		HomeTeamName: tokens[0],
		AwayTeamName: tokens[1],
	}

	if err := a.service.CreateGame(request); err != nil {
		if errors.Is(err, football_world_cup_score_board.ErrGameAlreadyExist) {
			fmt.Printf("ERROR: game home team: %s away team: %s already exist\n", request.HomeTeamName, request.AwayTeamName)
			return
		}
		fmt.Printf("ERROR: a.service.CreateGame %s\n", err.Error())
		return
	}

	fmt.Println("successfully start game")
}

func (a *Api) UpdateGame(tokens []string) {
	if len(tokens) != 4 {
		fmt.Printf("ERROR: wrong command args, expect home team, home team score, away team, away team score"+
			", got %v\n", tokens)
		return
	}

	homeTeamScore, err := strconv.Atoi(tokens[1])
	if err != nil {
		fmt.Printf("ERROR: wrong home team score expect int got %s\n", tokens[1])
		return
	}

	if homeTeamScore < 0 {
		fmt.Printf("ERROR: wrong home team score expect int greater or equal 0 got %d\n", homeTeamScore)
		return
	}

	awayTeamScore, err := strconv.Atoi(tokens[3])
	if err != nil {
		fmt.Printf("ERROR: wrong away team score expect int got %s\n", tokens[2])
		return
	}

	if awayTeamScore < 0 {
		fmt.Printf("ERROR: wrong home team score expect int greater or equal 0 got %d\n", awayTeamScore)
		return
	}

	request := football_world_cup_score_board.UpdateGameRequest{
		HomeTeamName:  tokens[0],
		HomeTeamScore: homeTeamScore,
		AwayTeamName:  tokens[2],
		AwayTeamScore: awayTeamScore,
	}

	if err := a.service.UpdateGame(request); err != nil {
		if errors.Is(err, football_world_cup_score_board.ErrGameNotFound) {
			fmt.Printf("ERROR: game home team: %s away team: %s not found\n", request.HomeTeamName, request.AwayTeamName)
			return
		}
		fmt.Printf("ERROR: a.service.UpdateGame %s\n", err.Error())
		return
	}

	fmt.Println("successfully update game")
}

func (a *Api) FinishGame(tokens []string) {
	if len(tokens) != 2 {
		fmt.Printf("ERROR: wrong command args, expect home team and away team, got %v\n", tokens)
		return
	}

	request := football_world_cup_score_board.FinishGameRequest{
		HomeTeamName: tokens[0],
		AwayTeamName: tokens[1],
	}

	if err := a.service.FinishGame(request); err != nil {
		if errors.Is(err, football_world_cup_score_board.ErrGameNotFound) {
			fmt.Printf("ERROR: game home team: %s away team: %s not found\n", request.HomeTeamName, request.AwayTeamName)
			return
		}
		fmt.Printf("ERROR: a.service.CreateGame %s\n", err.Error())
		return
	}

	fmt.Println("successfully finish game")
}

func (a *Api) Summary() {
	response := a.service.Summary()

	fmt.Println("--- summary ---")
	fmt.Println("")

	for i, item := range response.Items {
		fmt.Printf("%d. %s %d - %d %s\n", i+1, item.HomeTeamName, item.HomeTeamScore, item.AwayTeamScore, item.AwayTeamName)
	}

	fmt.Println("")
}

func (a *Api) Help() {
	fmt.Println("Welcome to Football Score Board help utility!")
	fmt.Println("")
	fmt.Println("available actions and commands:")
	fmt.Println("")
	fmt.Println("START GAME")
	fmt.Println("start {{.HomeTeamName}} {{.AwayTeamName}}")
	fmt.Println("")
	fmt.Println("UPDATE GAME")
	fmt.Println("update {{.HomeTeamName}} {{.HomeTeamScore}} {{.AwayTeamName}} {{.AwayTeamScore}}")
	fmt.Println("")
	fmt.Println("FINISH GAME")
	fmt.Println("finish {{.HomeTeamName}} {{.AwayTeamScore}}")
	fmt.Println("")
	fmt.Println("SUMMARY")
	fmt.Println("summary")
}
