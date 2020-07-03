// B''H

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/barucharky/go-shenanigans/projects/weather/metaweather"
)

// !+Template
const templ = `5-Day Forecast for {{.Title}}:
{{range .Items}}----------------------------------------
Date:     {{.Date}}
Condition {{.Condition}}
Low:      {{.MinTemp}}
High:     {{.MaxTemp}}
Humidity: {{.Humidity}}
{{end}}`

// !-Template

//!+exec
var report = template.Must(template.New("forecast").Parse(templ))

func main() {

	locResult, err := metaweather.SearchLocation(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	// -- -------------------------------------------
	// Get selection
	selNum, err := getSel(locResult)

	if err != nil {
		log.Fatal(err)
	}

	// -- -------------------------------------------
	// Get the city name from the result struct
	var cityWoeid int

	for cityNum, city := range *locResult {
		if cityNum == selNum {
			cityWoeid = city.Woeid
			break
		}
	}

	// -- -------------------------------------------
	// Get forecast using cityWoeid
	forcastResult, err := metaweather.GetForecast(cityWoeid)

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, forcastResult); err != nil {
		log.Fatal(err)
	}

}

func getSel(locResult *[]metaweather.LocationSearchResult) (int, error) {

	// Get the number of options
	var resultLen int = len(*locResult)

	// If there are no options, return an error. If one, return that option
	if resultLen == 0 {
		return 0, fmt.Errorf("No city options available")
	} else if resultLen == 1 {
		return 0, nil
	}

	// Show the user the list of options
	for cityNum, city := range *locResult {
		fmt.Printf("%3d   %s\n", cityNum, city.Title)
	}

	// Start the selection process
	fmt.Println("Select a city:")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var selNum int

	// User must enter a viable option at this point
	for true {

		input, err := reader.ReadString('\n')

		if err != nil {
			return 0, err
		}

		input = strings.TrimSpace(input)

		selNum, err = strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please type a number")
			continue
		}

		if selNum > resultLen {
			fmt.Println("Please select from the list")
			continue
		} else {
			break
		}
	}

	// Return the selection
	return selNum, nil
}
