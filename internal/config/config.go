//Loads all env vars from .env into config struct which rest of app can use 

package config

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

)

//container which holds all configs app needs
type Config struct {
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
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

//returning the loaded app config
func Get() *Config {
	return &cfg
}

//holds config
func init(){
	//loads the .env file if it exists 
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (make sure its in the root)")
	}

	//populate cfg from env variables
	if err := envconfig.Process("", &cfg); err != nil {
		log.Printf("Failed to Load Config:%v", err)
	}

	if cfg.JWT.Expiration == "" {
		cfg.JWT.Expiration = "24H"
	} 

}