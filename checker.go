package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
)

// Site meta data
type Site struct {
	Url   string
	Title string
}

// Check the site title
func Check(site Site) bool {
	//log.Printf("Check %s ..\n", site.Url)
	doc, err := goquery.NewDocument(site.Url)
	if err != nil {
		log.Print(err)
		return false
	}

	title := doc.Find("title").Text()
	matched, err := regexp.MatchString(site.Title, title)
	if !matched {
		log.Printf("Title for %s is '%s', but should '%s'\n", site.Url, title, site.Title)
		return false
	}
	return true
}
