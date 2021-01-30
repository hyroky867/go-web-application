package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	fmt.Sprintf("%sapp", otherWord),
	fmt.Sprintf("%ssite", otherWord),
	fmt.Sprintf("%stime", otherWord),
	fmt.Sprintf("get%s", otherWord),
	fmt.Sprintf("go%s", otherWord),
	fmt.Sprintf("lets%s", otherWord),
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
