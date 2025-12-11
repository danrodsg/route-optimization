# 🗺️ Route Optimization API em GoLang

Este projeto implementa um otimizador de rotas utilizando o algoritmo heurístico **Nearest Neighbor** (Vizinho Mais Próximo) em Go, ideal para resolver o Problema do Caixeiro Viajante (TSP) para um conjunto de pontos.

[![Go](https://img.shields.io/badge/Golang-v1.21+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

---

## 💻 Tecnologias e Arquitetura

| Componente | Pacote Go | Função no Projeto |
| :--- | :--- | :--- |
| **Ponto e Distância** | `optimizer/point.go` | Define a estrutura `Point` (ID, Lat, Lon) e a interface `DistanceCalculator`. |
| **Calculadora** | `optimizer/haversine.go` | Implementação da métrica de **Distância Euclidiana** (fácil de estender para Haversine). |
| **Otimizador** | `optimizer/optimizer.go` | Contém a lógica principal do algoritmo **Nearest Neighbor**. |
| **Execução** | `cmd/main.go` | Configura o `deposito` e os `campos` e executa a otimização. |

---

## ✨ Funcionalidades (Nearest Neighbor)

O algoritmo **Vizinho Mais Próximo** é uma heurística greedy que constrói a rota passo a passo:

1.  **Ponto de Partida:** Começa sempre no `DEPOT`.
2.  **Iteração:** No ponto atual, ele calcula a distância para **todos** os outros pontos ainda não visitados.
3.  **Seleção:** O próximo ponto na rota é o ponto **mais próximo** (com a menor distância) do ponto atual.
4.  **Conclusão:** O processo se repete até que todos os pontos tenham sido visitados.

### 📐 Cálculo de Distância

A implementação padrão utiliza a **Distância Euclidiana** para simplificar o exemplo, baseada nas coordenadas (Latitude, Longitude):

$$
d = \sqrt{(\text{Lat}_2 - \text{Lat}_1)^2 + (\text{Lon}_2 - \text{Lon}_1)^2}
$$

A interface `DistanceCalculator` em `point.go` facilita a troca por métodos mais precisos para coordenadas geográficas, como a **Distância de Haversine**.

---

## 🚀 Como Executar

### Pré-requisitos
* Go (versão 1.21 ou superior).

### 1. Clonar o Repositório
```bash
git clone [seu-link-do-repositório]
cd route-optimization
````
### 2. Executar

Execute o arquivo principal dentro do pacote cmd:

```
bash
go run ./cmd/main.go
