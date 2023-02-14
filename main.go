package main

import (
	"flag"
	"fmt"
	"httpredirecter/redirection"
	"httpredirecter/webserver"
	"os"
)

func main() {
	filePath := flag.String("routes", "default", "Route file name")
	port := flag.Uint("port", 7878, "Webserver port")
	flag.Parse()

	fmt.Printf("Looking for redirections in routes/%s\n", *filePath)
	redirections, err := redirection.GetRedirections(*filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Starting server on 0.0.0.0:%d\n", *port)
	err = webserver.CreateWebserver(uint16(*port), redirections)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
