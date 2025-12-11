package api

type OptimizationRequest struct {
    StartPoint *Point       `json:"start_point" binding:"required"`
    Destinations []Point    `json:"destinations" binding:"required"`
}

type Point struct {
    ID        string `json:"id"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}