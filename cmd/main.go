package main

import "time"

func main() {
	startTime := time.Now()
	server, err := InitHttpServer()
	if err != nil {
		panic(err)
	}
	server.Run(startTime)
}
