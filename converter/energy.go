// Package converter provides utilities for energy calculations and cost/efficiency helpers.
package converter

// EnergyConverter provides utilities for energy calculations.
type EnergyConverter struct{}

// KWhToMWh converts kilowatt-hours to megawatt-hours.
func (ec *EnergyConverter) KWhToMWh(kwh float64) float64 {
	return kwh / 1000.0
}

// CalculateCost computes cost with optional tier multipliers ("peak", "offpeak", or default).
func (ec *EnergyConverter) CalculateCost(kwh float64, rate float64, tier string) float64 {
	multiplier := 1.0

	switch tier {
	case "peak":
		multiplier = 1.5
	case "offpeak":
		multiplier = 0.7
	default:
		multiplier = 1.0
	}

	return kwh * rate * multiplier
}

// CalculateEfficiency returns output/input as a percentage (0..100). If input is 0, returns 0.
func (ec *EnergyConverter) CalculateEfficiency(inputKWh, outputKWh float64) float64 {
	if inputKWh == 0 {
		return 0
	}
	return (outputKWh / inputKWh) * 100.0
}
