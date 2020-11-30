package main

import (
	"fmt"
	"github.com/svenwiltink/adventofdiscord"
	"log"
	"os"
)

func main() {
	session, exists := os.LookupEnv("AOC_SESSION")
	if !exists {
		log.Fatalln("AOC_SESSION env variable required")
	}

	leaderboard, exists := os.LookupEnv("AOC_LEADERBOARD")

	fetcher := adventofdiscord.NewStatsCollector(leaderboard, session)
	s, err := fetcher.FetchStats()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", s)
}
