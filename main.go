// This file is a command-line tool that reads a text file from input and
// encodes it with the given radix (e.g, base-16, base-32, etc.)
// This program is for learning purposes, so the stdlib `hex.Encode`` was not used.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const base16Map = "0123456789abcdef"

func main() {
	n := flag.Int("base", 16, "radix size to base-N encode. defaults to 16")
	f := flag.String("f", "", "file to base-N encode")
	flag.Parse()

	if *n != 16 {
		panic(fmt.Sprintf("base-%d not handled yet", *n))
	}

	bytes, err := os.ReadFile(*f)
	if err != nil {
		log.Fatalf(err.Error())
	}

	enc := base16Encode(bytes)

	fmt.Println("Content")
	fmt.Println(string(bytes))
	fmt.Println()
	fmt.Println("Encoded (base16)")
	fmt.Printf("%s\n", enc)
}

// stdlib hex.Encode does this without string <> byte conversions ...
func base16Encode(bytes []byte) string {
	var bits string
	for _, b := range bytes {
		bits += fmt.Sprintf("%08b", b)
	}
	var enc string
	for i := 0; i < len(bits); i += 4 {
		baseChar := bits[i:min(i+4, len(bits))]
		d, _ := strconv.ParseInt(baseChar, 2, 64)
		enc += string(base16Map[int(d)])
	}
	return enc
}
