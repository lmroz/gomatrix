// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"rand";
)

/*
A matrix backed by a flat array of all elements.
*/
type DenseMatrix struct {
	matrix;
	// flattened matrix data. elements[i*step+j] is row i, col j
	elements	[]float64;
	// actual offset between rows
	step	int;
}

/*
Returns an array of slices referencing the matrix data. Changes to
the slices effect changes to the matrix.
*/
func (A *DenseMatrix) Arrays() [][]float64 {
	a := make([][]float64, A.rows);
	for i := 0; i < A.rows; i++ {
		a[i] = A.elements[i*A.step : i*A.step+A.cols]
	}
	return a;
}

/*
Returns the contents of this matrix stored into a flat array (row-major).
*/
func (A *DenseMatrix) Array() []float64 {
	a := make([]float64, A.rows*A.cols)
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			a[i*A.rows+j] = A.elements[i*A.step+j]
		}
	}
	return a;
}

/*
Get the element in the ith row and jth column.
*/
func (A *DenseMatrix) Get(i int, j int) (v float64) {
	v = A.elements[i*A.step+j];
	return;
}


/*
Set the element in the ith row and jth column to v.
*/
func (A *DenseMatrix) Set(i int, j int, v float64) {
	A.elements[i*A.step+j] = v
}

/*
Get a submatrix starting at i,j with rows rows and cols columns. Changes to
the returned matrix show up in the original.
*/
func (A *DenseMatrix) GetMatrix(i, j, rows, cols int) *DenseMatrix {
	B := new(DenseMatrix);
	B.elements = A.elements[i*A.step+j : (i+rows)*A.step];
	B.rows = rows;
	B.cols = cols;
	B.step = A.step;
	return B;
}

func (A *DenseMatrix) GetColVector(j int) *DenseMatrix {
	return A.GetMatrix(0, j, A.rows, j+1)
}

func (A *DenseMatrix) GetRowVector(i int) *DenseMatrix {
	return A.GetMatrix(i, 0, i+1, A.cols)
}

/*
Get a copy of this matrix with 0s above the diagonal.
*/
func (A *DenseMatrix) L() *DenseMatrix {
	B := A.Copy();
	for i := 0; i < A.rows; i++ {
		for j := i + 1; j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

/*
Get a copy of this matrix with 0s below the diagonal.
*/
func (A *DenseMatrix) U() *DenseMatrix {
	B := A.Copy();
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i && j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func (A *DenseMatrix) Copy() *DenseMatrix {
	B := new(DenseMatrix);
	B.rows = A.rows;
	B.cols = A.cols;
	B.step = A.cols;
	B.elements = make([]float64, len(A.elements));
	for i := 0; i < len(A.elements); i++ {
		B.elements[i] = A.elements[i]
	}
	return B;
}

/*
Get a new matrix [A B].
*/
func (A *DenseMatrix) Augment(B *DenseMatrix) (*DenseMatrix, Error) {
	if A.Rows() != B.Rows() {
		return nil, ErrorDimensionMismatch
	}
	C := Zeros(A.Rows(), A.Cols()+B.Cols());
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j))
		}
		for j := 0; j < B.Cols(); j++ {
			C.Set(i, j+A.Cols(), B.Get(i, j))
		}
	}
	return C, NoError;
}


/*
Get a new matrix [A; B], with A above B.
*/
func (A *DenseMatrix) Stack(B *DenseMatrix) (*DenseMatrix, Error) {
	if A.Cols() != B.Cols() {
		return nil, ErrorDimensionMismatch
	}
	C := Zeros(A.Rows()+B.Rows(), A.Cols());
	for j := 0; j < A.Cols(); j++ {
		for i := 0; i < A.Rows(); i++ {
			C.Set(i, j, A.Get(i, j))
		}
		for i := 0; i < B.Cols(); i++ {
			C.Set(i+A.Rows(), j, B.Get(i, j))
		}
	}
	return C, NoError;
}

/*
Create a sparse matrix copy.
*/
func (A *DenseMatrix) SparseMatrix() *SparseMatrix {
	B := ZerosSparse(A.rows, A.cols);
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			v := A.Get(i, j);
			if v != 0 {
				B.Set(i, j, v)
			}
		}
	}
	return B;
}

func (A *DenseMatrix) DenseMatrix() *DenseMatrix {
	return A.Copy()
}

func Zeros(rows, cols int) *DenseMatrix {
	A := new(DenseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;
	return A;
}

func Ones(rows, cols int) *DenseMatrix {
	A := new(DenseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;

	for i := 0; i < len(A.elements); i++ {
		A.elements[i] = 1
	}

	return A;
}

func Numbers(rows, cols int, num float64) *DenseMatrix {
	A := Zeros(rows, cols);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, num)
		}
	}

	return A;
}

/*
Create an identity matrix with span rows and span columns.
*/
func Eye(span int) *DenseMatrix {
	A := Zeros(span, span);
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A;
}

func Normals(rows, cols int) *DenseMatrix {
	A := Zeros(rows, cols);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, rand.NormFloat64())
		}
	}

	return A;
}

func Diagonal(d []float64) *DenseMatrix {
	n := len(d);
	A := Zeros(n, n);
	for i := 0; i < n; i++ {
		A.Set(i, i, d[i])
	}
	return A;
}

func MakeDenseCopy(A MatrixRO) *DenseMatrix {
	B := Zeros(A.Rows(), A.Cols());
	for i := 0; i < B.rows; i++ {
		for j := 0; j < B.cols; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	return B;
}

func MakeDenseMatrix(elements []float64, rows, cols int) *DenseMatrix {
	A := new(DenseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;
	A.elements = elements;
	return A;
}

func MakeDenseMatrixStacked(data [][]float64) *DenseMatrix {
	rows := len(data);
	cols := len(data[0]);
	elements := make([]float64, rows*cols);
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elements[i*cols+j] = data[i][j]
		}
	}
	return MakeDenseMatrix(elements, rows, cols);
}

func (A *DenseMatrix) String() string	{ return String(A) }
