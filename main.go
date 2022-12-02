package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("food.txt")
	if err != nil {
		log.Fatal(err)
	}
	var s int
	type fs2 struct {
		v int
		i int
	}
	fs := []fs2{}
	i := strings.Split(string(f), "\r\n")
	c := 0
	for j := 0; j < len(i); j++ {
		n, _ := strconv.Atoi(i[j])
		s += n
		if n == 0 {
			fs = append(fs, fs2{s, c})
			s = 0
			c++
		}
	}
	sort.Slice(fs, func(i, j int) bool {
		return fs[i].v > fs[j].v
	})
	fmt.Printf("✅ Elf %v is the best. They have %v calories", fs[0].i, fs[0].v)
}
