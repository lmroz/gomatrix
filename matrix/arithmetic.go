// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math"
	//"reflect"
)

/*
Finds the sum of two matrices.
*/
func Sum(A, B MatrixRO) *DenseMatrix {
	C := MakeDenseCopy(A)
	err := C.Add(MakeDenseCopy(B))
	if err.OK() {
		return C
	}
	return nil
}

/*
Finds the difference between two matrices.
*/
func Difference(A, B MatrixRO) *DenseMatrix {
	C := MakeDenseCopy(A)
	err := C.Subtract(MakeDenseCopy(B))
	if err.OK() {
		return C
	}
	return nil
}

/*
Finds the Product of two matrices.
*/
func Product(A, B MatrixRO) *DenseMatrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	C := Zeros(A.Rows(), B.Cols())

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0)
			for k := 0; k < A.Cols(); k++ {
				sum += A.Get(i, k) * B.Get(k, j)
			}
			C.Set(i, j, sum)
		}
	}

	return C
}

/*
Uses a number of goroutines to do the dot products necessary
for the matrix multiplication in parallel.
*/
func ParallelProduct(A, B MatrixRO) *DenseMatrix {
	if A.Cols() != B.Rows() {
		return nil
	}

	C := Zeros(A.Rows(), B.Cols())

	in := make(chan int)
	quit := make(chan bool)

	dotRowCol := func() {
		for {
			select {
			case i := <-in:
				sums := make([]float64, B.Cols())
				for k := 0; k < A.Cols(); k++ {
					for j := 0; j < B.Cols(); j++ {
						sums[j] += A.Get(i, k) * B.Get(k, j)
					}
				}
				for j := 0; j < B.Cols(); j++ {
					C.Set(i, j, sums[j])
				}
			case <-quit:
				return
			}
		}
	}

	threads := 2

	for i := 0; i < threads; i++ {
		go dotRowCol()
	}

	for i := 0; i < A.Rows(); i++ {
		in <- i
	}

	for i := 0; i < threads; i++ {
		quit <- true
	}

	return C
}

/*
Scales a matrix by a scalar.
*/
func Scaled(A MatrixRO, f float64) *DenseMatrix {
	B := MakeDenseCopy(A)
	B.Scale(f)
	return B
}

/*
Tests the element-wise equality of the two matrices.
*/
func Equals(A, B MatrixRO) bool {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return false
	}
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j) != B.Get(i, j) {
				return false
			}
		}
	}
	return true
}

/*
Tests to see if the difference between two matrices,
element-wise, exceeds ε.
*/
func ApproxEquals(A, B MatrixRO, ε float64) bool {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return false
	}
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if math.Fabs(A.Get(i, j)-B.Get(i, j)) > ε {
				return false
			}
		}
	}
	return true
}

/*
Finds the product of any number of matrices.
*/
/*
//this stopped compiling
func MultipleProduct(values ...) Matrix {
	v := reflect.NewValue(values).(*reflect.StructValue)
	if v.NumField() < 2 {
		return nil
	}

	inter := v.Field(0).Interface()
	B, ok := inter.(MatrixRO)
	if ok {
		C := MakeDenseCopy(B)
		for i := 1; i < v.NumField(); i++ {
			inter := v.Field(i).Interface()
			if A, ok := inter.(MatrixRO); ok {
				C = Product(C, A)
			} else {
				return nil
			}
		}
		return C
	}

	return nil
}
*/
