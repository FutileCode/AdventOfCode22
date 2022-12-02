package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("strategy.txt")
	if err != nil {
		log.Fatal(err)
	}

	split := strings.Split(string(file), "\n")

	const (
		rock     int = 1
		paper    int = 2
		scissors int = 3
	)

	type strat struct {
		opponent string
		player   string
	}

	moves := make(map[strat]int)

	for _, v := range split {
		move := strings.Fields(v)
		moves[strat{move[0], move[1]}] = 0
	}

	for strat, score := range moves {
		switch {
		case strat.player == "X":
			score += rock
		case strat.player == "Y":
			score += paper
		case strat.player == "Z":
			score += scissors
		}
		moves[strat] = score
		switch {
		case strat.player == "X" && strat.opponent == "C", strat.player == "Y" && strat.opponent == "A", strat.player == "Z" && strat.opponent == "B":
			score += 6
		case strat.player == "X" && strat.opponent == "A", strat.player == "Y" && strat.opponent == "B", strat.player == "Z" && strat.opponent == "C":
			score += 3
		}
		moves[strat] = score
	}

	var total int
	for _, v := range moves {
		total += v
	}

	fmt.Printf("Perfect strategy total score = %v", total)
}
