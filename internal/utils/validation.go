// validation.go this file contains all the validation functions for the user input
// user input it includes checks for email, username, password strength and image file validation
// inputs are sanitized by trimming whitespaces, html and tags to prevent XSS attacks

package utils

import (
	"regexp"
	"strings"
	"net/mail"
	"html"
)

//this func checks if the image file is valid
//file is only valid if its size is less than maxfilesize and greater than 0 and filename is not empty
func ImageFileisValid(filename string , filesize int64, maxfilesize int64) bool{
	filename = strings.TrimSpace(filename) // sanitize input
	if filesize > maxfilesize*1024*1024 || filesize <= 0 || filename == "" {
		return false
	}
	return true
}

//this func checks if the username is valid
//user name is only alphanumeric and between 3 to 20 characters and contains only _ and . as special characters
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
//email is only valid if it contains @ and . and no spaces
func EmailIsValid(email string) bool {
	email = strings.ToLower(strings.TrimSpace(email)) // sanitize 
	if email == "" {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

// this func checks for the pasword strength
//password is only valid if it contains at least 8 characters, one uppercase letter, one special character and one digit
func PasswordStrength(password string) bool {
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