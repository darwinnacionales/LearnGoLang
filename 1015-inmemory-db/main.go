package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Here are the commands:")
	fmt.Println("SET key val")
	fmt.Println("GET key")
	fmt.Println("DEL key")

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
	m := map[string]string{}

	for scanner.Scan() {
		command := scanner.Text()
		splitted := strings.Split(command, " ")

		if splitted[0] == "GET" {
			key := splitted[1]
			val, isExisting := m[key]

			if isExisting {
				fmt.Fprintf(conn, val+"\n")
			} else {
				fmt.Fprintf(conn, key+" is not existing\n")
			}
		} else if splitted[0] == "SET" {
			key := splitted[1]
			val := splitted[2]

			m[key] = val
		} else if splitted[0] == "DEL" {
			key := splitted[1]
			delete(m, key)
		}
	}
}
