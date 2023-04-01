package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func whois(ip string) string {
	conn, err := net.Dial("tcp", "whois.radb.net:43")
	if err != nil {
		return ""
	}
	defer conn.Close()

	fmt.Fprintf(conn, "-i origin %s\r\n", ip)
	scanner := bufio.NewScanner(conn)
	var response []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "route:") {
			response = append(response, line)
		}
	}
	return strings.Join(response, "\n")
}

func getIPBlocks(asn string) []string {
	if strings.HasPrefix(asn, "AS") {
		asn = asn[2:]
	}

	output := whois(asn)
	var ipBlocks []string
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "route:") {
			ipBlocks = append(ipBlocks, strings.Fields(line)[1])
		}
	}
	return ipBlocks
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h"{
		fmt.Println("Usage: Asn2IP <ASN> | <file>")
		fmt.Println("")
    		fmt.Println("Asn2IP is a tool for converting an ASN to a list of IP blocks.")
		fmt.Println("")
		fmt.Println("Arguments:")
    		fmt.Println("  <ASN>   The ASN to convert to IP blocks.")
    		fmt.Println("  <file>  A file containing a list of ASNs to convert to IP blocks.")
		os.Exit(1)
	}

	var asns []string
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			asns = append(asns, arg)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			asns = append(asns, scanner.Text())
		}
	}

	for _, asn := range asns {
		ipBlocks := getIPBlocks(asn)
		//fmt.Printf("%s:\n", asn)
		for _, ipBlock := range ipBlocks {
			fmt.Println(ipBlock)
		}
	}
}
