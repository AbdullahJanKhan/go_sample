package logger

import (
	"io/ioutil"
	"os"

	"github.com/abdullahjankhan/go_sample/config"

	log "github.com/sirupsen/logrus"
)

var logger_instance = log.New()
var filepath = "./log/logrus.log"

func SetPath(path string) {
	// sample code to set the path of the log file
	filepath = path
}

func Init(conf *config.GlobalConfig) {

	//getting the credentials in the conf file
	//setting the format of the logs to be a JSON one
	logger_instance.SetFormatter(&log.JSONFormatter{PrettyPrint: true})

	//getting the log level set in the configuration file
	logLevel, err := log.ParseLevel("info" /* pass the log level you want to set*/)
	//If the log level in conf file can't be parsed, log level should be the default info level
	if err != nil {
		logLevel = log.InfoLevel
	}
	//setting the log level
	logger_instance.SetLevel(logLevel)

	//If we want to throw logs on the network, like redis
	if "conf.LogEnvironment" == "network" {
		//we dont want to display logs in the output, just on the network
		logger_instance.SetOutput(ioutil.Discard)
	} else if "conf.LogEnvironment" == "local" { //If we want to throw logs into a local file

		logger_instance.SetOutput(os.Stdout)
		//setting it to a file writer
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err == nil {
			logger_instance.Out = file
		} else {
			logger_instance.Info("Failed to log to file, using default stderr")
		}
		logger_instance.Info("Logrus has been initiated")
	}
}

func Instance() *log.Logger {
	// return the logger instance to be used in files
	return logger_instance
}
