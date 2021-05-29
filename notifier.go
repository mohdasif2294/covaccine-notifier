package main

import (
	"fmt"
	"net/smtp"

	"github.com/PrasadG193/covaccine-notifier/awsclient"
)

func sendMail(id, pass, body string) error {
	msg := "From: " + id + "\n" +
		"To: " + id + "\n" +
		"Subject: Vaccination slots are available\n\n" +
		"Vaccination slots are available at the following centers:\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", id, pass, "smtp.gmail.com"),
		id, []string{id}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}

func sendSMS(vaccType, num string) error {
	msg := fmt.Sprintf("ALERT: %s available on CoWIN App at %s", vaccType,district)
	awsclient.SendSMS(num, msg)
	return nil
}
