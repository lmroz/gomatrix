package matrix

import (
	"testing"
	"fmt"
	"time"
)

func BenchmarkTransposeTimes(b *testing.B) {
	for s := 25; s<=100; s+=25 {
		w, h := s/2, s*2
	
		A := Normals(h, w)
		B := Normals(w, h)
		
		var times [2]float64
		
		const Count = 500
		
		MaxProcs = 1
		WhichSyncMethod = 1
		start := time.Nanoseconds()
		for i := 0; i < Count; i++ {
			A.Times(B)
		}
		end := time.Nanoseconds()
		duration := end - start
		times[0] = float64(duration) / 1e9
		
		WhichSyncMethod = 2
		start = time.Nanoseconds()
		for i := 0; i < Count; i++ {
			A.Times(B)
		}
		end = time.Nanoseconds()
		duration = end - start
		times[1] = float64(duration) / 1e9
		fmt.Printf("%d: %.2f\n", h, times[1]/times[0])
	}
}