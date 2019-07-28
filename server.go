package main

import (
	"os"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/vasyahuyasa/july/opds"
)

func getPort() int {
	s := os.Getenv("PORT")
	if s != "" {
		p, err := strconv.Atoi(s)
			if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
			}
		return p 
	}
	return 8080
}

func main() {
	rootDir := flag.String("d", "", "Root storage directory")
	port := flag.Int("p", getPort(), "Service port")
	host := flag.String("i", "0.0.0.0", "Service network interface")
	flag.Parse()

	srv := &opds.Server{
		FileRoot: *rootDir,
	}
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Run server: ", "http://"+addr+"/opds")
	err := srv.Run(addr)
	log.Fatal(err)
}
