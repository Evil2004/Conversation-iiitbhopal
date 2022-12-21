package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"log"
	"net/http"

	"github.com/google/uuid"
	gomail "gopkg.in/gomail.v2"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}

func Requestcookie(r *http.Request) bool {
	_, err := r.Cookie("logged")

	return err == nil
}

// set cookie.

func set_get(w http.ResponseWriter, r *http.Request) {

	cookie, err = r.Cookie("logged")
	if err != nil {
		// creating a new session id.
		sID, err := uuid.NewRandom()

		s := sID.String() + "|" + user1.UserName
		checkerr(err)
		cookie = &http.Cookie{
			Name:  "logged",
			Value: s,
		}
		// user session struct.
		usersession[cookie.Value] = user1.UserName

		http.SetCookie(w, cookie)
	} else {
		split := strings.Split(cookie.Value, "|")

		usersession[split[0]] = split[1]

	}

}

func otpgenerator() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func sendEmail(email string) {

	otp, err = otpgenerator()
	checkerr(err)
	body := "Hello User, welcome to Conversation. Your Email verification code is " + otp

	m := gomail.NewMessage()
	m.SetHeader("From", "letsconversate@zohomail.in")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", body)
	d := gomail.NewDialer("smtp.zoho.in", 465, "letsconversate@zohomail.in", "conversation12")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		// checkerr(err)
		panic(err)
	}

}

func (thread *thread) Created_time() string {
	return time.Now().Format("02-01-2006")
}

func (post *post) Created_time() string {
	return time.Now().Format("02-01-2006")
}
