package main

import (
	"fmt"
	"time"
	"crypto/md5"
	"strings"
	"flag"
)

func calculate(nilNumbers int) {
	secret := "yzbqklnj"
	var result string
	index := 0
	resultIndex := 0
	startHash := strings.Repeat("0", nilNumbers)

	for {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", secret, index))))
		if hash[0:nilNumbers] == startHash {
			result = hash
			resultIndex = index
			break
		}
		index++
	}

	fmt.Println("Result = ", result, resultIndex)
}


var n int
func initFlags() {
	flag.IntVar(&n, "n", 5, "")
	flag.Parse()
}

func main() {
	start := time.Now()
	initFlags()
	calculate(n)
	end := time.Now()
	fmt.Println(end.Sub(start))
}