package tools

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Graph struct {
	P         *plot.Plot
	Functions []Function
}

type Function struct {
	X      []float64
	Y      []float64
	Legend string
}

func (g *Graph) Draw() {

	g.P.Add(plotter.NewGrid())

	d := make([]interface{}, 0, len(g.Functions)*2)
	for _, f := range g.Functions {
		data0 := make(plotter.XYs, len(f.X))
		for i := 0; i < len(f.X); i++ {
			data0[i].X = f.X[i]
			data0[i].Y = f.Y[i]
		}

		data, err := plotter.NewScatter(data0)
		if err != nil {
			log.Fatal(err)
		}
		d = append(d, f.Legend, data)
	}

	if err := plotutil.AddLines(g.P, d...); err != nil {
		log.Fatal(err)
	}
}

func PlotData(p *plot.Plot, x, y []float64, leg string) {

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

	if err := plotutil.AddLines(p, leg, data); err != nil {
		log.Fatal(err)
	}
}
