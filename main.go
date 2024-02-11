package main

import (
	"flag"
	"fmt"
	stdnet "net"

	"github.com/shirou/gopsutil/net"
)

func getInterfaces() ([]stdnet.Interface, error) {
	return stdnet.Interfaces()
}

func getAddresses(iface stdnet.Interface) ([]stdnet.Addr, error) {
	return iface.Addrs()
}

func getIOCounters() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}

func printInterfaceStats(iface stdnet.Interface, ioCounter net.IOCountersStat) {
	addrs, err := getAddresses(iface)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		fmt.Printf("Name: %v, MTU: %v, Flags: %v, BytesSent: %v, BytesRecv: %v, Address: %v\n",
			iface.Name, iface.MTU, iface.Flags, ioCounter.BytesSent, ioCounter.BytesRecv, addr.String())
	}
}

func printNetworkStats() {
	interfaces, err := getInterfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ioCounters, err := getIOCounters()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, iface := range interfaces {
		for _, ioCounter := range ioCounters {
			if ioCounter.Name == iface.Name {
				printInterfaceStats(iface, ioCounter)
			}
		}
	}
}

func main() {
	showStats := flag.Bool("stats", false, "Show network statistics")
	help := flag.Bool("help", false, "Show help information")

	flag.Parse()

	if *help {
		fmt.Println("Usage of CLI tool:")
		fmt.Println("  -stats: Show network statistics")
		fmt.Println("  -help: Show help information")
		return
	}

	if *showStats {
		printNetworkStats()
	} else {
		fmt.Println("Use -stats flag to show network statistics")
	}
}
