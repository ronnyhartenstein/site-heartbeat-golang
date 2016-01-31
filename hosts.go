package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var file = "hosts.txt"

//ReadHosts Read the `hosts.txt` file to get the Sites
func ReadHosts() []Site {
	// read whole the file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("File %s not found. (%s)", file, err)
	}

	var sites = make([]Site, 0)

	rows := strings.Split(string(b), "\n")
	for _, row := range rows {
		cols := strings.Split(row, ": ")
		if len(cols) == 1 || cols[0][0] == '#' {
			continue
		}
		//log.Printf("cols: len %d cap %d - %s", len(cols), cap(cols), cols)
		url := cols[0]
		title := cols[1]
		sites = append(sites, Site{Url: url, Title: title})
	}

	return sites
}
