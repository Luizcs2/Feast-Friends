// 2. Validate JWT token signature and expiration
// 3. Extract claims from token (user ID, email, etc.)
// 4. Set token expiration (24 hours recommended)

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

//this function will validate a JWT token and return the user ID if valid
func ValidateToken(tokenString string) (userID string, err error) {

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	client := gotrue.New(url,key)
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

//this function will extract claims from a JWT token
func ExtractClaims(tokenString string) (map[string]interface{}, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//parse the token
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
