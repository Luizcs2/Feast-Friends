package utils

//set importing necessary packages
import (
	"context"
	"feast-friends-api/internal/config"
	"feast-friends-api/pkg/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/supabase-community/supabase-go"
)

var DB *pgxpool.Pool
var SupabaseClient *supabase.Client

func Connection() error {
	//loading in config
	cfg := config.Get()
	connStr := cfg.Supabase.URL
	Key := cfg.Supabase.Skey
	maxRetries := 3
	retryDelay := 2 * time.Second
	var lastErr []error
	var check error

	for i := 0; i < maxRetries; i++ {

		clientOptions := supabase.ClientOptions{}

		//attemping to conenct to supabase client
		SupaClient, err := supabase.NewClient(connStr, Key, &clientOptions)
		if err != nil {
			lastErr = append(lastErr, fmt.Errorf("failed to connect to supabase client: %v", err))
			logger.Warn("Supabase client creation failed on attempt %d: %v", i+1, err)
			time.Sleep(retryDelay)
			continue //trying until max retries reached
		}
		SupabaseClient = SupaClient

		//using health check to confirm connection is alive
		if err := SupabaseCheck(); err != nil {
			check = err
			logger.Warn("Supabase health check failed: %v", check)
			time.Sleep(retryDelay)
			continue //trying until max retries reached
		}

		// asking for connection pool && return error if any

		dbpool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("%s?connect_timeout=10", connStr))
		if err != nil {
			lastErr = append(lastErr, fmt.Errorf("failed to connect to DB: %v", err))
			logger.Warn("DB connection failed on attempt %d: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}
		DB = dbpool

		//using db health check to confirm connection is alive
		if err := DBCheck(); err != nil {
			check = err
			logger.Warn("DB health check failed: %v", check)
			time.Sleep(retryDelay)
			continue //trying until max retries reached
		}

		// if we reach here both connections are successful
		logger.Info("Supabase and DB connection successful")
		return nil
	}
	// if we reach here all retries failed
	return fmt.Errorf("failed to connect after %d attempts: %v", maxRetries, lastErr)

}

// ExecuteQuery executes a query and returns the resulting rows
func ExecuteQuery(query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := DB.Query(context.Background(), query, args...)
	if err != nil {
		logger.Error("query execution failed:%s : %v", query, err)
		return nil, err
	}
	return rows, nil
}

// Execute a non-query command (insert, update, delete)
// commandTag contains info about the executed command
// tells rows affected and command that was executed)
func ExecuteNonQuery(query string, args ...interface{}) (pgconn.CommandTag, error) {
	commandTag, err := DB.Exec(context.Background(), query, args...)
	if err != nil {
		logger.Error("non-query failed: %s : %v", query, err)
	}
	return commandTag, err
}

// checks if supabase connection is alive
func SupabaseCheck() error {
	_, err := SupabaseClient.Auth.GetUser()
	if err != nil {
		logger.Error("Supabase health check failed : %v", err)
		return err
	}
	logger.Info("Supabase connection is healthy")
	return nil
}

// checks if db connection is alive
func DBCheck() error {
	if err := DB.Ping(context.Background()); err != nil {
		logger.Error("Database health check failed : %v", err)
		return err
	}
	logger.Info("DB connection is healthy")
	return nil
}
