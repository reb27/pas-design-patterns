package pattern

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
		return result
	}

	result := p.realCalculator.CalculateROI(client)

	if p.cache == nil {
		p.cache = make(map[uuid.UUID]float64)
	}
	p.cache[client.ID] = result

	return result
}

func Run() {
	var totalTime time.Duration

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
		roi := proxy.CalculateROI(client)
		duration := time.Since(start)
		fmt.Printf("ROI calculado: %.2f%%\n", roi)
		fmt.Printf("Tempo para calcular o ROI pela %d vez: %v\n", i, duration)
		totalTime += duration
	}

	fmt.Printf("[COM CACHE PROXY] O tempo total para executar duas requests para um client foi de : %d\n", totalTime)
	fmt.Printf("_____________________________________________________________________________________________\n")
}
