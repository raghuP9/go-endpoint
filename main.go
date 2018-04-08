package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	tpt "github.com/rpsraghu/go-endpoint/transport"
)

// Global vars
var (
	SupportedProtocols = [2]string{"http", "https"}
)

// Constants to be used
const (
	VERSION = "0.1.0"
)

func main() {
	showPtr := flag.Bool("show", false, "Display the response content")
	protocolPtr := flag.String("protocol", "", "Provide the protocol")
	intervalPtr := flag.Int("interval", 300, "Provide interval over which to poll monitor in seconds")
	hostPtr := flag.String("host", "", "Provide host's IP/FQDN")
	versionPtr := flag.Bool("version", false, "shows version of this tool")
	insecurePtr := flag.Bool("insecure", false, "Accept self-signed certificate")

	if len(os.Args) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Parse args
	flag.Parse()

	if flag.Parsed() {
		// check for version flag
		if *protocolPtr == "" || *hostPtr == "" {
			if *versionPtr == true {
				fmt.Printf("VERSION: %s", VERSION)
			} else {
				flag.PrintDefaults()
			}
			os.Exit(1)
		}

		// Validate interval
		if *intervalPtr < 0 {
			fmt.Printf("Negative interval value provided")
			os.Exit(1)
		}

		// Validate protocol
		if !(*protocolPtr == "http" || *protocolPtr == "https") {
			fmt.Printf("Supported Protocols: %v", SupportedProtocols)
			os.Exit(1)
		}
	}

	// setup signal catching
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	//signal.Notify(sigs,syscall.SIGQUIT)

	// method invoked upon seeing signal
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s", s)
		AppCleanup()
		os.Exit(1)
	}()

	// infinite print loop
	for {
		log.Printf("Starting Application!")
		log.Printf("Use Ctrl+C to stop...")

		pass := make(chan string, 1)
		fail := make(chan string, 1)

    tp := tpt.NewTransport(*insecurePtr)
		client := new(Client)
		client.Http = RealWebClient{Transport: tp}

		go monitor(client, *protocolPtr, *hostPtr, *showPtr, pass, fail)
		select {
		case msg := <-pass:
			log.Printf("Successfully monitored %s://%s",
				*protocolPtr, *hostPtr)
      log.Printf(msg)
      log.Println("Duration:", tp.Duration())
      log.Println("Request duration:", tp.ReqDuration())
      log.Println("Connection duration:", tp.ConnDuration())

		case msg := <-fail:
			log.Printf("Failed to monitor %s://%s", *protocolPtr, *hostPtr)
      log.Printf("Error message: %s", msg)
		}

		// wait random number of milliseconds
		Nsecs := *intervalPtr
		log.Printf("About to sleep %ds before looping again", Nsecs)
		time.Sleep(time.Second * time.Duration(Nsecs))
	}

}

// AppCleanup performs cleanup
func AppCleanup() {
	log.Println("CLEANUP APP BEFORE EXIT!!!")
}

func monitor(client *Client, protocol, host string, show bool,
  pass, fail chan string) {

	body, err := client.Http.Get(protocol + "://" + host)
	if err != nil {
		log.Printf("get error: %s: %s", err, host)
		fail <- err.Error()
		return
	}

	if show {
		log.Printf(string(body))
	}

	pass <- "Success"
	return
}
