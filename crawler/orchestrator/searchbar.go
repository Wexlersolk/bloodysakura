package orchestrator

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"bloodysakura/crawler/visitor"

	"github.com/PuerkitoBio/goquery"
	"github.com/anthdm/hollywood/actor"
)

// HandleSearchBar is the method that initiates the search bar handling process
func (orchestrator *Orchestrator) HandleSearchBar(context *actor.Context) error {
	slog.Info("Starting HandleSearchBar")

	initialLink := "https://audible.com"
	slog.Info("Parsing initial link", "link", initialLink)

	parsedLink, err := url.Parse(initialLink)
	if err != nil {
		slog.Error("Failed to parse initial link", "error", err)
		return err
	}

	slog.Info("Making GET request to the parsed link", "parsedLink", parsedLink.String())
	resp, err := http.Get(parsedLink.String())
	if err != nil {
		slog.Error("Failed to make GET request", "error", err)
		return err
	}
	defer resp.Body.Close()

	slog.Info("GET request successful", "status", resp.Status)

	// Use Goquery to parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		slog.Error("Failed to parse HTML", "error", err)
		return err
	}

	slog.Info("Searching for the search bar in the response body")
	searchFormData, err := orchestrator.findSearchBar(doc, orchestrator.wantedText)
	if err != nil {
		slog.Error("Error while finding search bar", "error", err)
		return err
	}

	if searchFormData != nil {
		slog.Info("Search bar found")
		searchResultsPage, err := orchestrator.submitSearchForm(parsedLink, searchFormData)
		if err != nil {
			slog.Error("Error while submitting search form", "error", err)
			return err
		}

		newLinks, err := orchestrator.extractLinksFromPage(searchResultsPage)
		if err != nil {
			slog.Error("Error while extracting links from search results page", "error", err)
			return err
		}

		slog.Info("Extracted links", "links", newLinks)

		return orchestrator.HandleVisitRequest(context, visitor.NewVisitRequest(newLinks, orchestrator.wantedText))
	}

	slog.Info("No search bar found")
	return nil
}

// findSearchBar searches for the search bar in the HTML document
func (orchestrator *Orchestrator) findSearchBar(doc *goquery.Document, searchText string) (url.Values, error) {
	formData := url.Values{}

	// Find input fields of type 'search' or 'text'
	doc.Find("input[type='search'], input[type='text']").Each(func(i int, s *goquery.Selection) {
		inputName, _ := s.Attr("name")
		if inputName != "" {
			formData.Set(inputName, searchText)
			slog.Debug("Search text set for input", "searchText", searchText)
		}
	})

	// Check if any form data was added
	if len(formData) > 0 {
		slog.Info("Search form data constructed successfully", "formData", formData)
		return formData, nil
	}

	slog.Error("No valid search form found")
	return nil, fmt.Errorf("no search form found")
}

// submitSearchForm submits the search form to the server
func (orchestrator *Orchestrator) submitSearchForm(baseURL *url.URL, formData url.Values) (io.Reader, error) {
	// Adjusting searchURL based on the base URL
	searchURL := baseURL.ResolveReference(&url.URL{Path: formData.Get("action")})
	slog.Info("Submitting search form", "searchURL", searchURL.String(), "formData", formData)

	resp, err := http.PostForm(searchURL.String(), formData)
	if err != nil {
		slog.Error("Failed to submit search form", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	slog.Info("Search form submitted successfully", "status", resp.Status)
	return resp.Body, nil
}

// extractLinksFromPage extracts links from the HTML document
func (orchestrator *Orchestrator) extractLinksFromPage(body io.Reader) ([]string, error) {
	slog.Info("Starting to extract links from page")
	links, err := (&visitor.Visitor{URL: &url.URL{}}).ExtractLinks(body)
	if err != nil {
		slog.Error("Failed to extract links", "error", err)
		return nil, err
	}

	slog.Info("Links extracted successfully", "links", links)
	return links, nil
}
