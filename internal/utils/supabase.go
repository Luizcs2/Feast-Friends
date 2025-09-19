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

import (
	"context"
	"feast-friends-api/internal/config"
	"feast-friends-api/pkg/logger"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/supabase-community/supabase-go"
)

var DB *pgxpool.Pool

func connection() error {
	cfg := config.Get()
	connStr, sKey := cfg.Supabase.URL, cfg.Supabase.Skey

	SupaClient, err := supabase.NewClient(connStr, sKey, nil)

	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		logger.Log.Errorf("Failed to connect to supabase: %v", err)
		return err
	}

	DB = dbpool
	return nil

}
