# dim2000 fixtures

Coordinates in `[0, 2000] × [0, 2000]`.

| File | Aspect |
|------|--------|
| `01_interior_cloud.csv` | Dense interior noise; hull should be a few extremes (306 points) |
| `02_circle_ring.csv` | Points on a circle — nearly all vertices on the hull (72 points) |
| `03_bounding_square.csv` | Four corners of [0,2000]^2 plus interior scatter (84 points) |
| `04_upper_heavy.csv` | Dense upper arc; sparse lower edge (68 points) |
| `05_collinear_base.csv` | Many collinear points on the bottom edge (58 points) |
| `06_diamond.csv` | Diamond perimeter in the middle of the plane (78 points) |
| `07_two_clusters.csv` | Two separated blobs — hull bridges both (100 points) |
| `08_axis_cross.csv` | Points along horizontal/vertical axes through center (85 points) |
| `09_thin_vertical.csv` | Nearly vertical strip (stresses extreme-X handling) (72 points) |
| `10_grid_lattice.csv` | Regular lattice — hull is the outer rectangle corners (121 points) |
