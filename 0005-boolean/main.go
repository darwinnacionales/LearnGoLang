package main

import (
	"fmt"
)

func main() {
	s1 := "word1"
	s2 := "word1"
	s3 := "word2"
	result1 := s1 == s2
	result2 := s1 == s3
	sampleVar := true
	var explicitVar bool = false

	fmt.Println(result1, result2, sampleVar)
	fmt.Printf("%T\n", sampleVar)
	fmt.Println(explicitVar)
}
