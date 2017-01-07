package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	var args []string              		//manual type declaration   1/2
	args = os.Args                		// -||- 				    2/2
	hourOfDay := time.Now().Hour() 		//type inference
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(args) > 1 {
		if len(args[1]) < 3 {
			for i := 0; i < 3; i++ {
				fmt.Println(args[1])
			}
		} else {
			i := 0
			isLessThanSix := true

			for isLessThanSix {
				if i >= 6 {
					isLessThanSix = false
				} else if i == 1 {
					fmt.Println("Fear my spamming technique!!!")
				}
				fmt.Println(args[1])
				i++
			}
		}
	} else {
		fmt.Println(greeting)
	}
}

func getGreeting(hour int) (string, error) {
	var greeting [4]string
	greeting[0] = "Khhhhh!! Go back to sleep! Programs also need some rest"
	greeting[1] = "Good Morning Sweetheart"
	greeting[2] = "Splendid Afternoon You Beautiful Being"
	greeting[3] = "Dark Evening Your Vampiresy"
	var message string

	if hour < 4 {
		err := errors.New(greeting[0])
		return message, err
	} else if hour < 7 {
		message = greeting[1]
	} else if hour < 16 {
		message = greeting[2]
	} else {
		message = greeting[3]
	}
	return message, nil
}
