# hakrevdns

Small, fast, simple tool for performing reverse DNS lookups en masse.

You feed it IP addresses, it returns hostnames.

This can be a useful way of finding domains and subdomains belonging to a company from their IP addresses.

## Installation

```sh
go get github.com/hakluke/hakrevdns
```

### Alternative Installation using automated install script (Linux)

#### Note: Requires ```wget``` in order to download the golang tarball

  - Installs compatible version of golang
  - Builds hakrevdns and places the binary in ```/usr/local/bin```
  
```sh
chmod +x install.sh
./install.sh
```

## Usage
The most basic usage is to simply pipe a list of IP addresses into the tool, for example:

```sh
hakluke~$ prips 173.0.84.0/24 | hakrevdns 
173.0.84.110	he.paypal.com.
173.0.84.109	twofasapi.paypal.com.
173.0.84.114	www-carrier.paypal.com.
173.0.84.77	twofasapi.paypal.com.
173.0.84.102	pointofsale.paypal.com.
173.0.84.104	slc-a-origin-pointofsale.paypal.com.
173.0.84.111	smsapi.paypal.com.
173.0.84.203	m.paypal.com.
173.0.84.105	prm.paypal.com.
173.0.84.113	mpltapi.paypal.com.
173.0.84.8	ipnpb.paypal.com.
173.0.84.2	active-www.paypal.com.
173.0.84.4	securepayments.paypal.com.
...
```

### Parameters

```sh
hakluke~$ hakrevdns -h
Usage:
  hakrevdns [OPTIONS]

Application Options:
  -t, --threads=           Number of threads (too many may get you banned, too few will be slow)
  -r, --resolver=          IP of the DNS resolver to use for lookups
  -P, --protocol=[tcp|udp] Protocol to use for lookups (default: udp)
  -p, --port=              Port to bother the specified DNS resolver on (default: 53)
  -d, --domain             Output only domains

Help Options:
  -h, --help               Show this help message
```

If you want to use a resolver not specified by you OS, say: 1.1.1.1, try this:

```sh
hakluke~$ echo "173.0.84.110" | hakrevdns -r 1.1.1.1
173.0.84.110    he.paypal.com.
```

If you wish to obtain only a list of domains, perhaps to sort out working ones you may use:

```sh
$ echo "173.0.84.110" | hakrevdns -d
```
For filtering out only working domains, you may additionally pipe it through [httprobe](https://github.com/tomnomnom/httprobe):

```sh
$ echo "173.0.84.110" | hakrevdns -d | httprobe
```

## Contributors
- [hakluke](https://twitter.com/hakluke) wrote the tool
- [alphakilo](https://github.com/Alphakilo/) added the option to use custom resolvers
- [SaveBreach](https://twitter.com/SaveBreach/) added the -d flag and cleaned up the code
