package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	perceptron()
}

func perceptron() {
	fmt.Println(AND(0, 0)) // 0
	fmt.Println(AND(0, 1)) // 0
	fmt.Println(AND(1, 0)) // 0
	fmt.Println(AND(1, 1)) // 1

	fmt.Println(NAND(0, 0)) // 1
	fmt.Println(NAND(0, 1)) // 1
	fmt.Println(NAND(1, 0)) // 1
	fmt.Println(NAND(1, 1)) // 0

	fmt.Println(OR(0, 0)) // 0
	fmt.Println(OR(0, 1)) // 1
	fmt.Println(OR(1, 0)) // 1
	fmt.Println(OR(1, 1)) // 1

	fmt.Println(XOR(0, 0)) // 0
	fmt.Println(XOR(0, 1)) // 1
	fmt.Println(XOR(1, 0)) // 1
	fmt.Println(XOR(1, 1)) // 0
}

func AND(x1, x2 float64) float64 {
	x := mat.NewVecDense(2, []float64{x1, x2})
	w := mat.NewVecDense(2, []float64{0.5, 0.5})
	b := -0.7

	if v := mat.Dot(x, w) + b; v <= 0 {
		return 0
	}
	return 1
}

func NAND(x1, x2 float64) float64 {
	x := mat.NewVecDense(2, []float64{x1, x2})
	w := mat.NewVecDense(2, []float64{-0.5, -0.5})
	b := 0.7

	if v := mat.Dot(x, w) + b; v <= 0 {
		return 0
	}
	return 1
}

func OR(x1, x2 float64) float64 {
	x := mat.NewVecDense(2, []float64{x1, x2})
	w := mat.NewVecDense(2, []float64{0.5, 0.5})
	b := -0.2

	if v := mat.Dot(x, w) + b; v <= 0 {
		return 0
	}
	return 1
}

func XOR(x1, x2 float64) float64 {
	return AND(NAND(x1, x2), OR(x1, x2))
}
