# HullFinder Specs

Living list of requirements. Update as work progresses.

## Visualization
- [x] Render 2D point sets and hull perimeter as PNG (no Win32 GUI)
- [x] Use Go stdlib only for drawing/encoding (`image`, `image/png`, `image/color`)
- [x] Show all input points and, when present, perimeter/hull edges
- [x] Overlap visibility: larger blue outlines under smaller filled red hull markers
- [x] Single shared helper usable from all unit tests: write current image state to a given file path (`writeHullImage` → `WriteHullImage`)
- [x] Helper must work for any point set (inline, CSV-loaded, or fixture)

## Input
- [x] Load points from CSV (`LoadPointsFromCSV` / `ReadPointsCSV`)
- [x] Support visualizing any arbitrary set of points via that loader + helper
- [x] CLI: `go run . [points.csv] [out.png]`

## Correctness / error cases (tests)
- [x] Empty / no points
- [x] Single point
- [x] Greatest Y coincides with least X
- [x] Greatest Y coincides with most X
- [x] Least Y coincides with least X
- [x] Least Y coincides with most X
- [x] Existing happy path: 6×6 arrangement (`Test_6By6`) — expected set corrected to true convex hull
- [x] Existing helpers: `findHighestPoint` / `findLowestPoint`
- [x] Bad CSV row returns an error
- [x] CSV sample loads and visualizes

## Algorithm / starter code
- [x] Keep existing types (`point`, `points`, `hull`, `hullFinder`, `convexHullFinder`)
- [x] `findPerimeter` must compile and return perimeter `points`
- [x] Preserve `findHighestPoint` / `findLowestPoint` for calculations
- [x] Upper/lower chains use farthest-from-edge (cross product), not max/min Y alone

## Engineering constraints
- [x] Cheap: PNG file output is sufficient
- [x] Changes must be inspectable via generated images (tests write PNGs under `testdata/out/`)
- [x] Prefer minimal deps (stdlib only)

## Open / follow-ups
- [ ] Quiet debug `fmt.Printf` noise in hull helpers during tests
- [ ] Optional: visualize intermediate recursion steps
- [ ] Optional: label coordinates on PNG
