# hakrevdns
Small, fast, simple tool for performing reverse DNS lookups en masse.

You feed it IP addresses, it returns hostnames.

This can be a useful way of finding domains and subdomains belonging to a company from their IP addresses.

# Installation
```
go get github.com/hakluke/hakrevdns
```

# Usage
Pipe a list of IP addresses into the tool, for example:

```
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
