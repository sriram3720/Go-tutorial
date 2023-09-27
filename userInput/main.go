package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welocme := "welcome to userInput"
	fmt.Println(welocme)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating for pizza: ")

	// comma ok , error ok
	rating, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating :", rating)
	fmt.Printf("Type of this input  : %T", rating)
}
