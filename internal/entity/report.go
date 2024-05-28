package entity

type Report struct {
	TotalTime          float64
	TotalRequests      int
	SuccessfulRequests int
	StatusDistribution map[int]int
}
