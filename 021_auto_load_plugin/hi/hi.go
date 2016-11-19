package main

// // Nothing to do
import "C"
import (
	"fmt"
	"net/http"
)

//MyAppHandler represents the entry point of http app in this pluging
func MyAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// Say function will print Hellow world.
func Say() {
	fmt.Println("Hello World")
}

func main() {

}
