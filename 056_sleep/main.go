package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sleepMS, _ := strconv.Atoi(r.FormValue("sleep"))
		fmt.Fprintf(w, "sleep for [%d]", sleepMS)
		time.Sleep(time.Duration(sleepMS) * time.Millisecond)
		fmt.Fprint(w, "done ")
	})

	http.ListenAndServe(":8080", nil)
}
