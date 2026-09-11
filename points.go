package main

// Stores 2 coordinates, X and Y
type point [2]int

// Stores any number of Points
type points []point

// return -1 if lower than other, 0 if the same, or 1 if higher than
func (p *point) HigherThan(other *point) int {
	if p[1] < other[1] {
		return -1
	}
	if p[1] > other[1] {
		return 1
	}
	// therefore  p[1] == other[1]
	return 0
}

// return -1 if higher than other, 0 if the same, or 1 if lower than
func (p *point) LowerThan(other *point) int {
	return -1 * p.HigherThan(other)
}

// return -1 if lower than other, 0 if the same, or 1
func (p *point) IsBetween(left, right *point) bool {
	return p[0] >= left[0] && p[0] <= right[0]
}
