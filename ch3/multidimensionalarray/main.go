package main

import (
	"fmt"

	"github.com/nasjp/deep-learning-from-scratch-in-go/tools"
	"gonum.org/v1/gonum/mat"
)

func main() {
	dimensionArray()
	matrixDot()
	neuralNetwork()
}

func dimensionArray() {
	vecA := mat.NewVecDense(4, []float64{1, 2, 3, 4})
	tools.MatPrint(vecA)
	/*
		⎡1⎤
		⎢2⎥
		⎢3⎥
		⎣4⎦
	*/
	fmt.Println(vecA.Dims())
	// 4 1
	matA := mat.NewDense(1, 4, []float64{1, 2, 3, 4})
	tools.MatPrint(matA)
	/*
		[1  2  3  4]
	*/
	fmt.Println(matA.Dims())
	// 1 4

	B := mat.NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6})
	tools.MatPrint(B)
	/*
		⎡1  2⎤
		⎢3  4⎥
		⎣5  6⎦
	*/
	fmt.Println(B.Dims())
	// 3 2
}

func matrixDot() {
	func() {
		A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
		fmt.Println(A.Dims())
		// 2 2

		B := mat.NewDense(2, 2, []float64{5, 6, 7, 8})
		fmt.Println(B.Dims())
		// 2 2

		C := mat.NewDense(2, 2, nil)
		// 行列の内積
		C.Product(A, B)
		tools.MatPrint(C)
		/*
			⎡19  22⎤
			⎣43  50⎦
		*/
	}()

	func() {
		A := mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
		fmt.Println(A.Dims())
		// 2 3

		B := mat.NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6})
		fmt.Println(B.Dims())
		// 3 2

		// C := mat.NewDense(2, 2, nil)
		C := &mat.Dense{}
		// 行列の内積
		C.Product(A, B)
		tools.MatPrint(C)
		/*
			⎡22  28⎤
			⎣49  64⎦
		*/
	}()

	func() {
		A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
		fmt.Println(A.Dims())
		// 2 2

		B := mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
		fmt.Println(B.Dims())
		// 2 3

		C := &mat.Dense{}
		fmt.Println(C.Dims())
		// B と A の行と列が一致しない
		// C.Product(B, A)
		// tools.MatPrint(C)
		// panic: mat: dimension mismatch
	}()

	func() {
		A := mat.NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6})
		fmt.Println(A.Dims())
		// 3 2

		B := mat.NewDense(2, 1, []float64{7, 8})
		fmt.Println(B.Dims())
		// 2 1

		C := mat.NewDense(3, 1, nil)
		C.Product(A, B)
		tools.MatPrint(C)
		/*
			⎡23⎤
			⎢53⎥
			⎣83⎦
		*/
	}()
}

func neuralNetwork() {
	X := mat.NewDense(1, 2, []float64{1, 2})
	fmt.Println(X.Dims())
	// 1 2

	W := mat.NewDense(2, 3, []float64{1, 3, 5, 2, 4, 6})
	tools.MatPrint(W)
	/*
		⎡1  3  5⎤
		⎣2  4  6⎦
	*/
	fmt.Println(W.Dims())
	// 2 3

	Y := mat.NewDense(1, 3, nil)

	Y.Product(X, W)
	tools.MatPrint(Y)
	/*
		[5  11  17]
	*/
}
