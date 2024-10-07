package main

import (
	"bloodysakura/crawler/orchestrator"
	"bloodysakura/data"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

func main() {

	config, err := data.LoadData()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	engine, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	engine.Spawn(orchestrator.NewOrchestrator(config), "manager")

	time.Sleep(time.Millisecond * 200)

	time.Sleep(time.Second * 1000)
}
