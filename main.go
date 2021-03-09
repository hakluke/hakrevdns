package main

import (
	"bufio"
	"context"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"net"
	"os"
	"strings"
	"sync"
)

var opts struct {
	Threads    int    `short:"t" long:"threads" default:"8" description:"How many threads should be used"`
	ResolverIP string `short:"r" long:"resolver" description:"IP of the DNS resolver to use for lookups"`
	Protocol   string `short:"P" long:"protocol" choice:"tcp" choice:"udp" default:"udp" description:"Protocol to use for lookups"`
	Port       uint16 `short:"p" long:"port" default:"53" description:"Port to bother the specified DNS resolver on"`
	Domain     bool   `short:"d" long:"domain" description:"Output only domains"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	// default of 8 threads
	numWorkers := opts.Threads

	work := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			work <- s.Text()
		}
		close(work)
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go doWork(work, wg)
	}
	wg.Wait()
}

func doWork(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	resolvers := make([]*net.Resolver, 1)
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

	resolvers = append(resolvers, r)

	for ip := range work {
		for _, r := range resolvers {
			addr, err := r.LookupAddr(context.Background(), ip)
			if err != nil {
				continue
			}

			for _, a := range addr {
				if opts.Domain {
					fmt.Println(strings.TrimRight(a, "."))
				} else {
					fmt.Println(ip, "\t", strings.TrimRight(a, "."))
				}
			}
		}
	}
}
