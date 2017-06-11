# hexip
convert to/from hexadecimal IP addresses

## usage

```
$ hexip google.com
google.com 172.217.11.14 0xacd90b0e
```

or

```
$ hexip 172.217.11.14
lga25s60-in-f14.1e100.net 172.217.11.14 0xacd90b0e
```

or

```
$ hexip 0xacd90b0e
lga25s60-in-f14.1e100.net 172.217.11.14 0xacd90b0e
```

## install

```
$ go install github.com/generaltso/miscellaneous/hexip
```

## note

hostname lookup can fail, dig is a more reliable tool for reverse DNS lookups
