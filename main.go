package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "sync"
)

func worker(ip string, wg *sync.WaitGroup, res chan string) {
    defer wg.Done()

    addr, err := net.LookupAddr(ip)
    if err != nil {
        return
    }

    for _, a := range addr {
        res <- fmt.Sprintf("%s \t %s", ip, a)
    }
}

func main() {
    var wg sync.WaitGroup
    res := make(chan string)

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        wg.Add(1)
        go worker(scanner.Text(), &wg, res)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    go func() {
        wg.Wait()
        close(res)
    }()

    for r := range res {
        fmt.Println(r)
    }
}
