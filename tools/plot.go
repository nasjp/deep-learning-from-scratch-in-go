package tools

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

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
