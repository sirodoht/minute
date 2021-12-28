package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type Config struct {
	smtpServer string
	username   string
	password   string
	from       string
	to         string
}

type Checks []string

func notify(config Config, website string) {
	server := config.smtpServer
	serverWithPort := config.smtpServer + ":587"
	from := config.from
	username := config.username
	password := config.password
	to := config.to

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Uptime: " + website + "\n\n" +
		website + " is down."

	err := smtp.SendMail(serverWithPort,
		smtp.PlainAuth("", username, password, server),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Printf("Notification sent for %s\n", website)
}

func ping(config Config, websites []string) {
	for i := 0; i < len(websites); i++ {
		websitesSingle := websites[i]
		if !strings.Contains(websitesSingle, "http") {
			continue
		}
		fmt.Print("Pinging " + websitesSingle)

		resp, err := http.Get(websitesSingle)
		if err != nil {
			fmt.Printf("NOT OK\n%s\n", err)
			notify(config, websitesSingle)
			continue
		}

		if resp.StatusCode != 200 {
			fmt.Print(" is NOT OK")
			notify(config, websitesSingle)
		} else {
			fmt.Print(" is OK")
		}

		fmt.Println()
	}
}

func main() {
	fileName := os.Args[1]
	fmt.Printf("Reading from %s\n", fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := strings.Split(string(data), "\n")
	config := Config{
		smtpServer: contents[0],
		username:   contents[1],
		password:   contents[2],
		from:       contents[3],
		to:         contents[4],
	}
	fmt.Printf("Config:\n%+v\n", config)

	var websites Checks
	websites = contents[5:]
	fmt.Printf("Website list:\n%v\n", websites)

	for {
		nextTime := time.Now().Truncate(time.Minute)
		nextTime = nextTime.Add(time.Minute)
		time.Sleep(time.Until(nextTime))

		fmt.Printf("\n* TICK: %s\n", nextTime.String())
		ping(config, websites)
	}
}
