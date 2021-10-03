package coordinate_supplier

/* NextMode determines how coordinates should be handed out:
 - in ascending order: 1, 2, 3, ...
 - in descending order: 3, 2, 1, ...
 - in random order: 2, 1, 3, ...
Coordinate system origin is the top left.
These are the first 9 points handed out when in ascending order for a 3x3 grid:
             x0          x1          x2
	   =================================
   y0 || (1) 0,0     (2) 1,0     (3) 2,0
      ||
   y1 || (4) 0,1     (5) 1,1     (6) 2,1
      ||
   y2 || (7) 0,2     (8) 1,2     (9) 2,2
*/
type NextMode uint

const (
	Asc NextMode = iota
	Desc
	Random
)
