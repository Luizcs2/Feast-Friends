// jwt.go contains all the functions to handle JWT tokens
// it includes token generation, validation and extraction of claims
// it uses golang-jwt/jwt package to handle JWT tokens
// it also uses gotrue-go package to validate tokens against Supabase Auth
// JWT secret and expiration are loaded from config package

package utils

import (
	"errors"
	"feast-friends-api/internal/config"
	"feast-friends-api/pkg/logger"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase-community/gotrue-go"
)

// set variables for JWT and Supabase
var JWSecret, JWExpiration, url, key = config.Get().JWT.Secret, config.Get().JWT.Expiration, config.Get().Supabase.URL, config.Get().Supabase.AKey

//this function will verify if the JWT secret is set in the config
func VerifyJWTConfig() {
	if JWSecret == "" {
		logger.Error("JWT secret is not set in the config. Cannot start application.")
	} else {
		logger.Info("JWT secret is set in the config")
	}
		
}

// this function will validate a JWT token and return the user ID if valid
// it uses gotrue-go package to validate the token against Supabase Auth
func ValidateToken(tokenString string) (userID string, err error) {

	tokenString = strings.TrimPrefix(tokenString, "Bearer ") // trimmng the Bearer prefix if present

	client := gotrue.New(url,key) // create a new gotrue client using supabase url and anon key
	if client == nil {
		logger.Error("failed to create gotrue client as it is nil")
		return "", errors.New("gotrue client is nil")
	}

	//create user supabase auto checks tor the user the token belongs to 
	user, err := client.WithToken(tokenString).GetUser() 
	if err != nil {
		logger.Error("failed to get user from token: %v", err)
		return "", err
	}

	return user.ID.String(), nil
}

// this function will extract claims from a JWT token as we only need the claims
// it returns a map of claims and error if any
func ExtractClaims(tokenString string) (map[string]interface{}, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//parse the token without validating the signature as we just want the claims
	token,_,err := new (jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil{
		logger.Error("failed to parse token: %v", err)
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}else{
		return nil , errors.New("invalid token claims")
	}
}
