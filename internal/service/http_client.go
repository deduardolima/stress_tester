package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/deduardolima/stress_tester/internal/entity"
)

const maxRetries = 3
const delayBetweenRetries = 1 * time.Second

func PerformRequests(url string, totalRequests int, concurrency int) entity.Report {
	var wg sync.WaitGroup
	var mu sync.Mutex
	report := entity.Report{
		TotalRequests:      totalRequests,
		StatusDistribution: make(map[int]int),
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			for attempt := 0; attempt < maxRetries; attempt++ {
				resp, err := client.Get(url)
				if err != nil {
					fmt.Printf("Error making request (attempt %d): %v\n", attempt+1, err)
					if attempt == maxRetries-1 {
						mu.Lock()
						report.StatusDistribution[0]++
						mu.Unlock()
					}
				} else {
					mu.Lock()
					report.StatusDistribution[resp.StatusCode]++
					if resp.StatusCode == 200 {
						report.SuccessfulRequests++
					}
					mu.Unlock()
					resp.Body.Close()
					break
				}
				time.Sleep(delayBetweenRetries)
			}
			<-semaphore
		}()
	}

	wg.Wait()
	return report
}
