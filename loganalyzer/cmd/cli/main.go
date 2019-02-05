package main

import (
	"flag"
	"fmt"
	"honestbee/pkg/analyzer"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Please enter a valid path to a log file.")
		return
	}

	// Setup flags
	reportType := flag.String("r", "minimal", "Show full or minimal report")
	entries := flag.Int("e", 10, "Number of logs to be reported")
	displayList := flag.Bool("l", true, "Displayed as list or not")
	noGif := flag.Bool("g", true, "No GIFs")
	status := flag.Bool("s", true, "Only display request with status 200")
	getMethod := flag.Bool("m", true, "Only display request with status 200")

	flag.Parse()

	// Initialize and start analyzer
	a := analyzer.New(
		*reportType,
		*entries,
		*status,
		*getMethod,
		*noGif,
	)

	// Get result and display
	result := a.Analyze(args[len(args)-1])
	if *displayList {
		fmt.Println(result)
	} else {
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}

}
