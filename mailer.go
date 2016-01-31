package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strconv"
)

type emailConfig struct {
	Username  string
	Password  string
	Host      string
	Port      int
	Receivers []string
	Sender    string
}

const config = "mailer.conf"

//MailSiteDown mail on site down
func MailSiteDown(site Site) {
	// https://www.socketloop.com/tutorials/golang-send-email-and-smtp-configuration-example

	emailConf := readConfig()

	emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

	message := []byte("Site is down!") // your message

	log.Printf("Send mail for %s ..", site)
	err := smtp.SendMail(emailConf.Host+":"+strconv.Itoa(emailConf.Port),
		emailauth,
		emailConf.Sender,
		emailConf.Receivers,
		message,
	)

	if err != nil {
		fmt.Println(err)
	}
}

func readConfig() *emailConfig {

	b, err := ioutil.ReadFile(config)
	if err != nil {
		log.Panicf("Conf %s not found. (%s)", config, err)
	}

	// http://blog.golang.org/json-and-go
	var conf emailConfig
	err = json.Unmarshal(b, &conf)
	if err != nil {
		log.Panicf("Conf %s corrupt. (%s)", config, err)
	}
	log.Printf("Mailer conf: %s", conf)

	return &conf
}
