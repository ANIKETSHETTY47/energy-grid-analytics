// Package anomaly provides basic analytics to detect spikes and outliers in consumption readings.
package anomaly

import (
	"math"
	"sort"
)

// Reading is a simple time-series point containing consumption and an epoch timestamp.
type Reading struct {
	Consumption float64
	Timestamp   int64
}

// AnomalyDetector holds configuration for detection.
type AnomalyDetector struct {
	Threshold  float64 // number of standard deviations for spike detection
	WindowSize int     // sliding window size
}

// DetectSpikes flags points whose consumption deviates from the rolling mean by Threshold*stddev.
func (ad *AnomalyDetector) DetectSpikes(readings []Reading) []Reading {
	if ad.WindowSize <= 1 || len(readings) < ad.WindowSize {
		return []Reading{}
	}

	var anomalies []Reading
	for i := ad.WindowSize; i < len(readings); i++ {
		window := readings[i-ad.WindowSize : i]
		avg := ad.calculateAverage(window)
		stdDev := ad.calculateStdDev(window, avg)

		current := readings[i].Consumption
		if stdDev == 0 {
			continue
		}

		if math.Abs(current-avg) > ad.Threshold*stdDev {
			anomalies = append(anomalies, readings[i])
		}
	}
	return anomalies
}

// DetectOutliers uses IQR (Tukey) to return points outside [Q1-1.5*IQR, Q3+1.5*IQR].
func (ad *AnomalyDetector) DetectOutliers(readings []Reading) []Reading {
	if len(readings) < 4 {
		return []Reading{}
	}

	values := make([]float64, len(readings))
	for i, r := range readings {
		values[i] = r.Consumption
	}
	sort.Float64s(values)

	q1 := ad.percentile(values, 25)
	q3 := ad.percentile(values, 75)
	iqr := q3 - q1

	lowerBound := q1 - 1.5*iqr
	upperBound := q3 + 1.5*iqr

	var outliers []Reading
	for _, r := range readings {
		if r.Consumption < lowerBound || r.Consumption > upperBound {
			outliers = append(outliers, r)
		}
	}
	return outliers
}

func (ad *AnomalyDetector) calculateAverage(readings []Reading) float64 {
	sum := 0.0
	for _, r := range readings {
		sum += r.Consumption
	}
	return sum / float64(len(readings))
}

func (ad *AnomalyDetector) calculateStdDev(readings []Reading, mean float64) float64 {
	variance := 0.0
	for _, r := range readings {
		variance += math.Pow(r.Consumption-mean, 2)
	}
	variance /= float64(len(readings))
	return math.Sqrt(variance)
}

func (ad *AnomalyDetector) percentile(sortedValues []float64, p float64) float64 {
	if len(sortedValues) == 0 {
		return 0
	}
	index := (p / 100.0) * float64(len(sortedValues)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	if lower == upper {
		return sortedValues[lower]
	}
	return sortedValues[lower] + (sortedValues[upper]-sortedValues[lower])*(index-float64(lower))
}
