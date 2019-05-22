package main



type Person struct {
	Name  string
	Height int
}


type Dog struct {
	Person
	Name string
}


type Animal struct {
	Name string
	Age int
	Action struct{
		Height int
	}
}


func main() {
	dog := Dog{Name: "xx"}
	dog.Height = 1
	dog.Person.Height = 6


	an := Animal{
		Name: "xx",
		Age: 1,
	}
	an.Action.Height = 1
}