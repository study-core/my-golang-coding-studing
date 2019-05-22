package main

import "fmt"

func main() {
	map1 := make(map[string]*MyAnimal, 0)
	map2 := make(map[string]*MyAnimal, 0)
	map1["a"] = &MyAnimal{
		Name: "A",
	}
	fmt.Println("map1", map1)
	fmt.Println("map2", map2)
	map2["a"] = map1["a"]
	fmt.Println("map1", map1)
	fmt.Println("map2", map2)
	delete(map1, "a")
	fmt.Println("map1", map1)
	fmt.Println("map2", map2)

}

type MyAnimal struct {
	Name string
}