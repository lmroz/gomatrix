package main

import (
	"fmt"
	"time"
	"os"
	"runtime"
	. "gomatrix.googlecode.com/hg/matrix"
)

func timeCall(A, B *DenseMatrix, foo func(A, B *DenseMatrix) (*DenseMatrix, os.Error)) {
	foo(A, B)
}

func main() {
	var start, end, duration int64
	var count int

	var A, B *DenseMatrix

	w := 500
	h := 500
	A = Normals(h, w)
	B = Normals(w, h)

	count = 2

	for MaxProcs := 1; MaxProcs < 3; MaxProcs++ {
		runtime.GOMAXPROCS(MaxProcs)
		fmt.Printf("With MaxProcs=%d:\n", MaxProcs)
		start = time.Nanoseconds()
		for i := 0; i < count; i++ {
			A.Times(B)
		}
		end = time.Nanoseconds()
		duration = end - start
		fmt.Printf("%d %dx%d x %dx%d matrix multiplications in %fs\n", count, h, w, w, h, float64(end-start)/1000000000)

		start = time.Nanoseconds()
		for i := 0; i < count; i++ {
			A.TimesDense(B)
		}
		end = time.Nanoseconds()
		fmt.Printf("%d %dx%d x %dx%d dense matrix multiplications in %fs\n", count, h, w, w, h, float64(end-start)/1000000000)
		fmt.Printf("For a ratio of %f.\n", float64(duration)/float64(end-start))
	}

	if true {
		return
	}

	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.Plus(A)
	}
	end = time.Nanoseconds()
	duration = end - start
	fmt.Printf("%d %dx%d x %dx%d matrix additions in %fs\n", count, h, w, w, h, float64(end-start)/1000000000)

	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.PlusDense(A)
	}
	end = time.Nanoseconds()
	fmt.Printf("%d %dx%d x %dx%d dense matrix additions in %fs\n", count, h, w, w, h, float64(end-start)/1000000000)
	fmt.Printf("For a ratio of %f.\n", float64(duration)/float64(end-start))

	A = Normals(4, 4)
	B = Normals(4, 4)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.ElementMult(B)
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 4x4 matrix element multiplications in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(4, 4)
	B = Normals(4, 4)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.Plus(B)
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 4x4 matrix additions in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(6, 6)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.Inverse()
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 6x6 matrix inversions in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(6, 6)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.Det()
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 6x6 matrix determinants in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(6, 6)
	Bm, _ := A.Times(A.Transpose())
	B, _ = Bm.(*DenseMatrix)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		B.Cholesky()
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 6x6 cholesky decompositions in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(6, 6)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.QR()
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 6x6 QR decompositions in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(6, 6)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.Eigen()
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 6x6 Eigenvector decompositions in %fs\n", count, float64(end-start)/1000000000)

	A = Normals(6, 6)

	count = 100000
	start = time.Nanoseconds()
	for i := 0; i < count; i++ {
		A.SVD()
	}
	end = time.Nanoseconds()
	fmt.Printf("%d 6x6 singular value decompositions in %fs\n", count, float64(end-start)/1000000000)

}
