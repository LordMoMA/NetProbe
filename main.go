package main

import (
	"flag"
	"fmt"
	stdnet "net"

	"github.com/fatih/color"
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
			color.CyanString("%v", iface.Name),
			color.BlueString("%v", iface.MTU),
			color.GreenString("%v", iface.Flags),
			color.YellowString("%v", ioCounter.BytesSent),
			color.MagentaString("%v", ioCounter.BytesRecv),
			color.RedString("%v", addr.String()))
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
		color.Cyan("Usage of CLI tool:")
		color.Blue("  -stats: Show network statistics")
		color.Green("  -help: Show help information")
		return
	}

	if *showStats {
		printNetworkStats()
	} else {
		color.Red("Use -stats flag to show network statistics")
	}
}
