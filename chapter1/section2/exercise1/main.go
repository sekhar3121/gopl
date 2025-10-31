package main

import (
	"fmt"
	"os"
)

func main() {
	result, sep := "", " "
	result += os.Args[0]
	for i := 0; i < len(os.Args[1:]); i++ {
		result += sep + os.Args[i]
	}
	fmt.Println(result)

}
