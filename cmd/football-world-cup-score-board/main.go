package main

import "github.com/ErnestSzypula/football-world-cup-score-board/service"

func main() {
	api := service.NewApi()

	api.Start()
}
