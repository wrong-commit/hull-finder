package main

import "fmt"

type hullFinder interface {
	findPerimeter(h *hull) points
}

/**
This method will return the highest Y point between least and most X
*/
func findHighestPoint(h *hull, leastX *point, mostX *point) *point {
	lastGreatestY := *leastX
	// iterate over all points where X is between leastX, mostX.
	for i, p := range h.allPoints {
		if p.IsBetween(leastX, mostX) {
			// p less than last
			if p.HigherThan(&lastGreatestY) > 0 {
				fmt.Printf("Point %d:%v highest point between %v,%v\n", i, p, leastX, mostX)
				lastGreatestY = p
			}
		}
	}
	return &lastGreatestY
}

/**
This method will return the leastX argument if no lower Y point could be found
*/
func findLowestPoint(h *hull, leastX *point, mostX *point) *point {
	lastLowestY := *leastX
	// iterate over all points where X is between leastX, mostX.
	for i, p := range h.allPoints {
		if p.IsBetween(leastX, mostX) {
			// p less than last
			if p.LowerThan(&lastLowestY) > 0 {
				fmt.Printf("Point %d:%v lowest point between %v,%v\n", i, p, leastX, mostX)
				lastLowestY = p
			}
		}
	}
	return &lastLowestY
}
