package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var args []string 							//manual type declaration   1/2
    args = os.Args								// -||- 				    2/2
	hourOfDay := time.Now().Hour()				//type inference
	greeting := getGreeting(hourOfDay)

	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println(greeting)
	}
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning Sweetheart"
	} else if hour < 18 {
		return "Splendid Afternoon You Beautiful Being"
	} else {
		return "Dark Evening Your Vampiresy"
	}
}
