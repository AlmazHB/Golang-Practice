package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [3]string{"a", "b", "c"}
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
	fmt.Println(len(numbers))

	for index, volume := range letters {
		fmt.Println(index, " ", volume)
	}
	massiw := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	simpleCount := float64(len(massiw))
	for _, number := range massiw {
		sum += number
	}
	fmt.Printf("Orta Arifmrtik baha:%0.2f\n", sum/simpleCount)
}
