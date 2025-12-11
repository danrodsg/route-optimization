package api

import (
	"net/http"

	"github.com/danrodsg/route-optimization/optimizer"
	"github.com/gin-gonic/gin"
)

func OptimizeRouteHandler(c *gin.Context) {
	var req OptimizationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	opt := optimizer.NearestNeighborOptimizer{
		Calculator: optimizer.EuclideanDistance{},
	}

	start := optimizer.Point{
		ID: req.StartPoint.ID, Latitude: req.StartPoint.Latitude, Longitude: req.StartPoint.Longitude,
	}

	destinations := make([]optimizer.Point, len(req.Destinations))
	for i, p := range req.Destinations {
		destinations[i] = optimizer.Point{ID: p.ID, Latitude: p.Latitude, Longitude: p.Longitude}
	}

	rotaOtimizada := opt.Optimize(start, destinations)

	c.JSON(http.StatusOK, gin.H{
		"optimized_route": rotaOtimizada,
	})
}
