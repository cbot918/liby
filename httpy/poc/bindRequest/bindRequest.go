package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

var lg = fmt.Println
var lf = fmt.Println

type MyRequest struct {
	Method  string
	URL     string
	Headers map[string][]string
	Body    string
}

func parseRequest(requestDump string) (*MyRequest, error) {
	var req MyRequest

	// Split the request dump into lines
	scanner := bufio.NewScanner(strings.NewReader(requestDump))
	for scanner.Scan() {
		line := scanner.Text()

		// Parse request method and URL
		if strings.HasPrefix(line, "GET ") {
			req.Method = "GET"
			req.URL = strings.TrimSpace(strings.TrimPrefix(line, "GET "))
		} else if strings.HasPrefix(line, "POST ") {
			req.Method = "POST"
			req.URL = strings.TrimSpace(strings.TrimPrefix(line, "POST "))
		}

		// Parse headers
		if strings.Contains(line, ": ") {
			parts := strings.SplitN(line, ": ", 2)
			key := parts[0]
			value := parts[1]
			if req.Headers == nil {
				req.Headers = make(map[string][]string)
			}
			req.Headers[key] = append(req.Headers[key], value)
		}

		// Parse request body (this is a simplified example)
		if line == "" {
			// The request body is assumed to be everything after the empty line
			var bodyLines []string
			for scanner.Scan() {
				bodyLines = append(bodyLines, scanner.Text())
			}
			req.Body = strings.Join(bodyLines, "\n")
		}
	}

	return &req, nil
}

type Httpy struct{}

func (h *Httpy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// lg(string(requestDump))

	req, err := parseRequest(string(requestDump))
	if err != nil {
		fmt.Println("Error parsing request:", err)
		return
	}

	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("URL: %s\n", req.URL)
	fmt.Printf("Headers: %+v\n", req.Headers)
	lg("--------")
	// fmt.Printf("Body:\n%s\n", req.Body)

	for key, val := range req.Headers {
		lg(key + ": " + val[0])
	}

}

const port = ":8080"

func main() {
	engine := &Httpy{}
	fmt.Println("listening " + port)
	if err := http.ListenAndServe(port, engine); err != nil {
		log.Fatal(err)
	}
}
