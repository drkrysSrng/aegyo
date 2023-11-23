package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func reverse_ip(ip_add string) {
	addr, err := net.LookupAddr(ip_add)
	fmt.Println(addr, err)
}
func get_subdomains(domain string) {
	subfinderOpts := &runner.Options{
		Threads:            10, // Thread controls the number of threads to use for active enumerations
		Timeout:            30, // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10, // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		// ResultCallback: func(s *resolve.HostEntry) {
		// callback function executed after each unique subdomain is found
		// },
		// ProviderConfig: "your_provider_config.yaml",
		// and other config related options
	}

	// disable timestamps in logs / configure logger
	log.SetFlags(0)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		log.Fatalf("failed to create subfinder runner: %v", err)
	}

	output := &bytes.Buffer{}
	// To run subdomain enumeration on a single domain
	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		log.Fatalf("failed to enumerate single domain: %v", err)
	}

	// To run subdomain enumeration on a list of domains from file/reader
	// file, err := os.Open("domains.txt")
	// if err != nil {
	// 	log.Fatalf("failed to open domains file: %v", err)
	// }
	// defer file.Close()
	// if err = subfinder.EnumerateMultipleDomainsWithCtx(context.Background(), file, []io.Writer{output}); err != nil {
	// 	log.Fatalf("failed to enumerate subdomains from file: %v", err)
	// }

	// print the output
	log.Println(output.String())
}

func main() {

	// Prompt the user to enter a domain
	fmt.Println("Enter an IP to analyse: ")

	// Read the input from the command line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Get the user input
	domain := scanner.Text()

	// Prompt the user to enter a domain
	fmt.Println("Enter a domain to analyse: ")

	scanner.Scan()

	// Get the user input
	example_ip := scanner.Text()

	reverse_ip(example_ip)

	get_subdomains(domain)
	https_get(domain)
	nuclei_get(domain)

}
