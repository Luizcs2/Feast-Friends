//use this for file testing 

package main

import (
	"feast-friends-api/pkg/logger"
	"feast-friends-api/internal/utils"
	"fmt"
)

func main() {
	logger.Info("🧪 Testing Database Connection...")

	// Test 1: Connection
	logger.Info("\n1️⃣ Testing Connection()...")
	if err := utils.Connection(); err != nil {
		logger.Error("❌ Connection failed: %v", err)
	}
	logger.Info("✅ Connection successful!")

	// Test 2: ExecuteQuery
	logger.Info("\n2️⃣ Testing ExecuteQuery()...")
	rows, err := utils.ExecuteQuery("SELECT 1 as test_column")
	if err != nil {
		logger.Error("❌ ExecuteQuery failed: %v", err)
	} else {
		defer rows.Close()
		fmt.Println("✅ ExecuteQuery successful!")
		
		// Try to read the result
		for rows.Next() {
			var testValue int
			if err := rows.Scan(&testValue); err != nil {
				logger.Error("❌ Scan failed: %v", err)
			} else {
				logger.Info("📊 Query result: %d\n", testValue)
			}
		}
	}

	// Test 3: ExecuteNonQuery
	logger.Info("\n3️⃣ Testing ExecuteNonQuery()...")
	// Using a harmless SELECT as a non-query test (won't modify data)
	commandTag, err := utils.ExecuteNonQuery("SELECT NOW()")
	if err != nil {
		logger.Error("❌ ExecuteNonQuery failed: %v", err)
	} else {
		logger.Info("✅ ExecuteNonQuery successful! CommandTag: %s\n", commandTag.String())
	}

	// Test 4: Health Checks
	fmt.Println("\n4️⃣ Testing Health Checks...")
	
	if err := utils.SupabaseCheck(); err != nil {
		logger.Error("❌ Supabase health check failed: %v", err)
	} else {
		logger.Info("✅ Supabase health check passed!")
	}

	if err := utils.DBCheck(); err != nil {
		logger.Error("❌ Database health check failed: %v", err)
	} else {
		logger.Info("✅ Database health check passed!")
	}

	// Test 5: Connection Pool Stats
	fmt.Println("\n5️⃣ Connection Pool Stats:")
	if utils.DB != nil {
		stats := utils.DB.Stat()
		logger.Info("📈 Total connections: %d\n", stats.TotalConns())
		logger.Info("📈 Idle connections: %d\n", stats.IdleConns())	
		logger.Info("📈 Acquired connections: %d\n", stats.AcquiredConns())
	}

	// Cleanup
	logger.Info("\n🧹 Cleaning up...")
	if utils.DB != nil {
		utils.DB.Close()
		logger.Info("✅ Database connection closed")
	}

	logger.Info("\n🎉 All tests completed!")
}