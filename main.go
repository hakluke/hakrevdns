package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	ResolverIP string `short:"r" long:"resolver" description:"IP of the DNS resolver to use for lookups"`
	Protocol   string `short:"P" long:"protocol" choice:"tcp" choice:"udp" default:"udp" description:"Protocol to use for lookups"`
	Port       uint16 `short:"p" long:"port" default:"53" description:"Port to bother the specified DNS resolver on"`
	Domain 	   bool   `short:"d" long:"domain" description:"Output only domains"`
}

func worker(ip string, wg *sync.WaitGroup, res chan string) {
	defer wg.Done()

	var r *net.Resolver

	if opts.ResolverIP != "" {
		r = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, opts.Protocol, fmt.Sprintf("%s:%d", opts.ResolverIP, opts.Port))
			},
		}
	}

	addr, err := r.LookupAddr(context.Background(), ip)
	if err != nil {
		return
	}

	for _, a := range addr {
		if opts.Domain {
			res <- fmt.Sprintf("%s", strings.TrimRight(a, "."))
		} else {
			res <- fmt.Sprintf("%s \t %s", ip, a)
		}
	}
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

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
