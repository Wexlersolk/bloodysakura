package orchestrator

import (
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/tebeka/selenium"
)

func (orchestrator *Orchestrator) HandleSearchBar(context *actor.Context) error {
	slog.Info("Starting HandleSearchBar with Selenium and geckodriver")

	opts := []selenium.ServiceOption{}
	service, err := selenium.NewGeckoDriverService(orchestrator.data.GeckoPath, orchestrator.data.GeckoPort, opts...)
	if err != nil {
		slog.Error("Error starting geckodriver service", "error", err)
		return err
	}
	defer service.Stop()

	// Set up capabilities for headless mode
	caps := selenium.Capabilities{
		"browserName": "firefox",
		"moz:firefoxOptions": map[string]interface{}{
			"args": []string{"-headless"}, // Enable headless mode
		},
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d", orchestrator.data.GeckoPort))
	if err != nil {
		slog.Error("Failed to connect to WebDriver", "error", err)
		return err
	}
	defer wd.Quit()

	initialLink := orchestrator.data.VisitUrl.String()
	slog.Info("Opening link", "link", initialLink)
	if err := wd.Get(initialLink); err != nil {
		slog.Error("Failed to load page", "error", err)
		return err
	}

	searchBox, err := wd.FindElement(selenium.ByCSSSelector, "input[type='search']")
	if err != nil {
		slog.Error("Failed to find search bar", "error", err)
		return err
	}

	searchText := orchestrator.data.WantedText
	slog.Info("Typing text in search bar", "searchText", searchText)
	if err := searchBox.SendKeys(searchText); err != nil {
		slog.Error("Failed to type text in search bar", "error", err)
		return err
	}

	slog.Info("Submitting the search using JavaScript")
	jsScript := `var e = new KeyboardEvent('keydown', {bubbles: true, cancelable: true, keyCode: 13}); arguments[0].dispatchEvent(e);`
	if _, err := wd.ExecuteScript(jsScript, []interface{}{searchBox}); err != nil {
		slog.Error("Failed to execute JavaScript for Enter keypress", "error", err)
		return err
	}

	time.Sleep(5 * time.Second)

	newURL, err := wd.CurrentURL()
	if err != nil {
		slog.Error("Failed to get current URL", "error", err)
		return err
	}
	parsedNewURL, err := url.Parse(newURL)
	if err != nil {
		slog.Error("Failed to parse new URL", "error", err)
		return err
	}

	slog.Info("New link after form submission", "newLink", newURL)

	orchestrator.data.VisitUrl = parsedNewURL

	return nil
}
