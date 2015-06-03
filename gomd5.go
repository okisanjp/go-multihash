package main

import (
	"crypto/md5"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	h := []byte(flag.Arg(0))
	hash(h)
}

func hash(h []byte) {
	fmt.Printf("%x\n", md5.Sum(h))
}
