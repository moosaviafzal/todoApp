package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Todo App , Start with type: -salam=")

	salam := flag.String("salam", "command", "run do you want")
	flag.Parse()

	if *salam == "Ejad-kar" {
		var name, daste string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter name bache")
		scanner.Scan()
		name = scanner.Text()

		fmt.Println("Enter daste bache")
		scanner.Scan()
		daste = scanner.Text()
		fmt.Println(name, daste)

	}

}
