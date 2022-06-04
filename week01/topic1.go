package main

import "fmt"

func main() {
	var arr = [5]string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < 5; i++ {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		}
		if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}
	fmt.Println(arr)
}
