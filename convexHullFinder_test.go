package main

import (
	"testing"
)

func test6by6(t *testing.T) {
	expectedPoints := points{
		point{0, 1},
	}
	t.Log("Testing 6x6 arrangment")
	hull := hull{allPoints: points{
		point{0, 1},
		point{1, 2},
		point{1, 4},
		point{2, 0},
		point{2, 3},
		point{2, 5},
		point{3, 4},
		point{4, 1},
	},
	}

	hf := convexHullFinder{}

	hf.findPerimeter(&hull)
	// expected convex

	if len(hull.perimiterPoints) != len(expectedPoints) {
		t.Error("Did not get the same number of elements in returned Permeter")
	}
}
