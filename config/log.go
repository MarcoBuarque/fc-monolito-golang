package config

import (
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func NewLog() {

	log.SetLevel(getLevel(os.Getenv("LOG_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})
}

func getLevel(value string) log.Level {
	switch value {
	case "DEBUG":
		return log.DebugLevel
	case "WARN":
		return log.WarnLevel
	case "INFO":
		return log.InfoLevel
	default:
		return log.ErrorLevel
	}
}
