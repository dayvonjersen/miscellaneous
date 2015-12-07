// column
//
// because gnu column doesn't exist on windows
//
//
// install
//
//	$ go install github.com/generaltso/miscellaneous/column
//
// usage
//
// in
//
//	$ echo -e "a b c\naaaaaaaaaa b c\na bbbbbbbbbb c" | column
//
// out
//
//	a          b          c
//	aaaaaaaaaa b          c
//	a          bbbbbbbbbb c
//
// flags
//
//	-l="\n": Separator to use when parsing input rows
//	-n="\n": Separator to use when formatting output rows
//	-q=" ": Separator to use when formatting output
//	-s=" ": Separator to use when parsing input columns
package main
