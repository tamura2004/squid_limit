package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	host = flag.String("h", "sq", "hostname")
	port = flag.String("p", "3128", "port")
	num  = flag.Int("n", 10, "connection number to try keep")
)

var wg sync.WaitGroup

func init() {
	flag.Parse()
	countDown()
}

func countDown() {
	count := 10
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			log.Println(count)
			if count == 0 {
				t.Stop()
				return
			}
			count--
		}
	}
}

func main() {
	for i := 0; i < *num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", *host+":"+*port)
			handleErr(err)

			req := fmt.Sprintf("GET http://web%d/ HTTP/1.0\r\nHost")
			conn.Write([]byte(req))
		}()
	}
	wg.Wait()
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
