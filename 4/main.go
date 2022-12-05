package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("pairs.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	file.Close()

	var contains []string
	re := regexp.MustCompile("[0-9]+")

	for i, v := range lines {
		r := re.FindAllString(v, -1)
		p1 := r[:len(r)/2]
		p2 := r[len(r)/2:]
		switch {
		case p1[0] >= p2[0] && p1[1] <= p2[1]:
			contains = append(contains, fmt.Sprintf("Pair %v -> %v fully contains %v\n", i, p2, p1))
		case p2[0] >= p1[0] && p2[1] <= p1[1]:
			contains = append(contains, fmt.Sprintf("Pair %v -> %v fully contains %v\n", i, p1, p2))
		}
	}

	fmt.Printf("There are %v pairs that have a range fully containing the other:\n", len(contains))
	for _, v := range contains {
		fmt.Print(v)
	}
}
