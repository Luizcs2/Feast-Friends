// logger.go 
// Package logger provides a centralized logging setup for the Feast Friends API,
// configuring log formatting and levels based on environment and application settings.
package logger

import (
	"feast-friends-api/internal/config"
	"io"
	"os"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log = logrus.New()


// setting up wrapper functions for different log levels 
// uses interface{} to accept any type of arguments
func Info(format string, args ...interface{}){
	Log.Infof(format,args...)
}
func Error(format string ,args ...interface{}){
	Log.Errorf(format,args...)
}
func Debug(format string , args ...interface{}){
	Log.Debugf(format,args...)
}
func Warn(format string, args ... interface{}){
	Log.Warnf(format,args...)
}

// init configures the global logger instance (Log) based on application configuration,
// setting the log level and formatting according to the environment (development or production).
// It is automatically invoked when the package is imported.
func init() {
	
	// load config, get log level and env
	cfg := config.Get()
	levelString := cfg.Logging.Level
	env := cfg.Environment

	lumberjackLogger := &lumberjack.Logger{
		Filename:  ".logs/app.log",
		MaxSize:10, // megabytes
		MaxBackups: 3,
		MaxAge: 28, //days
		Compress: true, // disabled by default
	}

	// here we set up formatting for dev or production based on environment string
	if env == "development" {
		Log.SetFormatter(&logrus.TextFormatter{
			ForceColors:    true,
			FullTimestamp: true,
			PadLevelText:  true,
			TimestampFormat: "15:04:05",
		})
		Log.SetReportCaller(true)
		multiWriter := io.MultiWriter(os.Stdout,lumberjackLogger)
		Log.SetOutput(multiWriter)

	} else {
		Log.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: false, 
		})
		Log.SetReportCaller(false)
		Log.SetOutput(lumberjackLogger)

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

