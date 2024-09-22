package main

import (
	"bloodysakura/crawler"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	e.Spawn(crawler.NewOrchestrator(), "Orchestrator")

	time.Sleep(time.Second * 10)
}
