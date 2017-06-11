package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var port int

func main() {
	flag.IntVar(&port, "p", 80, "port to listen on")
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	checkErr(err)
	for {
		conn, err := ln.Accept()
		checkErr(err)
		b := make([]byte, 512)
		for {
			_, err := conn.Read(b)
			fmt.Print(string(b))
			if err == io.EOF {
				break
			}
			checkErr(err)
		}
	}
}
