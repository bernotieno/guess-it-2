package calculation

// CalculateRange calculates the predicted range based on the latest data.
func CalculateRange(data []float64) (float64, float64) {
	mean := calculateMean(data)
	variance := calculateVariance(data, mean)
	stddev := calculateStandardDeviation(variance)
	slope, intercept := linearRegression(data)
	nextIndex := float64(len(data))
	predictedValue := slope*nextIndex + intercept
	lower := predictedValue - 2.576*stddev
	upper := predictedValue + 2.576*stddev
	return lower, upper
}
