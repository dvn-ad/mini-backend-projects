package main

import (
	"flag"
	"fmt"
)

func main(){
	port:=flag.Int("port",8080,"Port server to run on")
	origin:=flag.String("origin","","Origin URL to proxy requests to")

	flag.Parse()

	if *origin ==""{
		fmt.Println("Error: --origin flag is required")
		flag.Usage()
		return
	}

	fmt.Printf("Starting caching proxy on port :%d\n", *port)
	fmt.Printf("Proxying traffic to origin: %s\n", *origin)
}