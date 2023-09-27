package main

import "fmt"

func main() {
	var fruitList [6]string

	fruitList[0] = "mango"
	fruitList[1] = "blackberry"
	fruitList[2] = "orange"
	fruitList[3] = "lemon"
	fruitList[5] = "seanese"

	fmt.Println("the array of fruitList :", fruitList)
	fmt.Println("the length of fruitList :", len(fruitList))
	var countryList = [4]string{"india", "canada", "usa"}

	fmt.Println("the array of countryList :", countryList)
	fmt.Println("the length of countryList :", len(countryList))

}
