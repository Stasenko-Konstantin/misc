package main

import (
	"os"
)

func count(file string) {

}

func matchR(arg string)  (deep int, ok bool) {

}

func main() {
	args := os.Args
	if len(args) == 1 {

	} else if len(args) == 2 && args[1] == "-r" {
		
	} else if _, ok := matchR(args[2]); len(args) >= 2 && ok {

	} else {
		for _, f := range args[1:] {
			count(f)
		}
	}
}
