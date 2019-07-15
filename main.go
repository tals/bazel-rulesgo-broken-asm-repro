package main

import (
	"github.com/golang/snappy"
)

func main() {
	hello_snap := []byte{13, 48, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	hello, _ := snappy.Decode(nil, hello_snap)
	println(string(hello))
}
