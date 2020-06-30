package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
	"io/ioutil"
)

// Dumps bytes like hexdump -C with trimming
func dump(bytes []byte) string {
	if len(bytes) == 0 {
		return "                                                  ||"
	}

	const threshold = 16
	b := new(strings.Builder)

	hexb := []byte(bytes)
	if len(hexb) > threshold {
		hexb = hexb[:threshold]
	}
	dump := hex.Dump(hexb)
	dump = dump[10:len(dump) - 1]
	b.WriteString(dump)
	if len(bytes) > threshold {
		b.WriteString(" ...")
	} else {
		b.WriteString("")
	}
	return b.String()
}

func main() {
	fmt.Printf("args length: %d\n", len(os.Args))
	stat, _ := os.Stdin.Stat()
	stdinExists := (stat.Mode() & os.ModeCharDevice) == 0
	var stdin []byte
	if stdinExists {
		stdin, _ = ioutil.ReadAll(os.Stdin)
	}
	for i, arg := range os.Args {
		fmt.Printf("args[%d]: %4d bytes %4d runes \n", i, len(arg), utf8.RuneCountInString(arg))
		fmt.Println(dump([]byte(arg)))
	}
	fmt.Println()
	if stdinExists {
		fmt.Println("stdin: exists")
		fmt.Println(dump(stdin))
	} else {
		fmt.Println("stdin: not available")
	}
}
