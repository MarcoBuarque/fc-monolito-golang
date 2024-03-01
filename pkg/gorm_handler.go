package pkg

import (
	"strings"

	"github.com/MarcoBuarque/fc-monolito-golang/constant"
	log "github.com/sirupsen/logrus"
)

func GormErrorHandler(msg string, err error) {
	if err == nil {
		log.Error("GormErrorHandler:  error is nil", msg)
		PanicException(constant.UnknownError)
	}

	errStr := err.Error()

	switch {
	case strings.Contains(errStr, "record not found"):
		PanicException(constant.DataNotFound)
	case strings.Contains(errStr, "constraint"):
		PanicException(constant.InvalidRequest)
	default:
		log.Error("GormErrorHandler: Cannot handler gorm error", msg, errStr)
		PanicException(constant.UnknownError)
	}

}
