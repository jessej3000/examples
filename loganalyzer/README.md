#  Log file analyzer


# NOTE: 
This was created using Go version go1.11 darwin/amd64
This is a CLI application so we need to build it first.
Please ensure Go is properly installed and all its GOPATH settings

# To build the code
1. Download or clone my code "git clone https://github.com/jessej3000/examples.git"
2. You will only need the `loganalyzer` folder. Please delete all other folders."
3. To build my code just go into the directory `examples/loganalyzer/cmd/cli`
4. Then in the console type "go build"

# To run the code (NOTE: you need to build it first. Please see instructions above.)
1. To test my code just go into the directory `examples/loganalyzer/cmd/cli`
2. Then in the console type "./cli access.log". (access.log is a sample log file.)

# Unit test
1. To test my code just go into the directory `examples/loganalyzer/pkg/analyzer`
2. And run "go test". See querylog_test.go for details of testing.

# Files and directory structure
Structure - I separated the core logic from the main program by making it as a package so it can be used by other applications. The structure is also devided into two. cmd and pkg. 
cmd - may contain all types of applications like CLI, or web app, in their own folders. In this case it is a CLI application so it is in a CLI folder. If you need a web app, you can simply create a folder under the cmd folder like web, and have your main.go web files in there. And your web app can call the analyzer package.

Files
- main.go               - contains func main()
- querylog.go           - log analyzer package
- querylog_test.go      - testing for querlog package
- *.log                 - sample log files to process

# Some explanation 
- Command line flags

    	
	* -r = minimal | complete : Default minimal. Shows full or minimal report
	* -e = n where n is any positive integer : Default 10. Number of log entries to be reported
	* -l = true | false : Default true, Displayed as list or not
	* -g = true | false : Default true. No GIFs if set to true and shows GIF if false
	* -s = true | false : Default true. True only display request with status 200 and false otherwise
	* -m = true | false : Default true. Only display request with status 200. Will display other status if false

- New: Creates a new QueryLog object with the following parameters:
    	* reportType string : taken from flag -r
	* numberOfEntry int : taken from flag -e
	* status bool : taken from flag -s
	* getMethod bool : taken from flag -m
	* noGif bool : taken from flag -g

- Analyze: Analyzes the log files and return an array of string as results.

- filterEntry: Filters the log entries according to flag values.

- parseLogEntry: Parses log entries from file for analization.

- displayResult: Format results for display.
