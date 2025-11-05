// Package aggregator provides aggregation helpers for energy data.
package aggregator

import "time"

// Point represents a timestamped numeric value.
type Point struct {
	Value     float64
	Timestamp time.Time
}

// Sum returns the sum of all values.
func Sum(points []Point) float64 {
	s := 0.0
	for _, p := range points {
		s += p.Value
	}
	return s
}

// Average returns the arithmetic mean. If empty, returns 0.
func Average(points []Point) float64 {
	if len(points) == 0 {
		return 0
	}
	return Sum(points) / float64(len(points))
}

// GroupByDay groups points by YYYY-MM-DD and returns totals per day in chronological order.
func GroupByDay(points []Point) map[string]float64 {
	out := make(map[string]float64)
	for _, p := range points {
		day := p.Timestamp.Format("2006-01-02")
		out[day] += p.Value
	}
	return out
}

// MovingAverage computes a simple moving average over a specified window size.
// If windowSize <= 0, returns empty slice.
func MovingAverage(points []Point, windowSize int) []float64 {
	if windowSize <= 0 || len(points) < windowSize {
		return []float64{}
	}
	res := make([]float64, 0, len(points)-windowSize+1)
	windowSum := 0.0
	for i := 0; i < len(points); i++ {
		windowSum += points[i].Value
		if i >= windowSize {
			windowSum -= points[i-windowSize].Value
		}
		if i >= windowSize-1 {
			res = append(res, windowSum/float64(windowSize))
		}
	}
	return res
}
