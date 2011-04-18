package main

import (
	"fmt"
	"time"
	"runtime"
	. "gomatrix.googlecode.com/hg/matrix"
)

func main() {
	for w := 0; w <= 200; w += 200 {
		for h := 0; h <= 200; h += 200 {
			if w == 0 && h != 0 {
				fmt.Printf("%d\t", h)
			}
			if w != 0 && h == 0 {
				fmt.Printf("%d\t", w)
			}
			if w == 0 && h == 0 {
				fmt.Printf("w/h\t")
			}
			if w == 0 || h == 0 {
				continue
			}
			A := Normals(h, w)
			B := Normals(w, h)
			times := []float64{0, 0, 0}
			
			const Count = 100
			
			MaxProcs := runtime.GOMAXPROCS(0)
			runtime.GOMAXPROCS(1)
			start := time.Nanoseconds()
			for i := 0; i < Count; i++ {
				A.Times(B)
			}
			end := time.Nanoseconds()
			duration := end - start
			times[0] = float64(duration) / 1e9
			
			for WhichParMethod=1; WhichParMethod<3; WhichParMethod++ {
				runtime.GOMAXPROCS(MaxProcs)
				start = time.Nanoseconds()
				for i := 0; i < Count; i++ {
					A.Times(B)
				}
				end = time.Nanoseconds()
				duration = end - start
				times[WhichParMethod] = float64(duration) / 1e9	
				
			}
			ratio1 := times[1] / times[0]
			ratio2 := times[2] / times[0]
			fmt.Printf("%.2f:%.2f\t", ratio1, ratio2)
		}
		fmt.Printf("\n")
	}
}
