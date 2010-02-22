package main

import (
	. "matrix"
	"fmt"
	"time"
)

func main() {
	for w := 0; w <= 400; w += 50 {
		for h := 0; h <= 400; h += 50 {
			if w == 0 && h != 0 {
				fmt.Printf("%d,", h)
			}
			if w != 0 && h == 0 {
				fmt.Printf("%d,", w)
			}
			if w == 0 && h == 0 {
				fmt.Printf("w/h,")
			}
			if w == 0 || h == 0 {
				continue
			}
			A := Normals(h, w)
			B := Normals(w, h)
			times := []float{0, 0}
			for MaxProcs = 1; MaxProcs < len(times)+1; MaxProcs++ {
				start := time.Nanoseconds()
				for i := 0; i < 10; i++ {
					A.Times(B)
				}
				end := time.Nanoseconds()
				duration := end - start
				times[MaxProcs-1] = float(duration)/1000000000
			}
			ratio := times[1]/times[0]
			fmt.Printf("%f,", ratio)
		}
		fmt.Printf("\n")
	}
}
