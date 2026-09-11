package main

import (
	"fmt"
)

type convexHullFinder struct{}

// findPerimeter finds the convex hull perimeter of h.allPoints.
//
// Approach (from starter notes):
//  1. Find least-X and most-X endpoints A and B.
//  2. Recursively build the upper chain via highest points between segments.
//  3. Recursively build the lower chain via lowest points between segments.
//  4. Return unique perimeter points (order: upper A→B, then lower B→A excluding endpoints).
func (hf convexHullFinder) findPerimeter(h *hull) points {
	if h == nil || len(h.allPoints) == 0 {
		return points{}
	}
	if len(h.allPoints) == 1 {
		return points{h.allPoints[0]}
	}

	leastX, mostX := extremeX(h.allPoints)
	fmt.Printf("Least X: %v, Most X: %v\n", leastX, mostX)

	if leastX == mostX {
		// All points share the same X (or only one distinct extreme); fall back to Y extremes.
		return verticalHull(h.allPoints)
	}

	upper := points{leastX}
	upper = append(upper, buildUpperChain(h, leastX, mostX)...)
	upper = append(upper, mostX)

	lowerMid := buildLowerChain(h, leastX, mostX)
	// lowerMid is left-to-right between A and B; append right-to-left excluding endpoints.
	for i := len(lowerMid) - 1; i >= 0; i-- {
		upper = append(upper, lowerMid[i])
	}
	return upper
}

func extremeX(ps points) (leastX, mostX point) {
	leastX, mostX = ps[0], ps[0]
	for _, p := range ps[1:] {
		if p[0] < leastX[0] || (p[0] == leastX[0] && p[1] < leastX[1]) {
			leastX = p
		}
		if p[0] > mostX[0] || (p[0] == mostX[0] && p[1] < mostX[1]) {
			mostX = p
		}
	}
	return leastX, mostX
}

func verticalHull(ps points) points {
	lo, hi := ps[0], ps[0]
	for _, p := range ps[1:] {
		if p[1] < lo[1] {
			lo = p
		}
		if p[1] > hi[1] {
			hi = p
		}
	}
	if lo == hi {
		return points{lo}
	}
	return points{lo, hi}
}

// buildUpperChain returns interior upper-hull points strictly between A and B (left→right).
func buildUpperChain(h *hull, a, b point) points {
	c, ok := farthestOnSide(h, a, b, 1)
	if !ok {
		return points{}
	}
	left := buildUpperChain(h, a, c)
	right := buildUpperChain(h, c, b)
	out := append(left, c)
	return append(out, right...)
}

// buildLowerChain returns interior lower-hull points strictly between A and B (left→right).
func buildLowerChain(h *hull, a, b point) points {
	c, ok := farthestOnSide(h, a, b, -1)
	if !ok {
		return points{}
	}
	left := buildLowerChain(h, a, c)
	right := buildLowerChain(h, c, b)
	out := append(left, c)
	return append(out, right...)
}

// farthestOnSide finds the point farthest from line ab on the requested side.
// side +1 = above (left of a→b), -1 = below (right of a→b).
func farthestOnSide(h *hull, a, b point, side int) (point, bool) {
	best := point{}
	bestDist := 0
	found := false
	for _, p := range h.allPoints {
		if p == a || p == b {
			continue
		}
		cr := cross(a, b, p)
		if side > 0 && cr <= 0 {
			continue
		}
		if side < 0 && cr >= 0 {
			continue
		}
		dist := cr
		if dist < 0 {
			dist = -dist
		}
		if !found || dist > bestDist {
			best = p
			bestDist = dist
			found = true
		}
	}
	return best, found
}

// cross returns (b-a)×(c-a). Positive => c is left of directed edge a→b.
func cross(a, b, c point) int {
	return (b[0]-a[0])*(c[1]-a[1]) - (b[1]-a[1])*(c[0]-a[0])
}
