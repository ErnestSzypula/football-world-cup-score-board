package main

import (
	"github.com/ErnestSzypula/football-world-cup-score-board/service"
	"github.com/ErnestSzypula/football-world-cup-score-board/storage"
)

func main() {
	db := storage.NewInMemoryStorage()

	s := service.NewService(db)

	api := service.NewApi(s)

	api.Start()
}
