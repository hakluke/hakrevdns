package main

import (
    "fmt"
    "net"
	"sync"
	"bufio"
	"os"
)

func worker(ip string, wg *sync.WaitGroup) {
	addr, err := net.LookupAddr(ip)	
    fmt.Println(ip, addr, err)
	wg.Done()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	for scanner.Scan() {
		go worker(scanner.Text(), &wg)
		wg.Add(1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	wg.Wait()
}
