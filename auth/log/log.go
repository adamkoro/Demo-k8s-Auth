package log

import (
	"log"
	"os"
)

var (
	WarningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.LUTC)
	ErrorLogger   = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.LUTC)
	InfoLogger    = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.LUTC)
)

func IsError(err error) bool {
	return err != nil
}
