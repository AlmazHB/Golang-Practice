package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printCircleArea(5)
	ppg(2.3, 5.08)
	pag(18, 32)
}

// PrintCircle Area
func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Circle radius is %d\n", radius)
	fmt.Println("Formule: S=pi*r^2")
	fmt.Printf("Circle Area is %f32\n", circleArea)

}

// Print Perimetre Gonibuchluk

func ppg(a, b float32) {
	pg, err := calculatePerimetrGonuburchluk(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Data:= a=%d, b=%d \n", int(a), int(b))
	fmt.Println("P=2*(a+b)")
	fmt.Printf("P=%d", int(pg))
}

// Print Area Gonuburchluk

func pag(a, b float32) {
	ag, err := calculateGuniburclukArea(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Data:= a=%d, b=%d \n", int(a), int(b))
	fmt.Println("S=a*b")
	fmt.Printf("P=%d", int(ag))
}

// CalculateCircleArea
func calculateCircleArea(radius int) (float32, error) {
	if radius > 0 {
		return float32(radius) * float32(radius) * math.Pi, nil
	}

	e := errors.New("Radiuse can't be uint")
	return float32(0), e
}

//Calculategoniburchluk

func calculatePerimetrGonuburchluk(a, b float32) (float32, error) {
	if a <= 0 || b <= 0 {
		massage := errors.New("Data invalide")
		return float32(0), massage
	}
	return float32(a+b) * float32(2), nil
}

// Calculate Gonuburchluk meydan
func calculateGuniburclukArea(a, b float32) (float32, error) {
	if a <= 0 || b <= 0 {
		massage := errors.New("Data invalide")
		return float32(0), massage
	}
	return float32(a * b), nil
}

// ucburchluk P
func calculatePerimetr3(a, b, c float32) (float32, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		massage := errors.New("Data invalide")
		return float32(0), massage
	}
	return float32(a+b) + float32(c), nil
}

// Calculate uchburchluk S
func calculate3Area(a, h float32) (float32, error) {
	if a <= 0 || h <= 0 {
		massage := errors.New("Data invalide")
		return float32(0), massage
	}
	return float32(a * h * 1 / 2), nil
}
