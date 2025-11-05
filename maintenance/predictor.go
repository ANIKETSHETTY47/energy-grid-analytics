// Package maintenance provides simple predictive helpers for maintenance scheduling.
package maintenance

import (
	"math"
	"time"
)

// AssetHealth captures minimal telemetry for a device/asset.
type AssetHealth struct {
	HoursRun           float64   // cumulative run-hours
	FailureRatePerYear float64   // lambda in failures/year for an exponential model
	LastService        time.Time // last service date
	ServiceInterval    time.Duration
}

// FailureRisk estimates probability of at least one failure within horizon using an exponential model.
func FailureRisk(lambdaPerYear float64, horizon time.Duration) float64 {
	if lambdaPerYear <= 0 || horizon <= 0 {
		return 0
	}
	years := horizon.Hours() / (24 * 365)
	return 1 - math.Exp(-lambdaPerYear*years)
}

// NextServiceDate returns the recommended next service date based on fixed interval and usage bias.
func NextServiceDate(h AssetHealth) time.Time {
	if h.ServiceInterval <= 0 {
		return h.LastService
	}
	// Simple usage-based bias: shorten interval by up to 20% if hours are high.
	bias := 1.0
	if h.HoursRun > 4000 {
		bias = 0.8
	} else if h.HoursRun > 2000 {
		bias = 0.9
	}
	adj := time.Duration(float64(h.ServiceInterval) * bias)
	return h.LastService.Add(adj)
}
