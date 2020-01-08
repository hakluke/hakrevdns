package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	ResolverIP string `short:"r" long:"resolver" description:"IP of the DNS resolver to use for lookups"`
	Protocol   string `short:"P" long:"protocol" choice:"tcp" choice:"udp" default:"udp" description:"Protocol to use for lookups"`
	Port       uint16 `short:"p" long:"port" default:"53" description:"Port to bother the specified DNS resolver on"`
}

func worker(ip string, wg *sync.WaitGroup) {
	defer wg.Done()

	if opts.ResolverIP != "" {
		r := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, opts.Protocol, fmt.Sprintf("%s:%d", opts.ResolverIP, opts.Port))
			},
		}

		addr, err := r.LookupAddr(context.Background(), ip)
		if err == nil {
			fmt.Printf("%s\t%s\n", ip, addr[0])
		} else {
			fmt.Println(err)
		}
	} else {
		addr, err := net.LookupAddr(ip)
		if err == nil {
			fmt.Printf("%s\t%s\n", ip, addr[0])
		}
	}
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

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
