package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
)

func main() {
	nums, err := getFloat64("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}
	sampleCount := float64(len(nums))
	fmt.Printf("OrtaArifmetikBaha:%0.2f\n", sum/sampleCount)
}
func getFloat64(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
