package main

import (
	"flag"
	"fmt"
)

// CLI Flags
var (
	ShowPtr     = flag.Bool("show", false, "Display the response content")
	ProtocolPtr = flag.String("protocol", "", "Provide the protocol")
	IntervalPtr = flag.Int("interval", 300, "Provide interval over which to poll monitor in seconds")
	HostPtr     = flag.String("host", "", "Provide host's IP/FQDN")
	VersionPtr  = flag.Bool("version", false, "shows version of this tool")
	InsecurePtr = flag.Bool("insecure", false, "Accept self-signed certificate")
)

// OSArgs struct
type OSArgs struct {
}

// FlagFetcher An interface to fetch CLI flags' value
type FlagFetcher interface {
	GetShowFlag() bool
	GetProtocolFlag() string
	GetIntervalFlag() int
	GetHostFlag() string
	GetVersionFlag() bool
	GetInsecureFlag() bool
}

// CLIArgs struct
type CLIArgs struct {
	ARGS FlagFetcher
}

// GetInsecureFlag function
func (a OSArgs) GetInsecureFlag() bool {
	if flag.Parsed() {
		return *InsecurePtr
	}

	return false
}

// GetVersionFlag function
func (a OSArgs) GetVersionFlag() bool {
	if flag.Parsed() {
		return *VersionPtr
	}

	return false
}

// GetIntervalFlag function
func (a OSArgs) GetIntervalFlag() int {
	if flag.Parsed() {
		return *IntervalPtr
	}

	return -1
}

// GetProtocolFlag function
func (a OSArgs) GetProtocolFlag() string {
	if flag.Parsed() {
		return *ProtocolPtr
	}

	return ""
}

// GetHostFlag function
func (a OSArgs) GetHostFlag() string {
	if flag.Parsed() {
		return *HostPtr
	}

	return ""
}

// GetShowFlag function
func (a OSArgs) GetShowFlag() bool {
	if flag.Parsed() {
		return *ShowPtr
	}

	return false
}

// Validate CLI args
func Validate(c CLIArgs) bool {
	proto := c.ARGS.GetProtocolFlag()
	host := c.ARGS.GetHostFlag()
	version := c.ARGS.GetVersionFlag()
	interval := c.ARGS.GetIntervalFlag()

	// check for version flag
	if proto == "" || host == "" {
		if version == true {
			fmt.Printf("VERSION: %s", VERSION)
		} else {
			flag.PrintDefaults()
		}
		return false
	}

	// Validate interval
	if interval < 0 {
		flag.PrintDefaults()
		fmt.Printf("Invalid interval value provided")
		return false
	}

	// Validate protocol
	if !(proto == "http" || proto == "https") {
		fmt.Printf("Supported Protocols: %v", SupportedProtocols)
		return false
	}
	return true
}
