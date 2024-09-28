package main

import (
	"bloodysakura/crawler/orchestrator"
	"bloodysakura/crawler/visitor"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	visitURL := os.Getenv("VISIT_URL")
	wantedText := os.Getenv("WANTED_TEXT")

	parsedURL, err := url.Parse(visitURL)
	if err != nil {
		log.Fatal("Invalid VISIT_URL")
	}

	engine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	pid := engine.Spawn(orchestrator.NewOrchestrator(wantedText, parsedURL.Host), "manager")

	time.Sleep(time.Millisecond * 200)

	engine.Send(pid, visitor.NewVisitRequest([]string{visitURL}, wantedText))

	time.Sleep(time.Second * 1000)
}
