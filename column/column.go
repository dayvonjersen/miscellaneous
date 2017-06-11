package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var separator string
var newline string
var out_sep string
var out_line string

func checkErr(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func main() {
	flag.StringVar(&separator, "s", " ", "Separator to use when parsing input columns")
	flag.StringVar(&newline, "l", "\n", "Separator to use when parsing input rows")
	flag.StringVar(&out_sep, "q", " ", "Separator to use when formatting output")
	flag.StringVar(&out_line, "n", "\n", "Separator to use when formatting output rows")

	flag.Parse()

	stat, err := os.Stdin.Stat()
	checkErr(err)

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if stat.Size() <= 0 {
		println("nothing to read")
		return
	}

	stdin, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)
	rows := strings.Split(string(stdin), newline)

	out := make(map[int][]string)
	widths := make(map[int]int)

	y := len(rows)

	for i := 0; i < y; i++ {
		row := rows[i]
		if len(row) == 0 {
			continue
		}
		cols := strings.Split(row, separator)
		for j, col := range cols {
			l := len(col)
			if l == 0 {
				continue
			}
			max_len, exists := widths[j]
			if !exists || l > max_len {
				widths[j] = l
			}
		}
		out[i] = cols
	}

	for i := 0; i < y; i++ {
		ln := out[i]
		for idx, c := range ln {
			width := widths[idx] - len(c)
			fmt.Printf("%s%s", c, strings.Repeat(out_sep, width+1))
		}
		fmt.Print(out_line)
	}
}
