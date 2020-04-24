package main

import (
	"testing"
)

func Test_6By6(t *testing.T) {
	expectedPoints := points{
		point{0, 1},
		point{1, 2},
		point{4, 1},
		point{3, 5},
		point{2, 6},
		point{1, 5},
	}
	t.Log("Testing 6x6 arrangment")
	hull := hull{allPoints: points{
		point{0, 1},
		point{1, 2},
		point{1, 5},
		point{2, 0},
		point{2, 2},
		point{2, 3},
		point{2, 6},
		point{3, 2},
		point{3, 5},
		point{4, 1},
	},
	}

	hf := convexHullFinder{}

	hull.perimiterPoints = hf.findPerimeter(&hull)
	// expected convex

	if len(hull.perimiterPoints) != len(expectedPoints) {
		t.Errorf("Did not get the same number of elements in returned Permeter, expected %d", len(expectedPoints))
		return
	}

	i := 0
	for i < len(expectedPoints) {
		var inLoop bool = false
		// iterate expected points, check if this point appears in our perimeter
		ii := 0
		for ii < len(hull.perimiterPoints) {
			if expectedPoints[i] == hull.perimiterPoints[ii] {
				inLoop = true
			}
			ii++
		}
		if !inLoop {
			t.Errorf("ExpectedPoint[%d] was not found in Perimeter", i)
			return
		}
		i++
	}
}

// TODO: test for what happens when greatestY is leastX
// TODO: test for what happens when greatestY is mostX
// TODO: test for what happens when leastestY is leastX
// TODO: test for what happens when leastestY is mostX
// TODO: test for what happens when no points
