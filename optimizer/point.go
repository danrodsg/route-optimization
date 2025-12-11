package optimizer

type Point struct {
	ID        string
	Latitude  float64
	Longitude float64
}

type DistanceCalculator interface {
	Calculate(p1, p2 Point) float64
}
