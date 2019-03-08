package main

/**
 * Show all lines of a file that contain the filter.
 * Example: go run src/grep.go --filename LICENSE --filter a
 */

import "log"
import "fmt"
import "strings"
import "flag"
import "io/ioutil"

func main() {
	filename := flag.String("filename", "", "filename to grep")
	filter := flag.String("filter", "", "filter for lines")
	flag.Parse()
	//log.Printf("Opening %v filter %v", *filename, *filter)
	readAndPrint(filename, filter)
}

func readAndPrint(filename *string, filter *string) {
	contents, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if strings.Contains(line, *filter) {
			fmt.Printf("%v", line+"\n")
		}
	}
}
