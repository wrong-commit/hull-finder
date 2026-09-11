package main

import (
	"testing"
)

var left = point{0, 0}
var right = point{3, -1}
var middle1 = point{1, -2}
var middle2 = point{2, 3}
var outside = point{4, 5}

func Test_findHighestPpint(t *testing.T) {
	hull := hull{allPoints: points{left, right, middle1, middle2, outside}}
	// expected output
	expected := middle2
	output := findHighestPoint(&hull, &left, &right)

	if output[0] != expected[0] {
		t.Errorf("X value wrong: %d", output[0])
	}
	if output[1] != expected[1] {
		t.Errorf("Y value wrong: %d", output[1])
	}
}
func Test_findLowestPoint(t *testing.T) {
	hull := hull{allPoints: points{left, right, middle1, middle2, outside}}
	// expected output
	expected := middle1
	output := findLowestPoint(&hull, &left, &right)

	if output[0] != expected[0] {
		t.Errorf("X value wrong: %d", output[0])
	}
	if output[1] != expected[1] {
		t.Errorf("Y value wrong: %d", output[1])
	}
}
