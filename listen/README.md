# listen
listen on a port and print all incoming data to stdout

## why
because I've had to write this more than once

you know how you use telnet to figure out how things like ftp and smtp work?

sometimes you need to analyse the request a client is sending

actually if telnet can do this already then [let me know](mailto:tso@teknik.io)

## install
`$ go install github.com/generaltso/miscellaneous/listen`

## usage
### in
1. `$ listen -p 80`
2. Open [http://localhost](http://localhost)

### out
```
GET / HTTP/1.1
Host: localhost
User-Agent: Mozilla/Something
Accept: */*
Accept-Language: en
Accept-Encoding: gzip, deflate
Connection: keep-alive
```
etc

### flags
```
-p=80: port to listen on
```
