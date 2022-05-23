build:
	go build -o bin/football-world-cup-score-board cmd/football-world-cup-score-board/main.go

run:
	go run cmd/football-world-cup-score-board/main.go

fmt:
	go fmt ./...