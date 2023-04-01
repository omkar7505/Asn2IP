#!/usr/bin/env python

import argparse
import socket

def whois(ip):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect(('whois.radb.net', 43))
    s.send(ip.encode() + b'\r\n')
    response = b''
    while True:
        data = s.recv(4096)
        response += data
        if not data:
            break
    s.close()
    return response.decode()

def get_ip_blocks(asn):
    if asn.startswith('AS'):
        asn = asn[2:]
    output = whois(f'-i origin {asn}')
    ip_blocks = []
    for line in output.split('\n'):
        if line.startswith('route:'):
            ip_blocks.append(line.split()[1])
    return ip_blocks

def main():
    parser = argparse.ArgumentParser(description='Get IP blocks for an ASN')
    parser.add_argument('asn', metavar='ASN', type=str, help='ASN or file containing ASNs')
    args = parser.parse_args()
    
    try:
        with open(args.asn) as f:
            asns = f.read().splitlines()
    except FileNotFoundError:
        asns = [args.asn]
    
    for asn in asns:
        try:
            ip_blocks = get_ip_blocks(asn)
            #print(f'{asn}:')
            for ip_block in ip_blocks:
                print(ip_block)
        except Exception as e:
            print(f'Error getting IP blocks for {asn}: {e}')

if __name__ == '__main__':
    main()
