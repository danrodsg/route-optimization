
package main

import (
	"fmt"
	"github.com/danrodsg/route-optimization/optimizer"
)

func main() {
	
	deposito := optimizer.Point{ID: "DEPOT", Latitude: -23.5505, Longitude: -46.6333} // São Paulo

	campos := []optimizer.Point{
		{ID: "CAMPO_A", Latitude: -23.54, Longitude: -46.65},
		{ID: "CAMPO_B", Latitude: -23.60, Longitude: -46.70},
		{ID: "CAMPO_C", Latitude: -23.50, Longitude: -46.60},
		{ID: "CAMPO_D", Latitude: -23.58, Longitude: -46.68},
	}

	opt := optimizer.NearestNeighborOptimizer{
		Calculator: optimizer.EuclideanDistance{},
	}


	rotaOtimizada := opt.Optimize(deposito, campos)

	
	fmt.Println("--- Rota Otimizada ---")
	for i, p := range rotaOtimizada {
		fmt.Printf("%d. %s (Lat: %.4f, Lon: %.4f)\n", i+1, p.ID, p.Latitude, p.Longitude)
	}
	
}