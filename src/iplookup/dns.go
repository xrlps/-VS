package iplookup

import (
	"Port-Scanner-/src/setup"
	"fmt"
	"net"
)

/*
* Determine and print (if applicable) an IP address for DNS; DNS lookup
 */
func DnsLookup(url string) {
	successLog, _, errorLog := setup.CreateStdoutLog()
	resp, err := net.LookupHost(url)
	if err != nil {
		errorLog.Println("Not a valid Host Name")
	} else {
		// Print the message to the console
		successLog.Println(resp)
		// Write to the log file
		message := fmt.Sprintf("success %s", resp)
		setup.Log(message)
	}

	for {
		var selection string
		fmt.Println("(1) Menu")
		fmt.Print("Select: ")
		fmt.Scan(&selection)
		if selection == "1" {
			return
		}
	}
}
