package main

import (
	"log"
	"os"

	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file found:", err)
	}

	host := os.Getenv("MAIL_HOST")
	user := os.Getenv("MAIL_USER")
	pass := os.Getenv("MAIL_PASS")

	if host == "" || user == "" || pass == "" {
		log.Fatal("missing MAIL_HOST / MAIL_USER / MAIL_PASS")
	}

	c, err := imapclient.DialTLS(host, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Println("connected")

	if err := c.Login(user, pass).Wait(); err != nil {
		log.Fatal("login:", err)
	}
	log.Println("logged in")

	if err := c.Logout().Wait(); err != nil {
		log.Fatal("logout:", err)
	}
	log.Println("bye")
}