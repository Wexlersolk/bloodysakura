package main

import (
	"bloodysakura/crawler"
	"log"
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

	engine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	pid := engine.Spawn(crawler.NewOrchestrator(wantedText), "manager")

	time.Sleep(time.Millisecond * 200)

	engine.Send(pid, crawler.NewVisitRequest([]string{visitURL}, wantedText))

	time.Sleep(time.Second * 1000)
}
