package main

import (
	"fmt"
	"time"
	. "gomatrix.googlecode.com/hg/matrix"
)

func main() {
	for w := 0; w <= 600; w += 100 {
		for h := 0; h <= 600; h += 100 {
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
				for i := 0; i < 50; i++ {
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
