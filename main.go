package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v] Got %v %v%v", time.Now().UTC(), r.Method, r.URL.Path, r.URL.RawQuery)
	headers := ""
	for h, v := range r.Header {
		headers += h + ": " + strings.Join(v, ", ") + "</br>\n"
	}
	// v, _ := ioutil.ReadAll(r.Body)
	v, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, `
<html>
<head>
	<title>Silver Surfing</title>
</head>
<body>
	Method: %v </br>
	URL: %v </br>
	QueryString: %v </br>
	Request Body: %v </br>
	-------------------------- </br>
	<h2>Headers</h2>
	</br>
	%v
</body>
</html>
`,
		r.Method,
		r.URL.Path,
		r.URL.RawQuery,
		string(v),
		headers)
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP Requests...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
