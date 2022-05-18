package main

import "fmt"

type Books struct {
	ID    string
	Name  string
	Price string
}

func main() {

	name := []string{} //list

	name = append(name, "Chaiyarin")
	name = append(name, "Atikom")
	name = append(name, "Kritsana")

	fmt.Println(name)

	json := make(map[string]int) // make key and value

	json["AAA"] = 1
	json["BBB"] = 2
	json["CCC"] = 3

	fmt.Println(json)

	// list + struc

	jsonBooks := []Books{}

	jsonBooks = append(jsonBooks, Books{
		ID:    "1",
		Name:  "TEST",
		Price: "22",
	})

	jsonBooks = append(jsonBooks, Books{
		ID:    "2",
		Name:  "TEST3",
		Price: "22",
	})

	fmt.Println(jsonBooks)

}
