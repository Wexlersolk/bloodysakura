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

	engine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := engine.Spawn(crawler.NewManager(), "manager")

	time.Sleep(time.Millisecond * 200)
	engine.Send(pid, crawler.NewVisitRequest([]string{visitURL}))

	time.Sleep(time.Second * 1000)
}
