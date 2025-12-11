package optimizer

import "math"

type NearestNeighborOptimizer struct {
	Calculator DistanceCalculator
}

func (n *NearestNeighborOptimizer) Optimize(startPoint Point, destinations []Point) []Point {
	if len(destinations) == 0 {
		return []Point{startPoint}
	}

	route := []Point{startPoint}

	unvisited := make(map[string]Point)
	for _, d := range destinations {
		unvisited[d.ID] = d
	}

	currentPoint := startPoint

	for len(unvisited) > 0 {
		var nextPoint Point
		minDistance := math.MaxFloat64

		for _, nextCandidate := range unvisited {
			dist := n.Calculator.Calculate(currentPoint, nextCandidate)

			if dist < minDistance {
				minDistance = dist
				nextPoint = nextCandidate
			}
		}

		route = append(route, nextPoint)
		delete(unvisited, nextPoint.ID)
		currentPoint = nextPoint
	}

	return route
}
