package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const maxConcurrentConnections = 50

func scanPort(target string, port int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	sem <- struct{}{}
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		<-sem
		return
	}
	fmt.Printf("Port %d open\n", port)
	conn.Close()
	<-sem
}

func main() {
	if os.Args[1] == "--help" {
		fmt.Printf("goscanner - A fast and lightweight portscanner build with multithreading\n")
		fmt.Printf("\n")
		fmt.Printf("Usage: goscanner [target IP]\n")
		fmt.Printf("\n")
		fmt.Printf("--help to see this screen")
		return
	}

	targetIp := os.Args[1]

	startPort := 1
	endPort := 65535

	sem := make(chan struct{}, maxConcurrentConnections)
	var wg sync.WaitGroup

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(targetIp, port, &wg, sem)
	}

	wg.Wait()
	fmt.Printf("Scan complete")
}
