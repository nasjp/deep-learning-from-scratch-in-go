package tools

import "gonum.org/v1/gonum/mat"

func NewVecDenseRange(start, stop, step float64) *mat.VecDense {
	data := make([]float64, 0, int((stop-start)/step))

	for i := start; i <= stop; i += step {
		data = append(data, i)
	}
	return mat.NewVecDense(len(data), data)
}
