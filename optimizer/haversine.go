package optimizer

import "math"

type EuclideanDistance struct{}

func (e EuclideanDistance) Calculate(p1, p2 Point) float64 {
	latDiff := p1.Latitude - p2.Latitude
	lonDiff := p1.Longitude - p2.Longitude
	return math.Sqrt(latDiff*latDiff + lonDiff*lonDiff)
}
