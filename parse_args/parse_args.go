package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// for _, arg := range os.Args {
	// 	// if idx == 0 {
	// 	// 	continue
	// 	// }
	// 	s += sep + arg
	// 	sep = " "

	// }

	fmt.Println(s)
}
