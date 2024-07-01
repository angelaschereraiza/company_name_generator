package main

import (
	"fmt"
	"net"
)

func main() {
	vocals := []rune("aeijouy")
	consonants := []rune("bcdfghklmnpqrstvwxz")

	// Create company name
	for i := 0; i < len(vocals); i++ {
		name := "A"
		name += string(vocals[i])

		for j := 0; j < len(consonants); j++ {
			name2 := name + string(consonants[j]) + "a"

			// Check if domain is used
			_, err := net.LookupCNAME(name2 + ".ch")

			if err != nil {
				fmt.Println("Name: ", name2)
				fmt.Println("-----")
			}
		}
	}
}
