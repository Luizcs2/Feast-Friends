// TODO - Luiz
// 3. Add log formatting for development vs production
// 4. File rotation for log files
// 5. Export logger functions: Info, Error, Debug, Warn	
package logger

import (
	"log"
	"github.com/sirupsen/logrus"
	"feast-friends-api/internal/config"
)

var Log = logrus.New()

func init() {

	// get the log level string fron the config 
	levelString := config.Get().Logging.Level

	// try to parse the string into logrus level
	level, err := logrus.ParseLevel(levelString)

	// if parsing fails log a warning and use default level	
	if err != nil{
		log.Printf("Wrong LOG_LEVEL: %s , Error:%v ", levelString, err)
		level = logrus.InfoLevel
	}

	// set final level 
	Log.SetLevel(level)

	Log.Info("Logger has been init")

}

