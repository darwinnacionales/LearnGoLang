package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `pass123word`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte("pass123word"))

	if err != nil {
		fmt.Println("Incorrect password")
	} else {
		fmt.Println("Correct Password")
	}
}
