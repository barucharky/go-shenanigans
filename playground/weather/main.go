// B''H

/*
go mod init sandbox/weather
go run main.go
*/

// See page 113.

// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"log"
	"os"
	"sandbox/weather/metaweather"
	"text/template"
)

// !+Template
const templ = `Cities found:
{{range .Cities}}----------------------------------------
City Name: {{.Title}}
Woeid:     {{.Woeid}}
{{end}}`

// !-Template

//!+exec
var report = template.Must(template.New("forecast").Parse(templ))

func main() {

	locResult, err := metaweather.SearchLocation(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, locResult); err != nil {
		log.Fatal(err)
	}

}
