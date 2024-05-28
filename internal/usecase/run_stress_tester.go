package usecase

import (
	"fmt"
	"time"

	"github.com/deduardolima/stress_tester/internal/service"
)

func RunStressTest(url string, totalRequests int, concurrency int) {
	start := time.Now()
	report := service.PerformRequests(url, totalRequests, concurrency)
	duration := time.Since(start).Seconds()
	report.TotalTime = duration

	fmt.Printf("Total Time: %.2f seconds\n", report.TotalTime)
	fmt.Printf("Total Requests: %d\n", report.TotalRequests)
	fmt.Printf("Successful Requests: %d\n", report.SuccessfulRequests)
	fmt.Printf("Status Distribution: %v\n", report.StatusDistribution)
}
