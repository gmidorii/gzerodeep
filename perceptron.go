package main

import (
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func perceptron() {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	err = plotutil.AddLinePoints(p, "perceptron", points(-5.0, 5.0, 0.1))
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "perceptron.png"); err != nil {
		log.Fatal(err)
	}
}

func points(xmin, xmax, inc float64) plotter.XYs {
	pts := make(plotter.XYs, 100)
	var counter = 0
	for i := xmin; i < xmax; i = i + inc {
		pts[counter].X = i
		pts[counter].Y = stepFunc(i)
		counter++
	}

	return pts
}

func stepFunc(x float64) float64 {
	if x > 0 {
		return 1.0
	} else {
		return 0.0
	}
}
