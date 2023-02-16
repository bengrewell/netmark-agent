package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bengrewell/netmark-agent/core"
	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v3"
	"log"
)

var (
	version string = "0.1.0"
	build   string
	rev     string
	branch  string
)

func PrintUsageLine(parameter string, defaultValue interface{}, description string, units string, extra string) {
	yellow := color.New(color.FgHiYellow)
	cyan := color.New(color.FgHiCyan)
	red := color.New(color.FgHiRed)
	yellow.Printf("    %-22s", parameter)
	cyan.Printf("  %-14v", defaultValue)
	yellow.Printf("  %-36s", description)
	cyan.Printf("  %-10s", units)
	red.Printf("  %s\n", extra)
}

func Usage() (usage func()) {
	return func() {
		white := color.New(color.FgWhite)
		boldWhite := color.New(color.FgWhite, color.Bold)
		boldGreen := color.New(color.FgGreen, color.Bold)
		usageLineFormat := "    %-22s  %-14v  %s\n"
		boldGreen.Printf("[+] netmark-agent :: Version %v :: Build %v :: Rev %v :: Branch %v\n", version, build, rev, branch)
		boldWhite.Print("Usage: ")
		fmt.Printf("netmark-agent <flags> ...\n")
		boldGreen.Print("  General Options:\n")
		white.Printf(usageLineFormat, "Parameter", "Default", "Description")
		//yellow.Printf(usageLineFormat, "-h", false, "show this help output")
		PrintUsageLine("--h[elp]", false, "show this help output", "[flag]", "")
		PrintUsageLine("--json", false, "output machine readable json", "[flag]", "")
		PrintUsageLine("--pretty", false, "pretty print the json output", "", "")
		PrintUsageLine("--yaml", false, "output machine readable yaml", "[flag]", "")
		PrintUsageLine("--host", "", "ip address of netmark server", "", "")
		PrintUsageLine("--port", 443, "port of the netmark service", "", "")
		PrintUsageLine("--timeout", 5, "connection timeout", "secs", "")
	}
}

func main() {

	// Setup flags
	var jout = flag.Bool("json", false, "print json output")
	var jpretty = flag.Bool("pretty", false, "pretty print json output")
	var yout = flag.Bool("yaml", false, "print yaml output")
	var host = flag.String("host", "", "hostname or ip address of server")
	var port = flag.String("port", "443", "port of the service")
	var timeout = flag.Int("timeout", 5, "amount of time in seconds to wait")
	flag.Usage = Usage()
	flag.Parse()

	// Ensure both json and yaml output are not selected
	if *yout && *jout {
		log.Fatalln("json/yaml output are mutually exclusive only one may be set")
	}

	// Ensure that either host and/or json/yaml is selected
	if *host == "" && !(*yout || *jout) {
		flag.Usage()
		log.Fatalln("a host must be specified or a local output format [json/yaml]")
	}

	// Run the core and gather all the information from enabled modules
	output := core.CoreEngine.Run()

	// If host is set then upload to the host
	if *host != "" {
		_ = port
		_ = timeout
		log.Fatalln("host upload has not yet been implemented")
	}

	// If json output is set then output in json
	if *jout {
		var outbuf []byte
		var err error

		// If pretty is set then pretty print the json
		if *jpretty {
			outbuf, err = prettyjson.Marshal(output)
		} else {
			outbuf, err = json.Marshal(output)
		}

		// Check for errors
		if err != nil {
			log.Fatalf("error: %w\n", err)
		}

		// Print the output
		fmt.Printf("%s\n", outbuf)

	} else if *yout {

		var outbuf []byte
		var err error

		// Marshal as yaml
		outbuf, err = yaml.Marshal(output)

		// Check for errors
		if err != nil {
			log.Fatalf("error: %w\n", err)
		}

		// Print the output
		fmt.Printf("%s\n", outbuf)
	}
}
