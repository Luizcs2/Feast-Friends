// TODO - Luiz
// 1. Define Config struct with all environment variables:
//    - DATABASE_URL, JWT_SECRET, SUPABASE_URL, SUPABASE_KEY
//    - SERVER_PORT, ENVIRONMENT (dev/prod)
// 2. Load config from .env file or environment
// 3. Validate required configuration values
// 4. Provide default values where appropriate
// 5. Export global config instance

package configs

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

)

type Config struct {
	Enviorment string `envconfig:"ENVIRONMENT" default:"dev"`
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8000"`
		GinMode string `envconfig:"GIN_MODE" default:"debug"`

	}
	Supabase struct{
		URL string `envconfig:"SUPABASE_URL" required:"true"`
		AKey string `envconfig:"SUPABASE_ANON_KEY" required:"true"`
		Skey string `envconfig:"SUPABASE_SERVICE_KEY" required:"true"`
	}
	JWT struct {
		Secret     string `envconfig:"JWT_SECRET" required:"true"`
		Expiration string `envconfig:"JWT_EXPIRATION" default:"7d"`
	}
	FileUpload struct {
		// Stored in bytes. 10485760 bytes = 10 MB
		MaxFileSize int64 `envconfig:"MAX_FILE_SIZE" default:"10485760"`
	}
	Logging struct {
		Level string `envconfig:"LOG_LEVEL" default:"debug"`
	}
}

//global var to hold the loaded config 
var cfg Config

//returning the loaded app cconfig 
func Get() Config{
	return cfg
}

//holds config
func init(){

}