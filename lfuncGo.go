package main

import "fmt"

func main() {
	fmt.Println("AAA")
	dosomting()
	var xxx = dosomtingandReturn()

	var aaa, bbb = multiReturn()
	defer println(aaa, "  ", bbb)
	defer println(xxx) //defer is do when function finish all
}

func dosomting() {
	fmt.Println("BBB")
}

func dosomtingandReturn() string {
	fmt.Println("CCC")

	return "ASDG"
}

func multiReturn() (string, string) {
	fmt.Println("GGGG")

	return "HHHH", "AAAA"
}
