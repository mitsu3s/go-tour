package functions

func Exercise(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
