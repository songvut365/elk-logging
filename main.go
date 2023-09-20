package main

import (
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	FILE_PATH = "./elk-stack/ingest_data/out.log"
)

func main() {
	// set format
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "severity",
		},
	})
	logrus.SetLevel(logrus.TraceLevel)

	// new file writer
	file, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Panicf("open file error: %v", err)
	}

	// write log to terminal and file
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	defer file.Close()

	// logging
	refId := logrus.Fields{
		"ref-id": uuid.NewString(),
	}

	logrus.WithFields(refId).Info("service is starting...")
	logrus.WithFields(refId).Warn("wait a minute, something's wrong here")
	logrus.WithFields(refId).Error("oh!! service error")
}
