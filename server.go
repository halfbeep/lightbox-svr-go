package main

import (
	"os"
	"flag"
	"fmt"
	"log"

	"github.com/vasyahuyasa/july/opds"
)

func main() {
	rootDir := flag.String("d", "", "Root storage directory")
	port := flag.String("p", os.Getenv("PORT"), "Service Port")
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
