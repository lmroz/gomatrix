// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y;
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y;
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y;
}

func minInt(x, y int) int {
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
