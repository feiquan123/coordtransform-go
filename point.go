package coordtransform

import (
	"fmt"
)

// Point point struct
type Point struct {
	Lon, Lat float64
}

func (p Point) String() string {
	return fmt.Sprintf("%.12f,%.12f", p.Lon, p.Lat)
}

// Equals this point is equivalent to the other point
func (p Point) Equals(other Point) bool {
	return p.Lon == other.Lon && p.Lat == other.Lat
}

// Round round this point to the nearest
func (p Point) Round(l int) Point {
	var a float64 = 1
	for i := 0; i < l; i++ {
		a *= 10
	}
	return Point{
		Lon: float64(int(p.Lon*a+0.5)) / a,
		Lat: float64(int(p.Lat*a+0.5)) / a,
	}
}

// RoundString limit number after dot
func (p Point) RoundString(l int) string {
	if l < 0 {
		l = 7
	}
	s := fmt.Sprintf("%%.%df,%%.%df", l, l)
	return fmt.Sprintf(s, p.Lon, p.Lat)
}
