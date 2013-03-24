package main

import (
	"flag"
	"gompress/format"
	"log"
	//"gompress/testing"
)

var (
	flagZip = flag.Bool("zip", false, "gompress -zip")
	flagCom = flag.Bool("c", false, "gompress -c")
	args    []string
)

func main() {

	flag.Parse()

	var allArgs = flag.Args()

	if *flagZip && *flagCom && (len(allArgs) > 1) {
		args = allArgs[1:len(allArgs)]
		format.CreateZip(allArgs[0], args)
	} else {
		log.Fatal("Invalid flags or arguments.")
	}

	// Testing suite
	/*
		// Begin the test	
		flag.Parse()
		tempArgs := flag.Args()

		// Test the arguments list
		//testing.StartTesting(tempArgs)

		// Test the sanitized file's name
		testing.TestSanitizedName(tempArgs)
	*/
}
