package orchestrator

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/tebeka/selenium"
)

func (orchestrator *Orchestrator) HandleSearchBar(context *actor.Context) error {
	slog.Info("Starting HandleSearchBar with Selenium and geckodriver")

	opts := []selenium.ServiceOption{}
	service, err := selenium.NewGeckoDriverService(main.gecko_path, main.gecko_port, opts...)
	if err != nil {
		slog.Error("Error starting geckodriver service", "error", err)
		return err
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d", main.gecko_port))
	if err != nil {
		slog.Error("Failed to connect to WebDriver", "error", err)
		return err
	}
	defer wd.Quit()

	initialLink := "https://www..com"
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

	searchText := orchestrator.wantedText
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

	slog.Info("New link after form submission", "newLink", newURL)
	return nil
}
