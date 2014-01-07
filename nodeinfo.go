package main

import (
	"fmt"
	"github.com/3M3RY/go-nodeinfo"
	"github.com/spf13/cobra"
	"os"
)

var GetNodeinfoCmd = &cobra.Command{
	Use:   "getnodeinfo [ADDRESS/HOSTNAME]",
	Short: "resolves a host or address through NodeInfo",
	Long:  `Performs a lookup or reverse lookup to the NodeInfo database.`,
	Run:   getNodeinfo,
}

func getNodeinfo(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "a single host must be specified")
		os.Exit(1)
	}
	/*
		if len(args) == 0 {
			faces, err := net.Interfaces()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to enumerate network interfaces, a host must be specified manually.")
			}

			for _, face := range faces {
				fmt.Println(face.Name())
				addrs, err := face.Addrs()
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to get addresses for %s, %s", face.Name, err)
				}
				for _, addr := range addrs {
					fmt.Println(addr)
					a := addr.String()
					if ipRegexp.MatchString(a) {
						entries, err := nodeinfo.LookupHost(args[0])
						if err != nil {
							fmt.Fprintf(os.Stderr, "failed to query %s, %s", addr, err)
						} else {
							for _, entry := range entries {
								fmt.Println(entry)
							}
						}
					}
				}
			}
			return
		}
	*/

	if ipRegexp.MatchString(args[0]) {
		hosts, err := nodeinfo.LookupAddr(args[0])
		if err != nil {
			fmt.Println(err)
		}
		for _, host := range hosts {
			fmt.Println(host)
		}
	} else {
		addrs, err := nodeinfo.LookupHost(args[0])
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {
			fmt.Println(addr)
		}
	}
}

var SetNodeinfoCmd = &cobra.Command{
	Use:   "setnodeinfo HOSTNAME",
	Short: "sets a NodeInfo hostname for this device",
	Long:  `Creates a hostname entry for the current device in the NodeInfo database.`,
	Run:   setNodeinfo,
}

func setNodeinfo(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "a single hostname must be specified")
		os.Exit(1)
	}
	if err := nodeinfo.SetHostname(args[0]); err != nil {
		fmt.Fprintln(os.Stderr, "failed to set NodeInfo hostname,", err)
		os.Exit(1)
	}
}