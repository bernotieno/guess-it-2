package calculation


func calculateVariance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += (v - mean) * (v - mean)
	}
	return sum / float64(len(data))
}