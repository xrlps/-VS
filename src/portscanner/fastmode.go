package portscanner

import (
	"Port-Scanner-/src/setup"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

// Create the global variable for logging
var logger *log.Logger
var successLog *log.Logger

func init() {
	successLog, logger, _ = setup.CreateStdoutLog()
}

/*
* Creates a tcp connection to check for open ports
 */
func worker(ports chan int, addr string, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports { // extracts the "first" value available in the port channel
		fullUrl := addr + strconv.Itoa(p)
		conn, err := net.Dial("tcp", fullUrl)
		if err != nil {
			continue
		}
		successLog.Printf(fullUrl)
		message := fmt.Sprintf("success %s", fullUrl)
		setup.Log(message)
		conn.Close()
	}
}

/*
* Runs goroutines for scans
 */
func startWorkers(url string, workerCount int, wg *sync.WaitGroup, numOfPorts int) {
	var ports = make(chan int, workerCount)

	// Run the worker functions concurrently
	for i := 0; i < cap(ports); i++ {
		wg.Add(1) // Add a wait group for every worker
		go worker(ports, url, wg)
	}

	// Create the ports available
	go func() {
		for i := 0; i <= numOfPorts; i++ {
			ports <- i
		}
		close(ports)
	}()
}

func concurrentRequest(url string, wg *sync.WaitGroup, workers, numOfPorts int) {
	url = url + ":" // Format the url to append port #'s

	startWorkers(url, workers, wg, numOfPorts)
}

func normalRequest(url string, numOfPorts int) {
	url = url + ":" // Format the url to append port #'s

	for i := 0; i <= numOfPorts; i++ {
		fullUrl := url + strconv.Itoa(i)
		conn, err := net.Dial("tcp", fullUrl)
		if err != nil {
			continue
		}
		successLog.Print(fullUrl)
		conn.Close()
	}
}

/*
* Handles option inputs
 */
func PortScanOption(url string) {
	var wg sync.WaitGroup
	var selection string
	fmt.Print("[1] Fast Mode \n[2] Safe Mode\nSelect: ")
	fmt.Scan(&selection)
	var numOfPorts int
	var err error

	for {
		var portCount string
		fmt.Print("Enter the range of ports to scan (i.e. 1023): ")
		fmt.Scan(&portCount)
		numOfPorts, err = strconv.Atoi(portCount)
		if err == nil {
			break
		}
	}

	if selection == "1" {
		var workers int
		for {
			var workerCount string
			fmt.Print("Workers: ")
			fmt.Scan(&workerCount)
			workers, err = strconv.Atoi(workerCount)
			if err == nil {
				break
			}
		}
		logger.Println("Running tasks...")
		logger.Println(numOfPorts)
		concurrentRequest(url, &wg, workers, numOfPorts)
		wg.Wait()
	} else if selection == "2" {
		normalRequest(url, numOfPorts)
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
