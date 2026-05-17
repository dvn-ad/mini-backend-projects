package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2{
		println("Error: Missing username argument.")
		println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	events, err := FetchGithubActivity(username); 
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Error: %v\n",err)
		os.Exit(1)
	}
	FormatAndPrintActivity(events)
}