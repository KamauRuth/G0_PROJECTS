package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")

	//comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("The type of input is :%T", input)

}
