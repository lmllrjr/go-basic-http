package none_test

import (
	"log"
	"os"
)

var dbLogger *log.Logger

func init() {
	dbLogger = log.New(os.Stdout, "TEST: ", 0)
	dbLogger.SetFlags(log.Ldate | log.Ltime)
}
