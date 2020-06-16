/*
ROT13 ("rotate by 13 places", sometimes hyphenated ROT-13) is a simple letter substitution
cipher that replaces a letter with the 13th letter after it, in the alphabet.
- Wikipedia
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		rotStr := rot(ln)
		fmt.Fprintf(conn, rotStr+"\n")
	}
}

func rot(str string) string {
	var result string

	for _, char := range str {
		x := char + 13
		if char >= 65 && char <= 90 {
			if x > 90 {
				x = (char + 13) - 90 + 64
			}
		} else {
			if x > 122 {
				x = (char + 13) - 122 + 96
			}
		}

		result += string(x)
	}

	return result
}
