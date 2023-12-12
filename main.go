package main

import (
    "bufio"
    "context"
    "fmt"
    "net"
    "os"
    "strings"
    "sync"

    flags "github.com/jessevdk/go-flags"
)

var opts struct {
    Threads      int      `short:"t" long:"threads" default:"8" description:"How many threads should be used"`
    ResolverIP   string   `short:"r" long:"resolver" description:"IP of the DNS resolver to use for lookups"`
    Resolvers    []string `short:"R" long:"resolvers" description:"List of DNS resolver IPs to use for lookups, separated by comma"`
    ResolverFile string   `short:"f" long:"resolver-file" description:"File containing list of DNS resolver IPs"`
    Protocol     string   `short:"P" long:"protocol" choice:"tcp" choice:"udp" default:"udp" description:"Protocol to use for lookups"`
    Port         uint16   `short:"p" long:"port" default:"53" description:"Port to bother the specified DNS resolver on"`
    Domain       bool     `short:"d" long:"domain" description:"Output only domains"`
}

func main() {
    _, err := flags.ParseArgs(&opts, os.Args)
    if err != nil {
        os.Exit(1)
    }

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

    resolvers, err := getResolvers()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading resolvers:", err)
        return
    }

    for ip := range work {
        var addr []string
        for _, resolverIP := range resolvers {
            r := &net.Resolver{
                PreferGo: true,
                Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
                    d := net.Dialer{}
                    return d.DialContext(ctx, opts.Protocol, fmt.Sprintf("%s:%d", resolverIP, opts.Port))
                },
            }

            addr, err = r.LookupAddr(context.Background(), ip)
            if err == nil {
                break // Exit the loop if a successful lookup is achieved
            }
        }

        if err != nil {
            fmt.Fprintln(os.Stderr, "Lookup failed for IP:", ip, "Error:", err)
            continue
        }

        for _, a := range addr {
            if opts.Domain {
                fmt.Println(strings.TrimRight(a, "."))
            } else {
                fmt.Println(ip, "\t", a)
            }
        }
    }
}

func getResolvers() ([]string, error) {
    if opts.ResolverFile != "" {
        file, err := os.ReadFile(opts.ResolverFile)
        if err != nil {
            return nil, err
        }
        return strings.Split(strings.TrimSpace(string(file)), "\n"), nil
    } else if len(opts.Resolvers) > 0 {
        return opts.Resolvers, nil
    } else if opts.ResolverIP != "" {
        return []string{opts.ResolverIP}, nil
    }
    return []string{}, nil // or return a default resolver if you wish
}
