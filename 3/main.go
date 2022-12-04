package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	type rucksack struct {
		first  string
		second string
	}
	rucksacks := []rucksack{}

	file, err := os.ReadFile("rucksacks.txt")
	if err != nil {
		log.Fatal(err)
	}

	rs := strings.Split(string(file), "\n")

	for _, v := range rs {
		f := strings.TrimSpace(v[:len(v)/2])
		s := strings.TrimSpace(v[len(v)/2:])
		r := rucksack{
			first:  f,
			second: s,
		}
		rucksacks = append(rucksacks, r)
	}

	var c int = 1
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var total int

	for _, v := range rucksacks {
		for _, char := range v.first {
			if strings.Contains(v.second, string(char)) {
				fmt.Printf("Rucksack %v: item in both compartments -> %v\n", c, string(char))
				c++
				if unicode.IsLower(char) {
					total += strings.Index(strings.ToLower(alpha), string(char)) + 1
				} else {
					total += strings.Index(alpha, string(char)) + 27
				}
				break
			}
		}
	}
	fmt.Printf("Sum of priorities = %v", total)
}
