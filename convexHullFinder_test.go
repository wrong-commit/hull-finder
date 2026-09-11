package main

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_6By6(t *testing.T) {
	expectedPoints := points{
		{0, 1},
		{1, 5},
		{2, 6},
		{3, 5},
		{4, 1},
		{2, 0},
	}
	h := newHull(points{
		{0, 1},
		{1, 2},
		{1, 5},
		{2, 0},
		{2, 2},
		{2, 3},
		{2, 6},
		{3, 2},
		{3, 5},
		{4, 1},
	})

	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "6by6", h)

	assertPerimeterContains(t, h.perimiterPoints, expectedPoints)
}

func Test_NoPoints(t *testing.T) {
	h := newHull(points{})
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "no_points", h)

	if len(h.perimiterPoints) != 0 {
		t.Fatalf("expected empty perimeter, got %v", h.perimiterPoints)
	}
}

func Test_SinglePoint(t *testing.T) {
	h := newHull(points{{3, 4}})
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "single_point", h)

	assertPerimeterContains(t, h.perimiterPoints, points{{3, 4}})
}

func Test_GreatestYIsLeastX(t *testing.T) {
	// Leftmost point is also the global max Y.
	h := newHull(points{
		{0, 5}, // least X and greatest Y
		{1, 1},
		{2, 2},
		{3, 0},
		{4, 1},
	})
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "greatestY_is_leastX", h)

	assertPerimeterContains(t, h.perimiterPoints, points{
		{0, 5},
		{4, 1},
		{3, 0},
		{1, 1},
	})
}

func Test_GreatestYIsMostX(t *testing.T) {
	// Rightmost point is also the global max Y.
	h := newHull(points{
		{0, 1},
		{1, 0},
		{2, 2},
		{3, 1},
		{4, 5}, // most X and greatest Y
	})
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "greatestY_is_mostX", h)

	assertPerimeterContains(t, h.perimiterPoints, points{
		{0, 1},
		{4, 5},
		{3, 1},
		{1, 0},
	})
}

func Test_LeastYIsLeastX(t *testing.T) {
	// Leftmost point is also the global min Y.
	h := newHull(points{
		{0, 0}, // least X and least Y
		{1, 3},
		{2, 4},
		{3, 2},
		{4, 1},
	})
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "leastY_is_leastX", h)

	assertPerimeterContains(t, h.perimiterPoints, points{
		{0, 0},
		{1, 3},
		{2, 4},
		{4, 1},
	})
}

func Test_LeastYIsMostX(t *testing.T) {
	// Rightmost point is also the global min Y.
	h := newHull(points{
		{0, 2},
		{1, 4},
		{2, 3},
		{3, 1},
		{4, 0}, // most X and least Y
	})
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "leastY_is_mostX", h)

	assertPerimeterContains(t, h.perimiterPoints, points{
		{0, 2},
		{1, 4},
		{2, 3},
		{4, 0},
	})
}

func Test_LoadCSVAndVisualize(t *testing.T) {
	path := filepath.Join("testdata", "sample_points.csv")
	ps, err := LoadPointsFromCSV(path)
	if err != nil {
		t.Fatalf("LoadPointsFromCSV: %v", err)
	}
	h := newHull(ps)
	hf := convexHullFinder{}
	h.perimiterPoints = hf.findPerimeter(h)
	writeHullImage(t, "csv_sample", h)

	if len(h.allPoints) == 0 {
		t.Fatal("expected points from CSV")
	}
	if len(h.perimiterPoints) == 0 {
		t.Fatal("expected non-empty perimeter for sample CSV")
	}
}

func Test_LoadCSV_BadRow(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.csv")
	if err := os.WriteFile(path, []byte("1,2\nnot-a-number,3\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := LoadPointsFromCSV(path)
	if err == nil {
		t.Fatal("expected error for bad CSV row")
	}
}

func assertPerimeterContains(t *testing.T, got, expected points) {
	t.Helper()
	if len(got) != len(expected) {
		t.Errorf("perimeter length %d, want %d; got %v", len(got), len(expected), got)
	}
	for i, ep := range expected {
		found := false
		for _, gp := range got {
			if ep == gp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expectedPoints[%d]=%v not found in perimeter %v", i, ep, got)
		}
	}
}
