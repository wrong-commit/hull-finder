package main

// Stores the points used by a convex hull
type hull struct {
	allPoints       points
	perimiterPoints points
}

func newHull(allPoints points) *hull {
	h := hull{}
	h.allPoints = allPoints
	return &h
}
