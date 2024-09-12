package iplookup

import (
	"Port-Scanner-/src/setup"
	"fmt"
	"net"
)

/*
* Determine and print (if applicable) a domain name from IP; Reverse DNS lookup
 */
func HostLookup(ip string) {
	successLog, _, errorLog := setup.CreateStdoutLog()
	resp, err := net.LookupAddr(ip)
	if err != nil {
		errorLog.Println("Not a valid IP address")
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
