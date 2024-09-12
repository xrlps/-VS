package setup

import (
	"log"
	"os"
)

/*
* Write logs to a "vulnScan.log" file for a given message
 */
func Log(message string) {
	f, err := os.OpenFile("vulnScan.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	writeLog := log.New(f, message, log.Ldate|log.Ltime|log.Lshortfile)
	writeLog.Println(message)
}
