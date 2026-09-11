package main

import (
	"fmt"
	"os"
)

func main() {
	path := "testdata/sample_points.csv"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ps, err := LoadPointsFromCSV(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load points: %v\n", err)
		os.Exit(1)
	}

	h := newHull(ps)
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)

	out := "hull.png"
	if len(os.Args) > 2 {
		out = os.Args[2]
	}
	if err := WriteHullImage(out, h); err != nil {
		fmt.Fprintf(os.Stderr, "write image: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s (%d points, %d on perimeter)\n", out, len(h.allPoints), len(h.perimiterPoints))
}
