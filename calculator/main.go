package main

import "fmt"

func main() {

	structs()
}

type Data struct {
	X float64
	S string
	Y float64
}

func structs() {
	data1 := Data{
		X: 12,
		S: "*",
		Y: 6,
	}
	fmt.Println(data1)
}
