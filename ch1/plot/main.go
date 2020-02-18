package main

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	simpleGraph()
	sinAndCos()
}

func simpleGraph() {
	x := arrayRange(0, 6, 0.1)
	y := arraySin(x)
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	plotData(p, x, y, "sin")
	p.Save(10*vg.Inch, 6*vg.Inch, "sin_function.png")
}

func sinAndCos() {
	x := arrayRange(0, 6, 0.1)
	ys := arraySin(x)
	yc := arrayCos(x)

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "sin & cos"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"
	plotData(p, x, ys, "sin")
	plotData(p, x, yc, "cos")
	p.Save(10*vg.Inch, 6*vg.Inch, "sin_cos_function.png")
}

func arrayRange(start, stop, step float64) []float64 {
	data := make([]float64, 0, int((stop-start)/step))

	for i := start; i <= stop; i += step {
		data = append(data, i)
	}
	return data
}

func arraySin(x []float64) []float64 {
	y := make([]float64, 0, len(x))
	for _, n := range x {
		y = append(y, math.Sin(n))
	}
	return y
}

func arrayCos(x []float64) []float64 {
	y := make([]float64, 0, len(x))
	for _, n := range x {
		y = append(y, math.Cos(n))
	}
	return y
}

func plotData(p *plot.Plot, x, y []float64, leg string) {

	p.Add(plotter.NewGrid())

	data0 := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data0[i].X = x[i]
		data0[i].Y = y[i]
	}

	data, err := plotter.NewScatter(data0)
	if err != nil {
		log.Fatal(err)
	}
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	data.GlyphStyle.Color = plotutil.Color(int(n.Int64()))
	data.Shape = &draw.PyramidGlyph{}
	p.Add(data)
	p.Legend.Add(leg, data)
}
