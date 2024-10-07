package data

import (
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	VisitUrl   *url.URL
	WantedText string
	GeckoPort  int
	GeckoPath  string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	visitURL := os.Getenv("VISIT_URL")
	wantedText := os.Getenv("WANTED_TEXT")
	geckoPortSTR := os.Getenv("GECKO_PORT")
	geckoPath := os.Getenv("GECKO_PATH")

	parsedURL, err := url.Parse(visitURL)
	if err != nil {
		return nil, err
	}

	geckoPort, err := strconv.Atoi(geckoPortSTR)
	if err != nil {
		log.Fatal("Invalid geckodriver port, must be an integer", "error", err)
	}

	return &Config{
		VisitUrl:   parsedURL,
		WantedText: wantedText,
		GeckoPort:  geckoPort,
		GeckoPath:  geckoPath,
	}, nil
}
