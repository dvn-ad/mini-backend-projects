package main

import (
	"caching-proxy/internal/cache"
	"caching-proxy/internal/proxy"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	port:=flag.Int("port",8080,"Port server to run on")
	origin:=flag.String("origin","","Origin URL to proxy requests to")
	clearCache:=flag.Bool("clear-cache",false,"CLear the stored cache and exit")	

	flag.Parse()

	c:=cache.NewCacheSystem()

	if *clearCache{
		c.Clear()
		fmt.Println("Cache cleared successfully")
		os.Exit(0)
	}

	if *origin ==""{
		fmt.Println("Error: --origin flag is required")
		flag.Usage()
		os.Exit(1)
	}

	p:=&proxy.ProxyServer{
		Cache: c,
		OriginURL: *origin,
	}
	http.HandleFunc("/",p.ProxyHandler)

	addr:=fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting caching proxy on port :%d\n", *port)
	fmt.Printf("Proxying traffic to origin: %s\n", *origin)

	err:=http.ListenAndServe(addr,nil)
	if err!=nil{
		log.Fatalf("Server failed to start: %v", err)
	}


	
}