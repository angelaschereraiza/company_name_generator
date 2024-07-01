package main

import (
	"fmt"
	"net"
)

func main() {
	vocals := []rune("aeijouy")
	consonants := []rune("bcdfghklmnpqrstvwxz")

	// Create company name
	for i := 0; i < len(consonants); i++ {
		name := "A"
		name += string(consonants[i])

		for j := 0; j < len(vocals); j++ {
			name2 := name + string(vocals[j])
			for k := 0; k < len(consonants); k++ {
				name3 := name2 + string(consonants[k]) + "a"

				// Check if domain is used
				_, err := net.LookupCNAME(name3 + ".ch")

				if err != nil {
					fmt.Println("Name: ", name3)
					fmt.Println("-----")
				}
			}
		}
	}
}
