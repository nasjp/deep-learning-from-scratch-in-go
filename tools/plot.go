package tools

import (
	"crypto/rand"
	"log"
	"math/big"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg/draw"
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
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	data.GlyphStyle.Color = plotutil.Color(int(n.Int64()))
	data.Shape = &draw.PyramidGlyph{}
	p.Add(data)
	p.Legend.Add(leg, data)
}
