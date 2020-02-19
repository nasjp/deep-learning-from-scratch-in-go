package tools

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func MatPrint(x mat.Matrix) {
	fa := mat.Formatted(x, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
