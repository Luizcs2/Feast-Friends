
package utils

import (
	"regexp"
	"strings"
	"net/mail"
	"html"
)


func ImageFileisValid(filename string , filesize int64, maxfilesize int64) bool{
	filename = strings.TrimSpace(filename) // sanitize input
	if filesize > maxfilesize*1024*1024 || filesize <= 0 || filename == "" {
		return false
	}
	return true
}

func UsernameIsValid(username string) bool {
	username = html.EscapeString(strings.TrimSpace(username)) // sanitize input making sure no html tags are present
	if len(username) < 3 || len(username) > 20{
		return false 
	}
	//check if username is alphanumeric
	re:= regexp.MustCompile(`^[a-zA-Z0-9_.]+$`) 

	return re.MatchString(username)
}

//this func checks for the email validity using net/mail package
func EmailIsValid(email string) bool {
	email = strings.ToLower(strings.TrimSpace(email)) // sanitize 
	if email == "" {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

// this func checks for the pasword strength
func Passwordstrength(password string) bool {
	password = strings.TrimSpace(password) // sanitize input making sure no html tags are present
	if len(password)<8 {
		return false
	}

	uppercase := regexp.MustCompile(`[A-Z]`)
	specialChar := regexp.MustCompile(`[!@#$&*]`)
	digit := regexp.MustCompile(`[0-9]`)

	if  !uppercase.MatchString(password) || !specialChar.MatchString(password) || !digit.MatchString(password) {
		return false /// we set the mimimum lenght of the password to 8
	}

	return true
}