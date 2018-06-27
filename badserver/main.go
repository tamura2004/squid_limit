package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"log"
)

var port = flag.String("p", "80", "listening port")

func init() {
	flag.Parse()
	log.Println("start")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3600 * time.Second)
		fmt.Fprint(w, "Hello")
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
