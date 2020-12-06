package main

import "fmt"

func main() {
	a := true
	if a || 1 > 0 {
		fmt.Println("ok")
	}

}
