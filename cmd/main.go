package main

import (
    "log"
   
    "github.com/danrodsg/route-optimization/api"
    "github.com/gin-gonic/gin"
)

func main() {
    
    router := gin.Default()

    router.POST("/api/v1/optimize", api.OptimizeRouteHandler)

    log.Println("Servidor de Otimização de Rotas rodando na porta 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}