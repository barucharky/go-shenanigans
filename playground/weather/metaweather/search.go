// B''H

package metaweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchLocation(terms []string) (*[]LocationSearchResult, error) {

	// -- -------------------------------------------
	q := url.QueryEscape(strings.Join(terms, " "))

	resp, err := http.Get(LocationURL + "search/?query=" + q)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// -- -------------------------------------------
	var result []LocationSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	return &result, nil
}

func GetForecast(woeid int) (*[]WeatherResult, error) {

	// -- -------------------------------------------
	q := url.QueryEscape(strconv.Itoa(woeid))

	resp, err := http.Get(LocationURL + q)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// -- -------------------------------------------
	var result []WeatherResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("here's the problem")
		return nil, err
	}

	// -- -------------------------------------------
	return &result, nil
}

//!-
