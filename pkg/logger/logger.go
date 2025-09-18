// TODO - Luiz
// 4. File rotation for log files
// 5. Export logger functions: Info, Error, Debug, Warn	

// Package logger provides a centralized logging setup for the Feast Friends API,
// configuring log formatting and levels based on environment and application settings.
package logger

import (
	"github.com/sirupsen/logrus"
	"feast-friends-api/internal/config"
)

var Log = logrus.New()

// init configures the global logger instance (Log) based on application configuration,
// setting the log level and formatting according to the environment (development or production).
// It is automatically invoked when the package is imported.
func init() {
	
	// load config, get log level and env
	cfg := config.Get()
	levelString := cfg.Logging.Level
	env := cfg.Environment

	// here we set up formatting for dev or production based on environment string
	if env == "development" {
		Log.SetFormatter(&logrus.TextFormatter{
			ForceColors:    true,
			FullTimestamp: true,
			PadLevelText:  true,
			TimestampFormat: "15:04:05",
		})
		Log.SetReportCaller(true)
	} else {
		Log.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: false, 
		})
		Log.SetReportCaller(false)

	}

	// try to parse the string into logrus level
	level, err := logrus.ParseLevel(levelString)

	// if parsing fails log a warning and use default level	
	if err != nil {
		Log.WithFields(logrus.Fields{
			"provided": levelString,
			"error":    err,
		}).Warn("Invalid log level provided, falling back to info")
		level = logrus.InfoLevel
	}

	// set final level 
	Log.SetLevel(level)

	Log.WithFields(logrus.Fields{
		"env" : env,
		"level": level,
	}).Info("logger initialized")

}

