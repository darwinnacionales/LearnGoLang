package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	lyrics := "Ladies and gents, this is the moment you've waited for\nBeen searchin' in the dark, your sweat soakin' through the floor\nAnd buried in your bones there's an ache that you can't ignore"

	scanner := bufio.NewScanner(strings.NewReader(lyrics))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	fmt.Printf("\n-------------------------\n\n")

	scanner = bufio.NewScanner(strings.NewReader(lyrics))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
