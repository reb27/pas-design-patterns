package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ROICalculatorInterface interface {
	CalculateROI(client Client) float64
}

type Client struct {
	ID         uuid.UUID
	Revenue    float64
	Investment float64
}

type ROICalculator struct{}

type ROICalculatorCacheProxy struct {
	realCalculator ROICalculatorInterface
	cache          map[uuid.UUID]float64
}

func (r *ROICalculator) CalculateROI(client Client) float64 {
	// simular demora no cálculo
	time.Sleep(10 * time.Second)
	return ((client.Revenue - client.Investment) / client.Investment) * 100
}

func (p *ROICalculatorCacheProxy) CalculateROI(client Client) float64 {
	if result, exists := p.cache[client.ID]; exists {
		fmt.Println("Resultado cacheado encontrado")
		return result
	}

	result := p.realCalculator.CalculateROI(client)

	if p.cache == nil {
		p.cache = make(map[uuid.UUID]float64)
	}
	p.cache[client.ID] = result
	fmt.Println("Resultado calculado e armazenado no cache")

	return result
}

func main() {
	client := Client{
		ID:         uuid.New(),
		Revenue:    5000,
		Investment: 2000,
	}

	realCalculator := &ROICalculator{}
	proxy := &ROICalculatorCacheProxy{realCalculator: realCalculator}

	for i := 1; i <= 2; i++ {
		// a segunda vez é o retorno do resultado cacheado
		start := time.Now()

		roi1 := proxy.CalculateROI(client)
		duration := time.Since(start)
		fmt.Printf("ROI calculado: %.2f%%\n", roi1)
		fmt.Printf("Tempo para calcular o ROI vez %d: %v\n", i, duration)
	}
}
