package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

)

var normalLog *log.Logger

var ErrorLog *log.Logger


openLogfile, err := os.OpenFile(filePath + "log.log"),
os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
	fmt.Println("Error opening log file: ", err)
	os.Exit(1)
}