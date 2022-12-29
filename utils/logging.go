package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// read and update a logfile, or create it
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// specify the output of the logfile
	mutiLogFile := io.MultiWriter(os.Stdout, logfile)

	// set format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(mutiLogFile)
}
