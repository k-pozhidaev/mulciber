package main

import "os"

func main() {
	for _, arg := range os.Args {
		println(arg)
	}
}