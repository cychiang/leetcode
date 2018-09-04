package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	var c = 0
	for _, s := range S {
		if strings.Contains(J, string(s)) {
			c += 1
		}
	}
	return c
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
