package pure

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

func (r *ROICalculator) CalculateROI(client Client) float64 {
	// simular demora no cálculo
	time.Sleep(10 * time.Second)
	return ((client.Revenue - client.Investment) / client.Investment) * 100
}

func Run() {
	var totalTime time.Duration

	client := Client{
		ID:         uuid.New(),
		Revenue:    5000,
		Investment: 2000,
	}

	realCalculator := &ROICalculator{}

	for i := 1; i <= 2; i++ {
		start := time.Now()
		roi := realCalculator.CalculateROI(client)
		duration := time.Since(start)
		fmt.Printf("ROI calculado: %.2f%%\n", roi)
		fmt.Printf("Tempo para calcular o ROI pela %d vez: %v\n", i, duration)
		totalTime += duration
	}

	fmt.Printf("[SEM CACHE] O tempo total para executar duas requests para um client foi de : %d\n", totalTime)
	fmt.Printf("_____________________________________________________________________________________________\n")

}
