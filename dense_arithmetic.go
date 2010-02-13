// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

func (A *DenseMatrix) Plus(B MatrixRO) (Matrix, Error) {
	C := A.Copy()
	err := C.Add(B)
	return C, err
}
func (A *DenseMatrix) PlusDense(B *DenseMatrix) (*DenseMatrix, Error) {
	C := A.Copy()
	err := C.AddDense(B)
	return C, err
}

func (A *DenseMatrix) Minus(B MatrixRO) (Matrix, Error) {
	C := A.Copy()
	err := C.Subtract(B)
	return C, err
}

func (A *DenseMatrix) MinusDense(B *DenseMatrix) (*DenseMatrix, Error) {
	C := A.Copy()
	err := C.SubtractDense(B)
	return C, err
}

func (A *DenseMatrix) Add(B MatrixRO) Error {
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

	return NoError
}

func (A *DenseMatrix) AddDense(B *DenseMatrix) Error {
	if A.cols != B.cols || A.rows != B.rows {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.elements[i*A.step+j] += B.elements[i*B.step+j]
		}
	}

	return NoError
}

func (A *DenseMatrix) Subtract(B MatrixRO) Error {
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

	return NoError
}

func (A *DenseMatrix) SubtractDense(B *DenseMatrix) Error {

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

	return NoError
}

func (A *DenseMatrix) Times(B MatrixRO) (Matrix, Error) {

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

	return C, NoError
}

func (A *DenseMatrix) TimesDense(B *DenseMatrix) (*DenseMatrix, Error) {
	if A.cols != B.rows {
		return nil, ErrorDimensionMismatch
	}
	C := Zeros(A.rows, B.cols)
	///*
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
	//*/
	/*

		for i := 0; i < A.rows; i++ {
			for j := 0; j < B.cols; j++ {
				sum := float64(0);
				for k := 0; k < A.cols; k++ {
					sum += A.elements[i*A.step+k] * B.elements[k*B.step+j]
				}
				C.elements[i*C.step+j] = sum;
			}
		}
	*/

	return C, NoError
}


func (A *DenseMatrix) ElementMult(B MatrixRO) (Matrix, Error) {
	C := A.Copy()
	err := C.ScaleMatrix(B)
	return C, err
}

func (A *DenseMatrix) ElementMultDense(B *DenseMatrix) (*DenseMatrix, Error) {
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

func (A *DenseMatrix) ScaleMatrix(B MatrixRO) Error {
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
	return NoError
}

func (A *DenseMatrix) ScaleMatrixDense(B *DenseMatrix) Error {
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
	return NoError
}
