package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: hexip [hostname, ip, or 32-bit hex number]")
		os.Exit(2)
	}

	input := os.Args[1]

	var ip net.IP
	if matched, _ := regexp.MatchString(`^0x[0-9a-f]+$`, input); matched {
		ip = net.ParseIP(hex2ip(input))
	} else {
		ip = net.ParseIP(input)
	}
	ip = ip.To4()
	if ip == nil {
		ips, err := net.LookupIP(input)
		checkErr(err)
		for _, ip := range ips {
			if ip4 := ip.To4(); ip4 != nil {
				fmt.Println(input, ip4, ip2hex(ip4))
			} else if ip6 := ip.To16(); ip6 != nil {
				fmt.Println(input, ip6, "(cannot convert ip6 to hex at this time...)")
			} else {
				fmt.Printf("could not convert: %v\n", ip)
			}
		}
	} else {
		hex := ip2hex(ip)
		ip4 := ip.String()
		if names, err := net.LookupAddr(ip4); err == nil {
			for _, host := range names {
				fmt.Println(strings.TrimRight(host, "."), ip4, hex)
			}
		} else {
			fmt.Println(err)
			fmt.Println(input, ip4, hex)
		}
	}
}

func hex2ip(hex string) string {
	ip, err := strconv.ParseInt(hex, 0, 0)
	checkErr(err)
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		int(ip)>>24&0xff,
		int(ip)>>16&0xff,
		int(ip)>>8&0xff,
		int(ip)&0xff,
	)
}

func ip2hex(ip4 net.IP) string {
	return fmt.Sprintf(
		"0x%x",
		(int(ip4[0])<<24)|(int(ip4[1])<<16)|(int(ip4[2])<<8)|int(ip4[3]),
	)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
