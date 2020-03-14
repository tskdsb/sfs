package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var (
		addr string
		dir  string
	)

	flag.StringVar(&addr, "addr", "0.0.0.0:4500", "ip:port")
	flag.StringVar(&dir, "dir", ".", "directory to export")
	flag.Parse()

	log.Printf("Listen on: %s\n", addr)
	log.Printf("Export on: %s\n", dir)
	panic(http.ListenAndServe(addr, http.FileServer(http.Dir(dir))))
}
