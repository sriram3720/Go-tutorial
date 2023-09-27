package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "orange", "blackberry"}

	fmt.Printf("Type of fruit list : %T\n", fruitList)

	fruitList = append(fruitList, "lemon", "pineapple")
	fmt.Println("The value of added new fruit list:", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("The value of new fruit list:", fruitList)

	highRate := make([]string, 4)

	highRate[0] = "apple"
	highRate[1] = "orange"
	highRate[2] = "lemon"
	highRate[3] = "pineapple"

	fmt.Println("The value of highRate:", highRate)

	highRate = append(highRate, "blackberry")
	fmt.Println("The value of highRate:", highRate)
	sort.Strings(highRate)
	fmt.Println(sort.StringsAreSorted(highRate))

	var country = []string{"india", "germany", "japan", "china"}
	var index int = 2
	country = append(country[:index], country[index+1:]...)
	fmt.Println(country)

}
