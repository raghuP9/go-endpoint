package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"
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

	if len(os.Args) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Parse args
	flag.Parse()

	cArgs := new(CLIArgs)
	cArgs.ARGS = new(OSArgs)
	if !Validate(*cArgs) {
		os.Exit(1)
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
		//AppCleanup()
		os.Exit(1)
	}()

	log.Printf("Starting Application!")
	log.Printf("Use Ctrl+C to stop...")

	// infinite print loop
	for {
		pass := make(chan string, 1)
		fail := make(chan string, 1)

		tp := NewTransport(cArgs.ARGS.GetInsecureFlag())
		client := new(Client)
		client.HTTP = RealWebClient{Transport: tp}

		go monitor(client, cArgs.ARGS.GetProtocolFlag(),
			cArgs.ARGS.GetHostFlag(), cArgs.ARGS.GetShowFlag(),
			pass, fail)
		select {
		case msg := <-pass:
			log.Printf("Successfully monitored %s://%s",
				cArgs.ARGS.GetProtocolFlag(), cArgs.ARGS.GetHostFlag())
			log.Printf(msg)
			log.Println("Duration:", tp.Duration())
			log.Println("Request duration:", tp.ReqDuration())
			log.Println("Connection duration:", tp.ConnDuration())

		case msg := <-fail:
			log.Printf("Failed to monitor %s://%s",
				cArgs.ARGS.GetProtocolFlag(),
				cArgs.ARGS.GetHostFlag())
			log.Printf("Error message: %s", msg)
		}

		// wait random number of milliseconds
		Nsecs := cArgs.ARGS.GetIntervalFlag()
		log.Printf("About to sleep %ds before looping again", Nsecs)
		time.Sleep(time.Second * time.Duration(Nsecs))
	}

}

// AppCleanup performs cleanup
// func AppCleanup() {
//	log.Println("CLEANUP APP BEFORE EXIT!!!")
//}

func monitor(client *Client, protocol, host string, show bool,
	pass, fail chan string) {

	body, err := client.HTTP.Get(protocol + "://" + host)
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
