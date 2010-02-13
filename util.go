// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "sync"

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sum(a []float64) (s float64) {
	for _, v := range a {
		s += v
	}
	return
}

func product(a []float64) float64 {
	p := float64(1)
	for _, v := range a {
		p *= v
	}
	return p
}

type box interface{}

func countBoxes(start, cap int) chan box {
	ints := make(chan box)
	go func() {
		for i := start; i < cap; i++ {
			ints <- i
		}
		close(ints)
	}()
	return ints
}


func parFor(inputs <-chan box, foo func(i box)) (wait func()) {
	m := new(sync.Mutex)
	block := make(chan bool, MaxThreads)
	for j := 0; j < MaxThreads; j++ {
		go func() {
			for {
				m.Lock()
				i, done := <-inputs, closed(inputs)
				m.Unlock()
				if done {
					break
				}
				foo(i)
			}
			block <- true
		}()
	}
	wait = func() {
		for i := 0; i < MaxThreads; i++ {
			<-block
		}
	}
	return
}
