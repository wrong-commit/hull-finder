# Hull Finder

![leastY_is_leastX example](docs/screenshots/leastY_is_leastX.png)

Computes the convex hull of a 2D point set and writes a PNG so you can inspect the result. No GUI — image output only, stdlib drawing.

This is the **outer perimeter** only: points strictly inside the hull are drawn but are not part of the red path. That is expected (see [SPECS.md](SPECS.md) investigation for `leastY_is_leastX`).

## How to build

Requires Go 1.17+.

```bash
go build -o hull-finder .
```

## How to run tests

```bash
go test ./...
```

Each hull test writes a PNG via `writeHullImage` under `testdata/out/` (e.g. `6by6.png`, `csv_sample.png`). Open those files to review cases.

## How to run program

```bash
go run . [points.csv] [out.png]
```

Defaults: `testdata/sample_points.csv` → `hull.png`.

```bash
go run . testdata/sample_points.csv hull.png
```

Or after building:

```bash
./hull-finder testdata/sample_points.csv hull.png
```

## How to populate CSV

One point per row as `x,y` (integers). Extra columns are ignored. Blank lines and `#` comments are skipped.

```csv
# sample
0,1
1,2
2,6
4,1
```

See `testdata/sample_points.csv` for a small example.

### Large 2000×2000 fixtures

Ten scenarios live under `testdata/dim2000/` (coords in `[0, 2000]`). Catalog: `testdata/dim2000/README.md`.

```bash
go run . testdata/dim2000/02_circle_ring.csv circle.png
go test -run Test_Dim2000Fixtures   # writes testdata/out/dim2000_*.png
go run testdata/gendata.go          # regenerate CSVs
```

## How to consume the generated images

Open the PNG in any image viewer:

- CLI: path from the second argument (default `hull.png`)
- Tests: `testdata/out/*.png` after `go test ./...`

### Legend

| Color | Meaning |
|-------|---------|
| **Blue** (larger outline) | Every input point |
| **Red** (smaller filled circle) | Hull vertex, drawn inside its blue outline |
| **Red** lines | Hull edges connecting consecutive perimeter points |
| Light grid / frame | Coordinate guide only |

Overlap: a hull point shows as a **red fill inside a blue outline**. Interior-only points are blue outline only.

Y increases upward in the plot (image Y is flipped so the picture matches normal math coordinates). Scale auto-fits the point bounds with a margin.
