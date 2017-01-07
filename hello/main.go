package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("Hello, world\n")
		fmt.Println("The current time is: ", time.Now())
	}
}
