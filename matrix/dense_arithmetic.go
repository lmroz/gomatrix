// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"os"
)

func (A *DenseMatrix) Plus(B MatrixRO) (Matrix, os.Error) {
	C := A.Copy()
	err := C.Add(B)
	return C, err
}
func (A *DenseMatrix) PlusDense(B *DenseMatrix) (*DenseMatrix, os.Error) {
	C := A.Copy()
	err := C.AddDense(B)
	return C, err
}

func (A *DenseMatrix) Minus(B MatrixRO) (Matrix, os.Error) {
	C := A.Copy()
	err := C.Subtract(B)
	return C, err
}

func (A *DenseMatrix) MinusDense(B *DenseMatrix) (*DenseMatrix, os.Error) {
	C := A.Copy()
	err := C.SubtractDense(B)
	return C, err
}

func (A *DenseMatrix) Add(B MatrixRO) os.Error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		index := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[index] += B.Get(i, j)
			index++
		}
	}

	return nil
}

func (A *DenseMatrix) AddDense(B *DenseMatrix) os.Error {
	if A.cols != B.cols || A.rows != B.rows {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.elements[i*A.step+j] += B.elements[i*B.step+j]
		}
	}

	return nil
}

func (A *DenseMatrix) Subtract(B MatrixRO) os.Error {
	if Bd, ok := B.(*DenseMatrix); ok {
		return A.SubtractDense(Bd)
	}

	if A.cols != B.Cols() || A.rows != B.Rows() {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		index := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[index] -= B.Get(i, j)
			index++
		}
	}

	return nil
}

func (A *DenseMatrix) SubtractDense(B *DenseMatrix) os.Error {

	if A.cols != B.cols || A.rows != B.rows {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		indexA := i * A.step
		indexB := i * B.step

		for j := 0; j < A.cols; j++ {
			A.elements[indexA] -= B.elements[indexB]
			indexA++
			indexB++
		}
	}

	return nil
}

func (A *DenseMatrix) Times(B MatrixRO) (Matrix, os.Error) {

	if Bd, ok := B.(*DenseMatrix); ok {
		return A.TimesDense(Bd)
	}

	if A.cols != B.Rows() {
		return nil, ErrorDimensionMismatch
	}
	C := Zeros(A.rows, B.Cols())

	for i := 0; i < A.rows; i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.elements[i*A.step+k] * B.Get(k, j)
			}
			C.elements[i*C.step+j] = sum
		}
	}

	return C, nil
}

func (A *DenseMatrix) TimesDense(B *DenseMatrix) (*DenseMatrix, os.Error) {
	if A.cols != B.rows {
		return nil, ErrorDimensionMismatch
	}
	C := Zeros(A.rows, B.cols)
	///*
	if MaxProcs > 1 {
		wait := parFor(countBoxes(0, A.rows), func(iBox box) {
			i := iBox.(int)
			sums := C.elements[i*C.step : (i+1)*C.step]
			for k := 0; k < A.Cols(); k++ {
				for j := 0; j < B.Cols(); j++ {
					sums[j] += A.elements[i*A.step+k] * B.elements[k*B.step+j]
				}
			}
		})

		wait()
	} else {
		for i := 0; i < A.rows; i++ {
			for j := 0; j < B.cols; j++ {
				sum := float64(0);
				for k := 0; k < A.cols; k++ {
					sum += A.elements[i*A.step+k] * B.elements[k*B.step+j]
				}
				C.elements[i*C.step+j] = sum;
			}
		}
	}

	return C, nil
}


func (A *DenseMatrix) ElementMult(B MatrixRO) (Matrix, os.Error) {
	C := A.Copy()
	err := C.ScaleMatrix(B)
	return C, err
}

func (A *DenseMatrix) ElementMultDense(B *DenseMatrix) (*DenseMatrix, os.Error) {
	C := A.Copy()
	err := C.ScaleMatrixDense(B)
	return C, err
}

func (A *DenseMatrix) Scale(f float64) {
	for i := 0; i < A.rows; i++ {
		index := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[index] *= f
			index++
		}
	}
}

func (A *DenseMatrix) ScaleMatrix(B MatrixRO) os.Error {
	if Bd, ok := B.(*DenseMatrix); ok {
		return A.ScaleMatrixDense(Bd)
	}

	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}
	for i := 0; i < A.rows; i++ {
		indexA := i * A.step
		for j := 0; j < A.cols; j++ {
			A.elements[indexA] *= B.Get(i, j)
			indexA++
		}
	}
	return nil
}

func (A *DenseMatrix) ScaleMatrixDense(B *DenseMatrix) os.Error {
	if A.rows != B.rows || A.cols != B.cols {
		return ErrorDimensionMismatch
	}
	for i := 0; i < A.rows; i++ {
		indexA := i * A.step
		indexB := i * B.step
		for j := 0; j < A.cols; j++ {
			A.elements[indexA] *= B.elements[indexB]
			indexA++
			indexB++
		}
	}
	return nil
}
