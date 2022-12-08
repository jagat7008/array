package main

import "fmt"

func main() {
	var a int = 7
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := []int{1, 2, 3}
	fmt.Println(b)
	fmt.Println(b[2])
	b = append(b, 4, 5, 6)
	fmt.Println(b)
	d := map[string]int{
		"vicky": 1,
		"king":  2,
	}
	fmt.Println(d)
}
