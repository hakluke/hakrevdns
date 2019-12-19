# hakrevdns
Small, fast, simple tool for performing reverse DNS lookups en masse.
This can be a useful way of finding domains and subdomains belonging to a company from their IP address.

# Installation
```
go get github.com/hakluke/hakrevdns
```

# Usage
Pipe a list of IP addresses into the tool, for example:

```
hakluke@home:/tmp$ prips 172.217.167.0/24 > ips.txt
hakluke@home:/tmp$ cat ips.txt | hakrevdns 
172.217.167.97 [syd09s17-in-f1.1e100.net.]
172.217.167.109 [syd09s17-in-f13.1e100.net.]
172.217.167.98 [syd09s17-in-f2.1e100.net.]
172.217.167.110 [syd09s17-in-f14.1e100.net.]
172.217.167.77 [syd15s06-in-f13.1e100.net.]
172.217.167.99 [syd09s17-in-f3.1e100.net.]
172.217.167.100 [syd09s17-in-f4.1e100.net.]
172.217.167.115 [syd09s17-in-f19.1e100.net.]
172.217.167.116 [syd09s17-in-f20.1e100.net.]
172.217.167.111 [syd09s17-in-f15.1e100.net.]
.... etc
```
