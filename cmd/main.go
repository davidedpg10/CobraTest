package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)



func main() {
	promt := promptui.Select{
		Label: "Select the person",
		Items: []string{"John", "Wick", "Lord", "Voldemort"},
	}
	i, res, err := promt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Println("Hello world")
	fmt.Printf("You choose %q\n", res)
	fmt.Printf("Number returned %v\n", i)
}
