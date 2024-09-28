package orchestrator

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/anthdm/hollywood/actor"
	"golang.org/x/net/html"
)

func (orchestrator *Orchestrator) handleSearchBar(context *actor.Context) error {
	// Get the current URL (assuming you have a starting point or current link)
	initialLink := "https://example.com" // Replace this with the starting link
	parsedLink, err := url.Parse(initialLink)
	if err != nil {
		return err
	}

	// Perform the visit to the initial page
	resp, err := http.Get(parsedLink.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Buffer for reading the response body
	w := &bytes.Buffer{}
	r := io.TeeReader(resp.Body, w)

	// Look for the search bar on the page
	searchFormData, err := orchestrator.findSearchBar(r, orchestrator.wantedText)
	if err != nil {
		return err
	}

	// If we found a search form, submit it with the wanted text
	if searchFormData != nil {
		searchResultsPage, err := orchestrator.submitSearchForm(parsedLink, searchFormData)
		if err != nil {
			return err
		}

		// Once we have the new page, extract links and call handleVisitRequest
		newLinks, err := orchestrator.extractLinksFromPage(searchResultsPage)
		if err != nil {
			return err
		}

		// Call handleVisitRequest with the new links
		return orchestrator.handleVisitRequest(context, NewVisitRequest(newLinks, orchestrator.wantedText))
	}

	return nil
}

// Function to find a search bar in the page's HTML
func (orchestrator *Orchestrator) findSearchBar(body io.Reader, searchText string) (url.Values, error) {
	tokenizer := html.NewTokenizer(body)
	formData := url.Values{}

	var isInForm bool
	var formAction string
	var inputName string

	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			return nil, fmt.Errorf("search bar not found")
		}

		token := tokenizer.Token()

		if token.Type == html.StartTagToken {
			// Look for the form tag
			if token.Data == "form" {
				isInForm = true
				// Find the form's action URL
				for _, attr := range token.Attr {
					if attr.Key == "action" {
						formAction = attr.Val
					}
				}
			}

			// Look for the input field inside the form
			if isInForm && token.Data == "input" {
				for _, attr := range token.Attr {
					if attr.Key == "name" {
						inputName = attr.Val
					}
				}

				// Add the wanted text as the value for the input field
				formData.Set(inputName, searchText)
			}
		}

		if token.Type == html.EndTagToken && token.Data == "form" {
			// We're done with the form
			break
		}
	}

	// Return the form data if we found a form and input field
	if isInForm && formAction != "" && inputName != "" {
		return formData, nil
	}

	return nil, fmt.Errorf("no search form found")
}

// Function to submit the search form and return the new page
func (orchestrator *Orchestrator) submitSearchForm(baseURL *url.URL, formData url.Values) (io.Reader, error) {
	searchURL := baseURL.ResolveReference(&url.URL{Path: formData.Get("action")})
	resp, err := http.PostForm(searchURL.String(), formData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Return the body of the search results page
	return resp.Body, nil
}

// Function to extract links from a page
func (orchestrator *Orchestrator) extractLinksFromPage(body io.Reader) ([]string, error) {
	// Extract links using the same method as in Visitor
	return (&Visitor{URL: &url.URL{}}).ExtractLinks(body)
}
