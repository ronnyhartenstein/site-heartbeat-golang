package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"strconv"
)

type emailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Receiver string
	Sender   string
}

const config = "mailer.conf"

//MailSiteDown mail on site down
func MailSiteDown(site Site) {
	// https://www.socketloop.com/tutorials/golang-send-email-and-smtp-configuration-example
	// https://gist.github.com/chrisgillis/10888032

	conf := readConfig()

	from := mail.Address{"", conf.Sender}
	to := mail.Address{"", conf.Receiver}
	subj := fmt.Sprintf("Site %s is down", site.Url)
	body := fmt.Sprintf("Site %s should contain '%s' in title but doesn't!", site.Url, site.Title)

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := conf.Host + ":" + strconv.Itoa(conf.Port)

	auth := smtp.PlainAuth("", conf.Username, conf.Password, conf.Host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         conf.Host,
	}

	//log.Printf("Send mail for %s ..", site)

	if conf.Port != 465 {
		log.Print("Mailing only implemented and tested for TLS over Port 465 (e.g. Gmail)")
		return
	}

	doRawTLSConn(servername, tlsconfig, auth, from, to, message)
}

// Here is the key, you need to call tls.Dial instead of smtp.Dial
// for smtp servers running on 465 that require an ssl connection
// from the very beginning (no starttls)
func doRawTLSConn(servername string, tlsconfig *tls.Config, auth smtp.Auth, from mail.Address, to mail.Address, message string) {

	host, _, _ := net.SplitHostPort(servername)

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}

// read mailer config
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
	//log.Printf("Mailer conf: %s", conf)

	return &conf
}
