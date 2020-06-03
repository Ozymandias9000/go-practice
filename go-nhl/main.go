package main

import (
	"go-nhl/nhlApi"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	log.Printf("Log Opened: %v", now.String())

	rosterFile, err := os.OpenFile("roster.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("error opening rosterFile, %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("error fetching all teams, %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(len(teams))

	var responses = make(chan []nhlApi.Player)

	for _, team := range teams {
		go func(team nhlApi.Team) {
			res, err := nhlApi.GetRoster(team)
			if err != nil {
				log.Fatalf("Error getting roster: %v", err)
			}
			responses <- res
			wg.Done()
		}(team)
	}

	go func() {
		wg.Wait()
		close(responses)
	}()

	display(responses)

	log.Printf("took %v", time.Now().Sub(now).String())
}

func display(responses chan []nhlApi.Player) {
	for res := range responses {
		for _, p := range res {
			log.Printf("Player Name: %s", p.Person.FullName)
			log.Printf("Number: %s", p.JerseyNumber)
			log.Printf("Position: %s", p.Position.Name)
			log.Println("----------------------")
		}
	}
}
