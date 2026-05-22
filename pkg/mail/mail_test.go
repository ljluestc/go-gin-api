package mail

import (
	"os"
	"strconv"
	"testing"
)

func TestSend(t *testing.T) {
	mailHost := os.Getenv("MAIL_TEST_HOST")
	mailPort := os.Getenv("MAIL_TEST_PORT")
	mailUser := os.Getenv("MAIL_TEST_USER")
	mailPass := os.Getenv("MAIL_TEST_PASS")
	mailTo := os.Getenv("MAIL_TEST_TO")
	if mailHost == "" || mailPort == "" || mailUser == "" || mailPass == "" || mailTo == "" {
		t.Skip("skip smtp integration test: set MAIL_TEST_HOST/MAIL_TEST_PORT/MAIL_TEST_USER/MAIL_TEST_PASS/MAIL_TEST_TO")
	}
	port, err := strconv.Atoi(mailPort)
	if err != nil {
		t.Fatalf("invalid MAIL_TEST_PORT: %v", err)
	}
	options := &Options{
		MailHost: mailHost,
		MailPort: port,
		MailUser: mailUser,
		MailPass: mailPass,
		MailTo:   mailTo,
		Subject:  "subject",
		Body:     "body",
	}
	err = Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
