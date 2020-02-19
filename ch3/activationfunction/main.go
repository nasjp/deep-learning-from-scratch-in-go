package main

import (
	"log"
	"math"

	"github.com/nasjp/deep-learning-from-scratch-in-go/tools"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func main() {
	x := tools.NewVecDenseRange(-5.0, 5.0, 0.1)

	ystep := step(x)
	ysigmoid := sigmoid(x)
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	tools.PlotData(p, x.RawVector().Data, ystep.RawVector().Data, "step")
	tools.PlotData(p, x.RawVector().Data, ysigmoid.RawVector().Data, "sigmoid")
	p.Save(10*vg.Inch, 6*vg.Inch, "activation_function.png")
}

func step(x1 *mat.VecDense) *mat.VecDense {
	f := func(x float64) float64 {
		if x > 0 {
			return 1
		}
		return 0
	}

	v := mat.NewVecDense(x1.Len(), nil)

	for i := 0; i < x1.Len(); i++ {
		v.SetVec(i, f(x1.AtVec(i)))
	}

	return v
}

func sigmoid(x1 *mat.VecDense) *mat.VecDense {
	f := func(x float64) float64 {
		return 1 / (1 + math.Exp(-x))
	}

	v := mat.NewVecDense(x1.Len(), nil)

	for i := 0; i < x1.Len(); i++ {
		v.SetVec(i, f(x1.AtVec(i)))
	}

	return v
}
