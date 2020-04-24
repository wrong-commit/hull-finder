package main

import "fmt"

type convexHullFinder struct {
}

/**
Finds the perimeter of a Hull.
This does not set the Perimeter points, but returns the Perimeter points
*/
func (hf convexHullFinder) findPerimeter(h *hull) points {
	var leastX, mostX point = point{0, 0}, point{0, 0}
	// keep an array of "captured" points
	var perimeterPoints []point = make([]point, len(h.allPoints))
	// find lowestX, greatestX
	for _, p := range h.allPoints {
		if p[0] <= leastX[0] {
			leastX = p
			fmt.Printf("Found new least X of %d\n", leastX[0])
		} else if p[0] >= mostX[0] {
			mostX = p
			fmt.Printf("Found new most X of %d\n", mostX[0])
		}
	}

	perimeterPoints = append(perimeterPoints, leastX)
	perimeterPoints = append(perimeterPoints, mostX)

	// find highest point between the leastX and mostX
	// keep track of points still to check
	var pointsToCheckForHighestPoint points = points{}
	pointsToCheckForHighestPoint = append(pointsToCheckForHighestPoint, leastX)
	pointsToCheckForHighestPoint = append(pointsToCheckForHighestPoint, mostX)

	i := 0
	// this keeps adding to the array, wtf ?
	for i < len(pointsToCheckForHighestPoint) {
		i++
		latestLeastX, latestMostX := pointsToCheckForHighestPoint[i-1], pointsToCheckForHighestPoint[i]
		pointBetween := findHighestPoint(h, &latestLeastX, &latestMostX)
		if pointBetween == &latestLeastX {
			fmt.Printf("TRACE: Could not find any higher point between %v:%v\n", latestLeastX, latestMostX)
			continue
		}

		perimeterPoints = append(perimeterPoints, *pointBetween)

		// add pointBetween to be compared against latestLeastX and latestMostX

		// if the pointBetween isn't our latestLeastX, add to the
		if pointBetween != &latestLeastX {
			fmt.Printf("TRACE: Need to check %v:%v\n", *pointBetween, latestLeastX)
			pointsToCheckForHighestPoint = append(pointsToCheckForHighestPoint, *pointBetween)
			pointsToCheckForHighestPoint = append(pointsToCheckForHighestPoint, latestLeastX)
		}
		if pointBetween != &latestMostX {
			fmt.Printf("TRACE: Need to check %v:%v\n", *pointBetween, latestMostX)
			pointsToCheckForHighestPoint = append(pointsToCheckForHighestPoint, *pointBetween)
			pointsToCheckForHighestPoint = append(pointsToCheckForHighestPoint, latestMostX)
		}
	}

	return perimeterPoints
}
