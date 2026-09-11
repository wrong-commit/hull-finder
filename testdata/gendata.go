//go:build ignore

// Generate large-coordinate CSV fixtures under testdata/dim2000/.
//
//	go run testdata/gendata.go
package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

const dim = 2000

func main() {
	outDir := filepath.Join("testdata", "dim2000")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		fatal(err)
	}

	fixtures := []struct {
		name string
		pts  [][2]int
		desc string
	}{
		{"01_interior_cloud", interiorCloud(), "Dense interior noise; hull should be a few extremes"},
		{"02_circle_ring", circleRing(72), "Points on a circle — nearly all vertices on the hull"},
		{"03_bounding_square", boundingSquare(), "Four corners of [0,2000]^2 plus interior scatter"},
		{"04_upper_heavy", upperHeavy(), "Dense upper arc; sparse lower edge"},
		{"05_collinear_base", collinearBase(), "Many collinear points on the bottom edge"},
		{"06_diamond", diamond(12), "Diamond perimeter in the middle of the plane"},
		{"07_two_clusters", twoClusters(), "Two separated blobs — hull bridges both"},
		{"08_axis_cross", axisCross(), "Points along horizontal/vertical axes through center"},
		{"09_thin_vertical", thinVertical(), "Nearly vertical strip (stresses extreme-X handling)"},
		{"10_grid_lattice", gridLattice(10), "Regular lattice — hull is the outer rectangle corners"},
	}

	readme := "# dim2000 fixtures\n\nCoordinates in `[0, 2000] × [0, 2000]`.\n\n| File | Aspect |\n|------|--------|\n"
	for _, f := range fixtures {
		path := filepath.Join(outDir, f.name+".csv")
		if err := writeCSV(path, f.desc, f.pts); err != nil {
			fatal(err)
		}
		readme += fmt.Sprintf("| `%s.csv` | %s (%d points) |\n", f.name, f.desc, len(f.pts))
		fmt.Printf("wrote %s (%d points)\n", path, len(f.pts))
	}
	if err := os.WriteFile(filepath.Join(outDir, "README.md"), []byte(readme), 0o644); err != nil {
		fatal(err)
	}
}

func writeCSV(path, desc string, pts [][2]int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "# %s\n# coords in [0,%d] x [0,%d]\n", desc, dim, dim)
	for _, p := range pts {
		fmt.Fprintf(f, "%d,%d\n", p[0], p[1])
	}
	return nil
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > dim {
		return dim
	}
	return v
}

func interiorCloud() [][2]int {
	// Deterministic pseudo-random fill well inside the bounds.
	pts := make([][2]int, 0, 320)
	for i := 0; i < 300; i++ {
		x := 200 + (i*97+13)%1601
		y := 200 + (i*53+29)%1601
		pts = append(pts, [2]int{x, y})
	}
	// Guaranteed extremes so the hull is interesting.
	pts = append(pts,
		[2]int{0, 400},
		[2]int{2000, 600},
		[2]int{800, 0},
		[2]int{1200, 2000},
		[2]int{50, 1900},
		[2]int{1950, 50},
	)
	return pts
}

func circleRing(n int) [][2]int {
	cx, cy, r := 1000.0, 1000.0, 900.0
	pts := make([][2]int, 0, n)
	for i := 0; i < n; i++ {
		a := 2 * math.Pi * float64(i) / float64(n)
		x := clamp(int(math.Round(cx + r*math.Cos(a))))
		y := clamp(int(math.Round(cy + r*math.Sin(a))))
		pts = append(pts, [2]int{x, y})
	}
	return pts
}

func boundingSquare() [][2]int {
	pts := [][2]int{
		{0, 0}, {2000, 0}, {2000, 2000}, {0, 2000},
	}
	for i := 0; i < 80; i++ {
		pts = append(pts, [2]int{
			300 + (i*41+7)%1400,
			300 + (i*73+11)%1400,
		})
	}
	return pts
}

func upperHeavy() [][2]int {
	pts := make([][2]int, 0, 120)
	// Dense samples along a high arc.
	for i := 0; i <= 60; i++ {
		t := float64(i) / 60
		x := int(math.Round(t * dim))
		y := int(math.Round(1400 + 500*math.Sin(math.Pi*t)))
		pts = append(pts, [2]int{clamp(x), clamp(y)})
	}
	// Sparse lower points.
	pts = append(pts,
		[2]int{0, 200},
		[2]int{500, 50},
		[2]int{1000, 0},
		[2]int{1500, 80},
		[2]int{2000, 150},
		[2]int{800, 900},
		[2]int{1100, 850},
	)
	return pts
}

func collinearBase() [][2]int {
	pts := make([][2]int, 0, 100)
	for x := 0; x <= dim; x += 40 {
		pts = append(pts, [2]int{x, 0})
	}
	pts = append(pts,
		[2]int{0, 2000},
		[2]int{2000, 1800},
		[2]int{400, 1600},
		[2]int{1000, 1900},
		[2]int{1600, 1400},
		[2]int{700, 600},
		[2]int{1300, 500},
	)
	return pts
}

func diamond(nPerSide int) [][2]int {
	cx, cy, r := 1000, 1000, 800
	pts := make([][2]int, 0, 4*nPerSide+30)
	for i := 0; i <= nPerSide; i++ {
		u := float64(i) / float64(nPerSide)
		ru := int(u * float64(r))
		// left → top → right → bottom → left (manhattan circle)
		pts = append(pts, [2]int{clamp(cx - r + ru), clamp(cy + ru)})
		pts = append(pts, [2]int{clamp(cx + ru), clamp(cy + r - ru)})
		pts = append(pts, [2]int{clamp(cx + r - ru), clamp(cy - ru)})
		pts = append(pts, [2]int{clamp(cx - ru), clamp(cy - r + ru)})
	}
	for i := 0; i < 30; i++ {
		pts = append(pts, [2]int{700 + (i*17)%600, 700 + (i*23)%600})
	}
	return unique(pts)
}

func twoClusters() [][2]int {
	pts := make([][2]int, 0, 120)
	for i := 0; i < 50; i++ {
		pts = append(pts, [2]int{100 + (i*13)%250, 100 + (i*19)%250})
		pts = append(pts, [2]int{1600 + (i*11)%300, 1500 + (i*17)%350})
	}
	return pts
}

func axisCross() [][2]int {
	pts := make([][2]int, 0, 120)
	for v := 0; v <= dim; v += 50 {
		pts = append(pts, [2]int{1000, v})
		pts = append(pts, [2]int{v, 1000})
	}
	pts = append(pts, [2]int{200, 200}, [2]int{1800, 200}, [2]int{200, 1800}, [2]int{1800, 1800})
	return unique(pts)
}

func thinVertical() [][2]int {
	pts := make([][2]int, 0, 80)
	for i := 0; i < 70; i++ {
		x := 990 + (i%5) // x in {990..994}
		y := (i * 29) % (dim + 1)
		pts = append(pts, [2]int{x, y})
	}
	pts = append(pts, [2]int{980, 0}, [2]int{1010, 2000})
	return pts
}

func gridLattice(step int) [][2]int {
	pts := make([][2]int, 0, (step+1)*(step+1))
	for i := 0; i <= step; i++ {
		for j := 0; j <= step; j++ {
			pts = append(pts, [2]int{i * dim / step, j * dim / step})
		}
	}
	return pts
}

func unique(pts [][2]int) [][2]int {
	seen := map[[2]int]bool{}
	out := make([][2]int, 0, len(pts))
	for _, p := range pts {
		if seen[p] {
			continue
		}
		seen[p] = true
		out = append(out, p)
	}
	return out
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
