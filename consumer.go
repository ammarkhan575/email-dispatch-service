package main

import (
	"fmt"
	"net/smtp"
	"sync"
)

func emailWorker(id int, ch chan Reciepient, wg *sync.WaitGroup) {
	defer wg.Done()
	for reciepient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// Here you would add logic to send the email using the SMTP server.
		// For demonstration, we will just print the email content.
		// msg := []byte(formattedEmail)
		// formattedEmail := fmt.Sprintf("To: %s <%s>\nSubject: Hello %s!\n\nThis is a test email.\n", reciepient.Name, reciepient.Email, reciepient.Name)

		msg, err := executeEmailTemplate(reciepient)
		if err != nil {
			fmt.Printf("Worker %d failed to execute email template for %s <%s>: %v\n", id, reciepient.Name, reciepient.Email, err)
			continue
		}
		err1 := smtp.SendMail(smtpHost+":"+smtpPort, nil, "ammarkhan575@gmail.com", []string{reciepient.Email}, []byte(msg))
		if err1 != nil {
			fmt.Printf("Worker %d failed to send email to %s <%s>: %v\n", id, reciepient.Name, reciepient.Email, err)
		} else {
			fmt.Printf("Worker %d sent email to %s <%s>\n", id, reciepient.Name, reciepient.Email)
		}
	}
}