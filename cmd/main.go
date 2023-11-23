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
	_, res, err := promt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose %q\n", res)
}
