package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := 10
	max := 1000
	ints := []int{}

	for i := 0; i <= num; i++ {
		ints = append(ints, r.Intn(max))
	}

	for i, integer := range ints {
		if integer%2 == 0 {
			fmt.Println(i, (strconv.Itoa(integer) + " is Even"))
		} else {
			fmt.Println(i, (strconv.Itoa(integer) + " is Odd"))
		}
	}
}
