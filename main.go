package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	lineBreak  = "----------------------------------------"
	portFlag   = flag.Int("port", 5555, "Port number")
	headerFlag = flag.Bool("header", true, "Print headers")
	bodyFlag   = flag.Bool("body", false, "Print body")
)

var requestParrotsChan = make(chan []string)

func main() {
	flag.Parse()

	go printRequestParrots(requestParrotsChan)
	addr := fmt.Sprintf(":%d", *portFlag)
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(requestPrinter)))
}

func requestPrinter(w http.ResponseWriter, r *http.Request) {
	requestParrots := make([]string, 0)

	requestParrots = append(requestParrots, fmt.Sprintf("%s", r.URL))

	if *headerFlag {
		for k, v := range r.Header {
			values := strings.Join(v, ", ")
			requestParrots = append(requestParrots, fmt.Sprintf("[%s] = '%s'", k, values))
		}
	}

	if *bodyFlag {
		var reader io.ReadCloser
		var err error
		var ignoreBody = false

		switch r.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(r.Body)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
				ignoreBody = true
				break
			}
		default:
			reader = r.Body
		}

		if !ignoreBody {
			b, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
			}

			requestParrots = append(requestParrots, fmt.Sprintf("[Body] = '%s'\n", b))
		}

		defer r.Body.Close()
	}

	requestParrotsChan <- requestParrots
}

func printRequestParrots(requestParrotsChan chan []string) {
	for requestParrot := range requestParrotsChan {
		for i, line := range requestParrot {
			fmt.Printf("%d: %s\n", i, line)
		}
		fmt.Println(lineBreak)
	}
}
