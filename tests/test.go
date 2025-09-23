package main

import (
	"fmt"
	"log"
	"feast-friends-api/internal/utils"
)

func main() {
	// Example token from Supabase login
	testToken := "eyJhbGciOiJIUzI1NiIsImtpZCI6Ik5JcGIyOXZmMzFTMFI5QXAiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2tza2Fsb2xwbWNlY3B6b3ZtdWV0LnN1cGFiYXNlLmNvL2F1dGgvdjEiLCJzdWIiOiJkMjk4OGIyMS03N2RmLTQzOWMtOWE5MS1kMGRjMjQzZWNjMzQiLCJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNzU4NTY1ODU3LCJpYXQiOjE3NTg1NjIyNTcsImVtYWlsIjoibHVpem9odG8yMDEyQGdtYWlsLmNvbSIsInBob25lIjoiIiwiYXBwX21ldGFkYXRhIjp7InByb3ZpZGVyIjoiZW1haWwiLCJwcm92aWRlcnMiOlsiZW1haWwiXX0sInVzZXJfbWV0YWRhdGEiOnsiZW1haWxfdmVyaWZpZWQiOnRydWV9LCJyb2xlIjoiYXV0aGVudGljYXRlZCIsImFhbCI6ImFhbDEiLCJhbXIiOlt7Im1ldGhvZCI6InBhc3N3b3JkIiwidGltZXN0YW1wIjoxNzU4NTYyMjU3fV0sInNlc3Npb25faWQiOiJmOGNlYTliYS1mYjJkLTRmMTUtOWRhZC02OTg3NjQwOWVhMTkiLCJpc19hbm9ueW1vdXMiOmZhbHNlfQ.P5qw-_gkCAUZhSAIGZNT6PVsHDHRpmfwr6QHZdLBrSs"

	// 1. Validate token
	userID, err := utils.ValidateToken(testToken)
	if err != nil {
		log.Fatalf("Token invalid: %v", err)
	}
	fmt.Println("Token valid! UserID:", userID)

	// 2. Extract claims (optional)
	claims, err := utils.ExtractClaims(testToken)
	if err != nil {
		log.Fatalf("Failed to extract claims: %v", err)
	}
	fmt.Println("Claims:", claims)
}
