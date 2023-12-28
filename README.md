# Asn2IP
Asn2IP is a command-line tool that converts Autonomous System Numbers (ASNs) to their associated IP blocks. An ASN is a unique identifier assigned to an Autonomous System (AS) by the Internet Assigned Numbers Authority (IANA). An AS is a collection of IP networks and routers under the control of one entity that presents a common routing policy to the Internet.

This repository contains both a Go and a Python implementation of the tool. The Go version was created to provide faster performance compared to the Python version.

## Installation
If you have Go installed, you can install Asn2IP using the following command:
```sh
go install github.com/Omkar7505/Asn2IP
```

## Usage
```sh
Asn2IP -h
```
This will display help for the tool.
```
Asn2IP <ASN> | <file>
Arguments:

  <ASN>: The ASN to convert to IP blocks.
  <file>: A file containing a list of ASNs to convert to IP blocks.
```

## Examples
### Converting a single ASN
To convert ASN AS55023 to its associated IP blocks, run the following command:
```sh
$ Asn2IP AS55023
```
This will output the list of IP blocks associated with ASN AS55023.

### Converting multiple ASNs from a file
To convert a list of ASNs stored in a file named asns.txt, run the following command:
```sh
Asn2IP asns.txt
```
This will output the list of IP blocks associated with each ASN listed in the asns.txt file.

Asn2IP can be useful in various applications, including network security and bug bounty hunting. By converting an ASN to its associated IP blocks, you can gain valuable information about the target’s infrastructure. This information can then be used for reconnaissance and vulnerability discovery.
