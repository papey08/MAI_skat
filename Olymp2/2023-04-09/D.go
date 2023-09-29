package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int
	var rank string
	fmt.Scan(&age, &rank)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = text[:len(text)-1]
	switch rank {
	case "G":
		fmt.Println("OK")
	case "PG":
		if age < 13 {
			fmt.Println("OK IF ACCOMPANIED")
		} else {
			fmt.Println("OK")
		}
	case "R-13":
		if age < 13 {
			fmt.Println("ACCESS DENIED")
		} else {
			fmt.Println("OK")
		}
	case "R-16":
		if age < 16 {
			fmt.Println("ACCESS DENIED")
		} else {
			fmt.Println("OK")
		}
	case "R-18":
		if age < 18 {
			fmt.Println("ACCESS DENIED")
		} else {
			fmt.Println("OK")
		}
	}
}
