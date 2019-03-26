package helpers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type MailHelper struct {
	SenderID string
	ToIDs    []string `json:"to"`
	Subject  string	`json:"subject"`
	Body     string `json:"body"`
}

type SmtpServer struct {
	host string
	port string
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *MailHelper) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.SenderID)
	if len(mail.ToIDs) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.ToIDs, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return message
}

func (mail MailHelper) SendMail () error {
	mail.SenderID = "yournextwish@gmail.com"
	fmt.Println("mail", mail)

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	log.Println(smtpServer.host)
	//build an auth
	auth := unencryptedAuth{
		smtp.PlainAuth("", mail.SenderID, "thephantom28", smtpServer.host),
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}

	// step 2: add all from and to
	if err = client.Mail(mail.SenderID); err != nil {
		return err
	}
	for _, k := range mail.ToIDs {
		if err = client.Rcpt(k); err != nil {
			return err
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()

	log.Println("Mail sent successfully")
	return err
}