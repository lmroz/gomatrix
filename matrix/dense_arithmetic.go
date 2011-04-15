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

func parTimes1(A, B *DenseMatrix) (C *DenseMatrix) {
	C = Zeros(A.rows, B.cols)
	wait := parFor(countBoxes(0, A.rows), func(iBox box) {
		i := iBox.(int)
		sums := C.elements[i*C.step : (i+1)*C.step]
		for k := 0; k < A.cols; k++ {
			for j := 0; j < B.cols; j++ {
				sums[j] += A.elements[i*A.step+k] * B.elements[k*B.step+j]
			}
		}
	})
	wait()

	return
}

//this is an adaptation of code from a go-nuts post made by Dmitriy Vyukov
func parTimes2(A, B *DenseMatrix) (C *DenseMatrix) {
	C = Zeros(A.rows, B.cols)

	const threshold = 8

	currentGoroutineCount := 1
	maxGoroutines := MaxProcs+2

	var aux func(sync chan bool, A, B, C *DenseMatrix, rs, re, cs, ce, ks, ke int)
	aux = func(sync chan bool, A, B, C *DenseMatrix, rs, re, cs, ce, ks, ke int) {
		switch {
		case currentGoroutineCount < maxGoroutines && re-rs >= threshold:
			sync0 := make(chan bool, 1)
			rm := (rs + re) / 2
			currentGoroutineCount++
			go aux(sync0, A, B, C, rs, rm, cs, ce, ks, ke)
			aux(nil, A, B, C, rm, re, cs, ce, ks, ke)
			<-sync0
			currentGoroutineCount--
		case currentGoroutineCount < maxGoroutines && ce-cs >= threshold:
			sync0 := make(chan bool, 1)
			cm := (cs + ce) / 2
			currentGoroutineCount++
			go aux(sync0, A, B, C, rs, re, cs, cm, ks, ke)
			aux(nil, A, B, C, rs, re, cm, ce, ks, ke)
			<-sync0
			currentGoroutineCount--
		case currentGoroutineCount < maxGoroutines && ke-ks >= threshold:
			km := (ks + ke) / 2
			//why don't we split here, too?
			//one answer - at this point we've already got way more goroutines than procs
			aux(nil, A, B, C, rs, re, cs, ce, ks, km)
			aux(nil, A, B, C, rs, re, cs, ce, km, ke)
		default:
			for row := rs; row < re; row++ {
				sums := C.elements[row*C.step : (row+1)*C.step]
				for k := ks; k < ke; k++ {
					for col := cs; col < ce; col++ {
						sums[col] += A.elements[row*A.step+k] * B.elements[k*B.step+col]
					}
				}
			}
		}
		if sync != nil {
			sync <- true
		}
	}

	aux(nil, A, B, C, 0, A.rows, 0, B.cols, 0, A.cols)

	return
}

var (
	WhichParMethod = 2
	WhichSyncMethod = 1
)

func (A *DenseMatrix) TimesDense(B *DenseMatrix) (*DenseMatrix, os.Error) {
	if A.cols != B.rows {
		return nil, ErrorDimensionMismatch
	}
	var C *DenseMatrix
	if MaxProcs > 1 {
		switch WhichParMethod {
		case 1:
			C = parTimes1(A, B)
		case 2:
			C = parTimes2(A, B)
		}
	} else {
		switch {
		case A.cols > 100 && WhichSyncMethod == 2:
			C = transposeTimes(A, B)
		default:
			C = Zeros(A.rows, B.cols)
			for i := 0; i < A.rows; i++ {
				sums := C.elements[i*C.step : (i+1)*C.step]
				for k := 0; k < A.cols; k++ {
					for j := 0; j < B.cols; j++ {
						sums[j] += A.elements[i*A.step+k] * B.elements[k*B.step+j]
					}
				}
			}
		}
	}

	return C, nil
}

func transposeTimes(A, B *DenseMatrix) (C *DenseMatrix) {
	Bt := B.Transpose()
	C = Zeros(A.rows, B.cols)
	
	Bcols := Bt.Arrays()
	
	for i := 0; i < A.rows; i++ {
		Arow := A.elements[i*A.step : i*A.step+A.cols]
		for j := 0; j < B.cols; j++ {
			Bcol := Bcols[j]
			for k := range Arow {
				C.elements[i*C.step+j] += Arow[k]*Bcol[k]
			} 
		}
	}
	
	return
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
