// TODO - Luiz
// 1. Initialize Supabase client with URL and service key
// 2. Create database connection pool
// 3. Implement helper functions:
//    - ExecuteQuery(query, args) for SELECT statements
//    - ExecuteNonQuery(query, args) for INSERT/UPDATE/DELETE
// 4. Handle database connection errors
// 5. Add connection health check function
// 6. Configure connection timeouts and retries

package utils

//set importing necessary packages
import (
	"context"
	"feast-friends-api/internal/config"
	"feast-friends-api/pkg/logger"

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

	SupaClient, err := supabase.NewClient(connStr, Key, &supabase.ClientOptions{})
	if err != nil {
		logger.Log.Errorf("Failed to create supabase client: %v", err)
		return err
	}
	SupabaseClient = SupaClient

	// asking for connection pool && return error if any
	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		logger.Log.Errorf("Failed to connect to supabase: %v", err)
		return err
	}

	//ping to confirm if dbpool connected successfully
	if err = dbpool.Ping(context.Background()); err != nil {
		logger.Log.Errorf("Failed to ping supabase: %v", err)
		return err
	}

	DB = dbpool
	logger.Log.Info("Connected to Supabase successfully")
	return nil

}

// ExecuteQuery executes a query and returns the resulting rows
func ExecuteQuery(query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := DB.Query(context.Background(), query, args...)
	if err != nil {
		logger.Log.Errorf("query execution failed:%s : %v", query, err)
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
		logger.Log.Errorf("non-query failed: %s : %v", query, err)
		return nil, err
	}
	return commandTag, nil
}

// checks if supabase connection is alive
func SupabaseCheck() error {
	_, err := SupabaseClient.Auth.GetUser()
	if err != nil {
		logger.Log.Errorf("Supabase health check failed : %v", err)
		return err
	}
	logger.Log.Info("Supabase connection is healthy")
	return nil
}

//checks if db connection is alive
func DBCheck() error {
	if err := DB.Ping(context.Background()); err != nil {
		logger.Log.Errorf("Database health check failed : %v", err)
		return err
	}
	return nil
}