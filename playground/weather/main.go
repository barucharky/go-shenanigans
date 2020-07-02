// B''H

/*
go mod init sandbox/weather
go run main.go
*/

// See page 113.

// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sandbox/weather/metaweather"
	"strconv"
	"strings"
	"text/template"
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

	// Print city options
	var resultLen int = len(*locResult)

	for cityNum, city := range *locResult {
		fmt.Printf("%3d   %s\n", cityNum, city.Title)
	}

	// Get selection
	fmt.Println("Select a city:")
	selNum, err := getSel(resultLen)

	// Get the city name from the result struct
	var cityWoeid int

	for cityNum, city := range *locResult {
		if cityNum == selNum {
			cityWoeid = city.Woeid
			break
		}
	}

	// Get forecast using cityWoeid

	forcastResult, err := metaweather.GetForecast(cityWoeid)

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, forcastResult); err != nil {
		log.Fatal(err)
	}

}

func getSel(cityNum int) (int, error) {

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var selNum int

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

		if selNum > cityNum {
			fmt.Println("Please select from the list")
			continue
		} else {
			break
		}
	}

	return selNum, nil

}
