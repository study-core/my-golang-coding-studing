package main 

import (
 	"fmt"
)


func main() {

	data := map[string]interface{}{}

	data["id"] = uint64(14)

	data["age"]  = int32(52)

	data["price"] = 363.7

	data["sale"] = 782.56

	data["mValue"] = 45.7


	data["fAge"] = -45.7 

	value := "45.7"

	id := fmt.Sprint(data["id"])

	age := fmt.Sprint(data["age"])
	
	price := fmt.Sprint(data["price"])

	sale := fmt.Sprint(data["sale"])

	mValue := fmt.Sprint(data["mValue"])

	fAge := fmt.Sprint(data["fAge"])




	fmt.Println(id < value) // true

	fmt.Println(age < value) //false

	fmt.Println(price < value) // false


	fmt.Println(sale < value) // false

	fmt.Println(mValue <= value && value <= sale) //true

	fmt.Println(fAge < value) // true
}