package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	colorBackground = color.RGBA{R: 250, G: 250, B: 252, A: 255}
	colorGrid       = color.RGBA{R: 230, G: 230, B: 235, A: 255}
	colorPoint      = color.RGBA{R: 40, G: 90, B: 200, A: 255}
	colorHullPoint  = color.RGBA{R: 200, G: 50, B: 50, A: 255}
	colorHullEdge   = color.RGBA{R: 220, G: 80, B: 60, A: 255}
	colorAxis       = color.RGBA{R: 180, G: 180, B: 180, A: 255}
)

const (
	defaultImgW     = 640
	defaultImgH     = 640
	defaultMargin   = 40
	bluePointRadius = 11 // outline only; larger so red fill sits inside and overlap is visible
	redPointRadius  = 6
)

// LoadPointsFromCSV reads points from a CSV file.
// Accepted row formats: "x,y" or "x,y,..." (extra columns ignored).
// Blank lines and lines starting with # are skipped.
func LoadPointsFromCSV(path string) (points, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadPointsCSV(f)
}

// ReadPointsCSV parses points from any CSV reader.
func ReadPointsCSV(r io.Reader) (points, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	cr.Comment = '#'
	cr.TrimLeadingSpace = true

	var out points
	line := 0
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("csv line %d: %w", line+1, err)
		}
		line++
		if len(rec) == 0 || (len(rec) == 1 && strings.TrimSpace(rec[0]) == "") {
			continue
		}
		if len(rec) < 2 {
			return nil, fmt.Errorf("csv line %d: need at least x,y", line)
		}
		x, err := strconv.Atoi(strings.TrimSpace(rec[0]))
		if err != nil {
			return nil, fmt.Errorf("csv line %d: bad x %q: %w", line, rec[0], err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(rec[1]))
		if err != nil {
			return nil, fmt.Errorf("csv line %d: bad y %q: %w", line, rec[1], err)
		}
		out = append(out, point{x, y})
	}
	return out, nil
}

// WriteHullImage renders h to a PNG at path. All points are drawn; perimeter
// points are highlighted and consecutive perimeter edges are connected.
// This is the shared helper for tests and CLI visualization.
func WriteHullImage(path string, h *hull) error {
	img := renderHull(h, defaultImgW, defaultImgH, defaultMargin)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func renderHull(h *hull, width, height, margin int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fill(img, colorBackground)

	if h == nil || len(h.allPoints) == 0 {
		drawAxes(img, margin)
		return img
	}

	minX, maxX, minY, maxY := bounds(h.allPoints)
	// Pad so single points / flat sets still have a visible scale.
	if minX == maxX {
		minX--
		maxX++
	}
	if minY == maxY {
		minY--
		maxY++
	}

	toPix := func(p point) (int, int) {
		nx := float64(p[0]-minX) / float64(maxX-minX)
		ny := float64(p[1]-minY) / float64(maxY-minY)
		x := margin + int(nx*float64(width-2*margin))
		// Invert Y for image coordinates (Y grows downward).
		y := height - margin - int(ny*float64(height-2*margin))
		return x, y
	}

	drawGrid(img, margin, minX, maxX, minY, maxY, toPix)
	drawAxes(img, margin)

	// Perimeter edges first (under points).
	pp := h.perimiterPoints
	if len(pp) >= 2 {
		for i := 0; i < len(pp); i++ {
			a := pp[i]
			b := pp[(i+1)%len(pp)]
			x0, y0 := toPix(a)
			x1, y1 := toPix(b)
			drawLine(img, x0, y0, x1, y1, colorHullEdge)
		}
	}

	// All input points as blue outlines; hull vertices as filled red disks on top.
	// Overlap reads as a blue ring around a red center.
	for _, p := range h.allPoints {
		x, y := toPix(p)
		drawCircleOutline(img, x, y, bluePointRadius, colorPoint)
	}
	for _, p := range h.perimiterPoints {
		x, y := toPix(p)
		drawDisk(img, x, y, redPointRadius, colorHullPoint)
	}
	return img
}

func bounds(ps points) (minX, maxX, minY, maxY int) {
	minX, maxX = ps[0][0], ps[0][0]
	minY, maxY = ps[0][1], ps[0][1]
	for _, p := range ps[1:] {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}
	return
}

func fill(img *image.RGBA, c color.Color) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			img.Set(x, y, c)
		}
	}
}

func drawAxes(img *image.RGBA, margin int) {
	b := img.Bounds()
	// Simple frame inset.
	drawLine(img, margin, margin, b.Max.X-margin, margin, colorAxis)
	drawLine(img, margin, b.Max.Y-margin, b.Max.X-margin, b.Max.Y-margin, colorAxis)
	drawLine(img, margin, margin, margin, b.Max.Y-margin, colorAxis)
	drawLine(img, b.Max.X-margin, margin, b.Max.X-margin, b.Max.Y-margin, colorAxis)
}

func drawGrid(img *image.RGBA, margin, minX, maxX, minY, maxY int, toPix func(point) (int, int)) {
	for x := minX; x <= maxX; x++ {
		x0, y0 := toPix(point{x, minY})
		_, y1 := toPix(point{x, maxY})
		drawLine(img, x0, y0, x0, y1, colorGrid)
	}
	for y := minY; y <= maxY; y++ {
		x0, y0 := toPix(point{minX, y})
		x1, _ := toPix(point{maxX, y})
		drawLine(img, x0, y0, x1, y0, colorGrid)
	}
}

func drawDisk(img *image.RGBA, cx, cy, r int, c color.Color) {
	r2 := r * r
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r2 {
				set(img, cx+dx, cy+dy, c)
			}
		}
	}
}

func drawCircleOutline(img *image.RGBA, cx, cy, r int, c color.Color) {
	// Midpoint circle; thickness ~2px for visibility on PNGs.
	for t := 0; t < 2; t++ {
		rt := r - t
		if rt <= 0 {
			continue
		}
		x, y := rt, 0
		err := 1 - rt
		for x >= y {
			plotCircleOctants(img, cx, cy, x, y, c)
			y++
			if err < 0 {
				err += 2*y + 1
			} else {
				x--
				err += 2*(y-x) + 1
			}
		}
	}
}

func plotCircleOctants(img *image.RGBA, cx, cy, x, y int, c color.Color) {
	set(img, cx+x, cy+y, c)
	set(img, cx+y, cy+x, c)
	set(img, cx-y, cy+x, c)
	set(img, cx-x, cy+y, c)
	set(img, cx-x, cy-y, c)
	set(img, cx-y, cy-x, c)
	set(img, cx+y, cy-x, c)
	set(img, cx+x, cy-y, c)
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
	dx := abs(x1 - x0)
	dy := -abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx + dy
	for {
		set(img, x0, y0, c)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

func set(img *image.RGBA, x, y int, c color.Color) {
	b := img.Bounds()
	if x < b.Min.X || x >= b.Max.X || y < b.Min.Y || y >= b.Max.Y {
		return
	}
	img.Set(x, y, c)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
