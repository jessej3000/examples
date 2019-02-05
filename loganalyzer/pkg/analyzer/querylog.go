package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// QueryLog definition
type QueryLog struct {
	LogFile       *os.File
	ReportType    string
	NumberOfEntry int
	DisplayList   bool
	Status        bool
	GetMethod     bool
	NoGIF         bool
}

type pair struct {
	Key   string
	Value float64
}

// New create a new instance of QueryLog
func New(filepath string,
	reportType string,
	numberOfEntry int,
	status bool,
	getMethod bool,
	noGif bool) *QueryLog {
	var err error
	var tmpFile *os.File

	tmpFile, err = os.Open(filepath)
	if err != nil {
		tmpFile.Close()
		panic(err)
	}

	q := &QueryLog{
		// Initialize property
		ReportType:    reportType,
		NumberOfEntry: numberOfEntry,
		LogFile:       tmpFile,
		Status:        status,
		GetMethod:     getMethod,
		NoGIF:         noGif,
	}

	return q
}

// Analyze log file
func (q *QueryLog) Analyze() []string {
	var urlsMap map[string]float64
	var urlsCounter map[string]int
	var pairs []pair

	urlsMap = make(map[string]float64)
	urlsCounter = make(map[string]int)

	scanner := bufio.NewScanner(q.LogFile)

	// Get log entries
	for scanner.Scan() {
		url, dur, stats := q.parseLogEntry(scanner.Text())

		if q.filterEntry(url, dur, stats) {
			_, ok := urlsMap[url]
			if ok {
				urlsMap[url] = urlsMap[url] + dur
				urlsCounter[url]++
			} else {
				urlsMap[url] = dur
				urlsCounter[url]++
			}
		}
	}

	// Get averages and assign to array of pair
	for k, v := range urlsMap {
		urlsMap[k] = v / float64(urlsCounter[k])
		pairs = append(pairs, pair{
			Key:   k,
			Value: urlsMap[k],
		})
	}

	// Sort
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	// Display results
	return (q.displayResult(pairs, urlsCounter))
}

// Filters log entry
func (q *QueryLog) filterEntry(url string, dur float64, stats int) bool {
	// Disallow GIFs
	if q.NoGIF {
		if strings.HasSuffix(url, ".gif") {
			return false
		}
	}

	// Check for status code 200
	if q.Status {
		if stats != 200 {
			return false
		}
	}

	// Checking for GET method is tricky as in the given: url, duration, status code
	// We can have a level of confidence in checking for GET method using the status code
	// as the url can be any time of of request, GET or POST.
	// And the usual status code that most likely not a request made with GET is 201
	// So we do not count any log entry with status code
	if q.GetMethod {
		if stats == 201 {
			return false
		}
	}

	return true
}

func (q *QueryLog) displayResult(pairs []pair, counter map[string]int) []string {

	var list []string
	switch strings.ToLower(q.ReportType) {
	case "minimal":
		for i := 0; (i < q.NumberOfEntry) && (i < len(pairs)); i++ {
			list = append(list, pairs[i].Key)
		}
	case "complete":
		for i := 0; (i < q.NumberOfEntry) && (i < len(pairs)); i++ {
			tmp := fmt.Sprintf("%.2f", pairs[i].Value)
			list = append(list, pairs[i].Key+" : "+tmp+"s"+" : "+strconv.Itoa(counter[pairs[i].Key])+" entries")
		}
	default:
		fmt.Println("Invalid report type.")
	}
	return list
}

// Get values from log entry
func (q *QueryLog) parseLogEntry(entry string) (string, float64, int) {
	if len(entry) == 0 {
		return "", 0.0, 0
	}

	// Extract log details
	entries := strings.Split(entry, ",")
	dur := strings.Trim(entries[1], "s ")
	finalDur, err := strconv.ParseFloat(dur, 64)
	if err != nil {
		panic(err)
	}

	// Extrac status
	tmpStatus := strings.Split(entries[2], ":")
	status := strings.Trim(tmpStatus[1], " ")
	finalStatus, err := strconv.Atoi(status)
	if err != nil {
		panic(err)
	}

	return strings.ToLower(strings.Trim(entries[0], " ")), finalDur, finalStatus
}
