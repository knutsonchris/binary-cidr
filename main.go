package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	color "github.com/logrusorgru/aurora"
)

// getBinary will return the binary representation of the IP as a string
func getBinary(ipaddr string) string {
	s := strings.Split(ipaddr, ".")
	lis := []string{}
	for str := range s {
		// parse each string chunk to a 64 bit int
		i, err := strconv.ParseInt(s[str], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		// append the base 2 representation of that int as a string to the lis
		lis = append(lis, fmt.Sprintf("%08b", i))
	}
	return (strings.Join(lis[:], "."))
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("please enter an ip address with cidr notation. Example: 192.168.2.2/16")
	}
	ipaddr := args[0]
	// seperate CIDR block from ip
	s := strings.Split(ipaddr, "/")
	ipaddr = s[0]
	CIDR := s[1]

	binaryIP := getBinary(ipaddr)

	// since there are also dots that we need to account for, add up to chars to the network counter
	networkDigits, _ := strconv.Atoi(CIDR)
	if networkDigits > 24 {
		networkDigits += 3
	} else if networkDigits > 16 {
		networkDigits += 2
	} else if networkDigits > 8 {
		networkDigits++
	}

	networkPart := binaryIP[:networkDigits]
	hostPart := binaryIP[networkDigits:]

	fmt.Print(color.Magenta("network"), color.Cyan("host"), "\n\n")

	fmt.Print("Binary representation\n")
	fmt.Print(color.Magenta(networkPart), "", color.Cyan(hostPart), "\n\n")

	// get the subnet in binary
	var re = regexp.MustCompile(`\d`)
	networkSubnet := re.ReplaceAllString(networkPart, "1")
	hostSubnet := re.ReplaceAllString(hostPart, "0")

	fmt.Println("Subnet mask")
	fmt.Print(color.Magenta(networkSubnet), "", color.Cyan(hostSubnet), "\n\n")

	fmt.Println("Network ID")
	fmt.Print(color.Magenta(networkPart), "", color.Cyan(hostSubnet), "\n\n")

	fmt.Println("Broadcast")
	hostBroadcast := re.ReplaceAllString(hostSubnet, "1")
	fmt.Print(color.Magenta(networkPart), "", color.Cyan(hostBroadcast), "\n\n")
}
