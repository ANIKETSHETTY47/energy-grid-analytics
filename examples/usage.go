package main

import (
	"fmt"
	"time"

	"github.com/ANIKETSHETTY47/energy-grid-analytics-go/aggregator"
	"github.com/ANIKETSHETTY47/energy-grid-analytics-go/anomaly"
	"github.com/ANIKETSHETTY47/energy-grid-analytics-go/converter"
	"github.com/ANIKETSHETTY47/energy-grid-analytics-go/maintenance"
)

func main() {
	// Converter demo
	conv := &converter.EnergyConverter{}
	fmt.Println("5,000 kWh -> MWh:", conv.KWhToMWh(5000))
	fmt.Println("Cost (peak, 100 kWh @ €0.20):", conv.CalculateCost(100, 0.20, "peak"))
	fmt.Println("Efficiency (input 120, output 90):", conv.CalculateEfficiency(120, 90))

	// Aggregator demo
	points := []aggregator.Point{
		{Value: 10, Timestamp: time.Now().Add(-72 * time.Hour)},
		{Value: 15, Timestamp: time.Now().Add(-48 * time.Hour)},
		{Value: 20, Timestamp: time.Now().Add(-24 * time.Hour)},
		{Value: 25, Timestamp: time.Now()},
	}
	fmt.Println("Sum:", aggregator.Sum(points))
	fmt.Println("Average:", aggregator.Average(points))
	fmt.Println("MovingAverage(2):", aggregator.MovingAverage(points, 2))

	// Anomaly demo
	readings := []anomaly.Reading{
		{Consumption: 10, Timestamp: 1},
		{Consumption: 11, Timestamp: 2},
		{Consumption: 10, Timestamp: 3},
		{Consumption: 200, Timestamp: 4}, // spike
		{Consumption: 12, Timestamp: 5},
		{Consumption: 13, Timestamp: 6},
	}
	detector := &anomaly.AnomalyDetector{Threshold: 2.0, WindowSize: 3}
	fmt.Println("Spikes:", detector.DetectSpikes(readings))
	fmt.Println("Outliers:", detector.DetectOutliers(readings))

	// Maintenance demo
	h := maintenance.AssetHealth{
		HoursRun:           2500,
		FailureRatePerYear: 0.3,
		LastService:        time.Now().Add(-180 * 24 * time.Hour),
		ServiceInterval:    365 * 24 * time.Hour,
	}
	fmt.Println("Failure risk in 90 days:", maintenance.FailureRisk(h.FailureRatePerYear, 90*24*time.Hour))
	fmt.Println("Next service date:", maintenance.NextServiceDate(h))
}
