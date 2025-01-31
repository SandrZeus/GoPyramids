package main

import "fmt"

func main() {

	var rows int = 5
	var k, i, j int

	fmt.Println("1st pyramid - Full Pyramid")

	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}

	fmt.Println("")

	fmt.Println("2nd pyramid - Inverted Full Pyramid")

	for i = rows; i >= 1; i-- {
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for j = i; j <= 2*i-1; j++ {
			fmt.Printf("* ")
		}
		for j = 0; j < i-1; j++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}

	fmt.Println("")

	fmt.Println("3rd pyramid - Right Half Pyramid")

    for i := 1; i <= rows; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }

}