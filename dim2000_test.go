package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_Dim2000Fixtures(t *testing.T) {
	dir := filepath.Join("testdata", "dim2000")
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read %s: %v", dir, err)
	}

	hf := convexHullFinder{}
	found := 0
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".csv") {
			continue
		}
		found++
		name := strings.TrimSuffix(e.Name(), ".csv")
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(dir, e.Name())
			ps, err := LoadPointsFromCSV(path)
			if err != nil {
				t.Fatalf("load: %v", err)
			}
			if len(ps) == 0 {
				t.Fatal("expected points")
			}
			for _, p := range ps {
				if p[0] < 0 || p[0] > 2000 || p[1] < 0 || p[1] > 2000 {
					t.Fatalf("point %v outside [0,2000]x[0,2000]", p)
				}
			}

			h := newHull(ps)
			h.perimiterPoints = hf.findPerimeter(h)
			writeHullImage(t, "dim2000_"+name, h)

			if len(h.perimiterPoints) < 2 && len(ps) >= 2 {
				t.Fatalf("expected a perimeter for %d points, got %v", len(ps), h.perimiterPoints)
			}
			// Every hull vertex must be an input point.
			in := map[point]bool{}
			for _, p := range ps {
				in[p] = true
			}
			for _, p := range h.perimiterPoints {
				if !in[p] {
					t.Fatalf("hull point %v not in input set", p)
				}
			}
			t.Logf("%d input → %d hull", len(ps), len(h.perimiterPoints))
		})
	}
	if found < 10 {
		t.Fatalf("expected at least 10 CSV fixtures, found %d", found)
	}
}
