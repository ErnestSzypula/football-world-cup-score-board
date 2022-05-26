# FootballWorldCupScoreBoard

Football World CupScore Board is simple console application storing current football game results in memory.

Run application:

```make run```

Run tests:

```make test```

## supported command

### Start game:

start game with starting score equal 0

```start {{.HomeTeamName}} {{.AwayTeamName}}```

### Update game:

update game score between teams if game exists

```update {{.HomeTeamName}} {{.HomeTeamScore}} {{.AwayTeamName}} {{.AwayTeamScore}}```

args:

```{{.HomeTeamScore}} | {{.AwayTeamScore}}``` -
integers greater than or equal zero

### Finish game:

finish game between teams if game exists

```finish {{.HomeTeamName}} {{.AwayTeamScore}}```

### Summary:

print score board with all active games

```summary```

### Help:

print help

```help```

### Exit

exit application

```exit```

