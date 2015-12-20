// listen
//
// because I've had to write this more than once
//
// what
//
// listens on a port and prints any data that comes in
//
// why
//
// you know how you use telnet to figure out how things like ftp and smtp work
//
// sometimes you need to analyse the request a client is sending
//
// actually if telnet can do this already then let me know: tso@teknik.io
//
// install
//
//		$ go install github.com/generaltso/miscellaneous/listen
//
// usage
//
// in
//
//		$ listen -p 80
//
//		Open http://localhost
//
// out
//
//		GET / HTTP/1.1
//		Host: localhost
//		User-Agent: Mozilla/Something
//		Accept: */*
//		Accept-Language: en
//		Accept-Encoding: gzip, deflate
//		Connection: keep-alive
//
//		etc...
//
// flags
//
//		-p=80: port to listen on
//
package main
