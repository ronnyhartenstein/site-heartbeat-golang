package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	//"time"
)

// Site meta data
type Site struct {
	Url   string
	Title string
}

// Check the site title
func Check(site Site) bool {
	// recover on e.g. "x509: certificate is valid errors"
	defer func() {
		if p := recover(); p != nil {
			log.Printf("internal error: %v", p)
		}
	}()

	//log.Printf("Check %s ..\n", site.Url)
	//start := time.Now()
	doc, err := goquery.NewDocument(site.Url)
	//runtime := time.Since(start).Seconds()
	if err != nil {
		log.Print(err)
		return false
	}

	title := doc.Find("title").Text()
	matched, err := regexp.MatchString(site.Title, title)
	if !matched {
		//log.Printf("Title for %s is '%s', but should '%s'\n", site.Url, title, site.Title)
		MailSiteDown(site)
		return false
	} else {
		//log.Printf("%s finished in %2fs, has Title '%s', must contain '%s' .. GOOD!\n", site.Url, runtime, title, site.Title)
		return true
	}
}
