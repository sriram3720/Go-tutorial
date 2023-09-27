package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["py"] = "python"
	languages["js"] = "javascript"
	languages["java"] = "java"
	fmt.Println("The following languages :", languages)
	fmt.Println("the language of Py :", languages["py"])

	delete(languages, "js")
	fmt.Println("The following languages :", languages)

	for key, value := range languages {
		fmt.Printf("for key: %v,value :%v \n", key, value)
	}
}
