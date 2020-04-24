package main

import "fmt"

type hullFinder interface {
	findPerimeter(h *hull) points
}

/**
This method will return the leastX argument if no higher Y point could be found
*/
func findHighestPoint(h *hull, leastX *point, mostX *point) *point {
	i := 0
	lastGreatestY := leastX
	// iterate over all points where X is between leastX, mostX.
	// return point with greatest Y
	for i < len(h.allPoints) {
		i++
		var p point = h.allPoints[i-1]

		if !(p[0] >= leastX[0] && p[0] <= mostX[0]) {
			fmt.Printf("TRACE: Point %v outside our X points\n", p)
			continue
		}

		if p[1] >= lastGreatestY[1] {
			fmt.Printf("TRACE: Point %v is new Highest Y Point\n", p)
			lastGreatestY = &p
		}
	}
	return lastGreatestY
}

/**
This method will return the leastX argument if no lower Y point could be found
*/
func findLowestPoint(h *hull, leastX *point, mostX *point) *point {
	i := 0
	lastLowestY := leastX
	// iterate over all points where X is between leastX, mostX.
	// return point with lowest Y
	for i < len(h.allPoints) {
		var p point = h.allPoints[i]

		if !(p[0] >= leastX[0] && p[0] <= mostX[0]) {
			// fmt.Printf("TRACE: Point %v outside our X points\n", p)
			continue
		}

		if p[1] <= lastLowestY[1] {
			// fmt.Printf("TRACE: Point %v is new Lowest Y Point\n", p)
			lastLowestY = &p
		}
		i++
	}
	return lastLowestY
}
