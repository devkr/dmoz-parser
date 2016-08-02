package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Input file
	fi, err := os.Open(`C:\temp\content.rdf.u8`)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	// Output file
	fo, err := os.Create("links.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	w := bufio.NewWriter(fo)

	scanner := bufio.NewScanner(fi)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "<ExternalPage about=") {
			continue
		}

		strings.TrimSpace(line)

		lineSize := len(line)
		if lineSize < 26 {
			continue
		}

		link := line[23 : lineSize-2]

		fmt.Fprintln(w, link)
		if i == 10000 {
			fmt.Print(".")
			w.Flush()
			i = 0
		}
		i++
	}
	w.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
