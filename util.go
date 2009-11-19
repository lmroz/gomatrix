package matrix

func max(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y;
}

func min(x float64, y float64) float64 {
	if x < y {
		return x
	}
	return y;
}

func sum(a []float64) (s float64) {
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return;
}

func product(a []float64) float64 {
	p := float64(1);
	for i := 0; i < len(a); i++ {
		p *= a[i]
	}
	return p;
}

