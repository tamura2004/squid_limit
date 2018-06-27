package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	host = flag.String("h", "sq", "hostname")
	port = flag.String("p", "3128", "port")
	num  = flag.Int("n", 10, "connection number to try keep")
)

func init() {
	flag.Parse()
}

func main() {
	conn, err := net.Dial("tcp", *host+":"+*port)
	handleErr(err)

	req := fmt.Sprintf("GET http://web%d/ HTTP/1.0\r\nHost")
	conn.Write([]byte(req))
}

func handleErr(err error) {
	log.Fatal(err)
}
