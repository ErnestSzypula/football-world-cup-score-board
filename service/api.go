package service

import (
	"bufio"
	"errors"
	"fmt"
	football_world_cup_score_board "github.com/ErnestSzypula/football-world-cup-score-board"
	"os"
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
		tokens := strings.Split(strings.ToLower(scanner.Text()), " ")
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
		fmt.Printf("wrong command args, expect home team and away team, got %v\n", tokens)
		return
	}

	request := football_world_cup_score_board.CreateGame{
		HomeTeam: tokens[0],
		AwayTeam: tokens[1],
	}

	if err := a.service.CreateGame(request); err != nil {
		if errors.Is(err, football_world_cup_score_board.ErrGameAlreadyExist) {
			fmt.Printf("ERROR: game home team: %s away team: %s already exist\n", request.HomeTeam, request.AwayTeam)
			return
		}

		return
	}

	fmt.Println("successfully start game")
}

func (a *Api) UpdateGame(tokens []string) {
	if len(tokens) != 4 {
		fmt.Printf("wrong command args, expect home team, away team, home team score, away team score"+
			", got %v\n", tokens)
		return
	}

	fmt.Println("successfully update game")
}

func (a *Api) FinishGame(tokens []string) {
	if len(tokens) != 2 {
		fmt.Printf("wrong command args, expect home team and away team, got %v\n", tokens)
		return
	}

	fmt.Println("successfully finish game")
}

func (a *Api) Summary() {
	fmt.Println("games summary")
}

func (a *Api) Help() {
	fmt.Println("help")
}
