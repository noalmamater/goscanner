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
