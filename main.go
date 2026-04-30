package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		mail := scanner.Text()
		fmt.Println("--------------------------------------------------------------")
		fmt.Printf("mail |  hasMX |  hasSPF |  spfRecord | hasDMARC |  dmarcRecord\n")
		fmt.Println("--------------------------------------------------------------")
		checkDomain(mail)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error could not read from the input: %v \n", err)
	}
}

func checkDomain(mail string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//LookupMX returns the DNS MX records for the given domain name sorted by preference

	mxRecords, err := net.LookupMX(mail)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	//LookupTXT returns the DNS TXT records for the given domain name.

	txtRecords, err := net.LookupTXT(mail)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + mail)
	if err != nil {
		log.Printf("Error: %v\n", err)

	}

	for _, value := range dmarcRecords {
		if strings.HasPrefix(value, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = value
		}
	}

	fmt.Printf("| MX Record    | Status: %-5v | Records: %v\n", hasMX, mxRecords)
	fmt.Printf("| SPF Record   | Status: %-5v | Value: %s\n", hasSPF, spfRecord)
	fmt.Printf("| DMARC Record | Status: %-5v | Value: %s\n", hasDMARC, dmarcRecord)
}
