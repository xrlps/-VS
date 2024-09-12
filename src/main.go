package main

import (
	"Port-Scanner-/src/iplookup"
	"Port-Scanner-/src/portscanner"
	"Port-Scanner-/src/setup"
	"fmt"
	"os"
	"os/exec"
)

/*
* Clear screen. Only works on Windows!
 */
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearScreen()
		var selection string
		logo := setup.CreateLogo()
		fmt.Println(logo)
		fmt.Println("-----------------------------------------------------------------------------------")

		fmt.Println("OPTIONS\n\n[1] Port Scanner\n[2] DNS Lookup\n[3] Hostname Lookup")
		fmt.Print("Select: ")
		fmt.Scan(&selection)

		var addr string
		fmt.Print("Target: ")
		fmt.Scan(&addr)

		switch selection {
		case "1":
			portscanner.PortScanOption(addr)
		case "2":
			iplookup.DnsLookup(addr)
		case "3":
			iplookup.HostLookup(addr)
		}

	}
}
