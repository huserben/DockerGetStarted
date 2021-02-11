package main

import (
    "fmt"
    "log"
	"os"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	return_value, present := os.LookupEnv("RETURN_VALUE")
	if present == true{		
		fmt.Fprintln(w, return_value)
	} else {
		fmt.Fprintln(w, "Hello from Go!")
	}
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}