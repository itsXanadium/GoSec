package tools

// Version 1.0
import (
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Go Routine, create a job for the threads to perform portscanner, making it faster to get the Scan result
func Worker(host string, ports, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		address := net.JoinHostPort(host, strconv.Itoa(p))
		// address := fmt.Sprintf("example.com:%d", websitename, p)
		conn, err := net.DialTimeout("tcp", address, time.Millisecond*300)
		if err != nil {
			results <- 0
			continue
		}
		_ = conn.Close()
		results <- p
	}
}

func Portscanner(portscan string) string {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var wg sync.WaitGroup
	const numWorkers = 100
	const maxport = 1024

	// // reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("{/}Input the Website Name or IP: ")
	// RawHost, _ := reader.ReadString('\n')
	// host := strings.TrimSpace(RawHost)
	// //Check if the input Empty or not
	// if host == "" {
	// 	fmt.Println("{-}No Input Provided...")
	// 	return
	// }
	//Checks if the Website/Port exists
	WebHost, err := net.LookupHost(portscan)
	if err != nil {
		log.Printf("{-}Website/Port not found/invalid\n %s:%v\n", WebHost, err)
	}

	//Handles the worker
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(portscan, ports, results, &wg)
	}
	//Scanning till max Port
	go func() {
		for i := 1; i <= maxport; i++ {
			ports <- i
		}
		close(ports)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	for port := range results {
		if port != 0 {
			openports = append(openports, port)
		}
	}
	fmt.Println("Scanning Result: ")
	sort.Ints(openports)
	for _, port := range openports {
		log.Printf("{+}Port: %d open\n", port)
	}
	//Outputs
	sort.Ints(openports)
	return fmt.Sprintf("**port opens **, on:  %s:\n```%v```", portscan, openports)
}
