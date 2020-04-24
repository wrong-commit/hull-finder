package main

type hullFinder interface {
	findPerimeter(h *hull) points
}

type convexHullFinder struct {
}

type concaveHullFinder struct {
}

func (hf convexHullFinder) findPerimeter(h *hull) points {
	// TODO: implement convexHullFinder.findPerimeter
	return h.allPoints
}

func (hf concaveHullFinder) findPerimeter(h *hull) points {
	// TODO: implement concaveHullFinder.findPerimeter
	return h.allPoints
}
