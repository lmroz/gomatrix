// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

/*
The sum of this matrix and another.
*/
func (A *SparseMatrix) Plus(B MatrixRO) (*SparseMatrix, Error) {
	C := A.Copy();
	err := C.Add(B);
	return C, err;
}

/*
The sum of this matrix and another sparse matrix, optimized for sparsity.
*/
func (A *SparseMatrix) PlusSparse(B *SparseMatrix) (*SparseMatrix, Error) {
	C := A.Copy();
	err := C.AddSparse(B);
	return C, err;
}

/*
The difference between this matrix and another.
*/
func (A *SparseMatrix) Minus(B MatrixRO) (*SparseMatrix, Error) {
	C := A.Copy();
	err := C.Subtract(B);
	return C, err;
}

/*
The difference between this matrix and another sparse matrix, optimized for sparsity.
*/
func (A *SparseMatrix) MinusSparse(B *SparseMatrix) (*SparseMatrix, Error) {
	C := A.Copy();
	err := C.SubtractSparse(B);
	return C, err;
}

/*
Add another matrix to this one in place.
*/
func (A *SparseMatrix) Add(B MatrixRO) Error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}

	return NoError;
}

/*
Add another matrix to this one in place, optimized for sparsity.
*/
func (A *SparseMatrix) AddSparse(B *SparseMatrix) Error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}

	for index, value := range B.elements {
		i, j := A.GetRowColIndex(index);
		A.Set(i, j, A.Get(i, j)+value);
	}

	return NoError;
}

/*
Subtract another matrix from this one in place.
*/
func (A *SparseMatrix) Subtract(B MatrixRO) Error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.Set(i, j, A.Get(i, j)-B.Get(i, j))
		}
	}

	return NoError;
}


/*
Subtract another matrix from this one in place, optimized for sparsity.
*/
func (A *SparseMatrix) SubtractSparse(B *SparseMatrix) Error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}

	for index, value := range B.elements {
		i, j := A.GetRowColIndex(index);
		A.Set(i, j, A.Get(i, j)-value);
	}

	return NoError;
}

/*
Get the product of this matrix and another.
*/
func (A *SparseMatrix) Times(B MatrixRO) (*SparseMatrix, Error) {
	if A.cols != B.Rows() {
		return nil, ErrorDimensionMismatch
	}

	C := ZerosSparse(A.rows, B.Cols());

	for index, value := range A.elements {
		i, k := A.GetRowColIndex(index);
		//not sure if there is a more efficient way to do this without using
		//a different data structure
		for j := 0; j < B.Cols(); j++ {
			v := B.Get(k, j);
			if v != 0 {
				C.Set(i, j, C.Get(i, j)+value*v)
			}
		}
	}

	return C, NoError;
}


/*
Get the product of this matrix and another, optimized for sparsity.
*/
func (A *SparseMatrix) TimesSparse(B *SparseMatrix) (*SparseMatrix, Error) {
	return A.Times(B)	//nothing clever yet
}

/*
Scale this matrix by f.
*/
func (A *SparseMatrix) Scale(f float64) Error {
	for index, value := range A.elements {
		A.elements[index] = value * f
	}

	return NoError;
}

/*
Get the element-wise product of this matrix and another.
*/
func (A *SparseMatrix) ElementMult(B MatrixRO) (*SparseMatrix, Error) {
	C := A.Copy();
	err := C.ScaleMatrix(B);
	return C, err;
}

/*
Get the element-wise product of this matrix and another, optimized for sparsity.
*/
func (A *SparseMatrix) ElementMultSparse(B *SparseMatrix) (*SparseMatrix, Error) {
	C := A.Copy();
	err := C.ScaleMatrixSparse(B);
	return C, err;
}

/*
Scale this matrix by another, element-wise.
*/
func (A *SparseMatrix) ScaleMatrix(B MatrixRO) Error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return ErrorDimensionMismatch
	}

	for index, value := range A.elements {
		i, j := A.GetRowColIndex(index);
		A.Set(i, j, value*B.Get(i, j));
	}

	return NoError;
}

/*
Scale this matrix by another sparse matrix, element-wise. Optimized for sparsity.
*/
func (A *SparseMatrix) ScaleMatrixSparse(B *SparseMatrix) Error {
	if len(B.elements) > len(A.elements) {
		if A.rows != B.Rows() || A.cols != B.Cols() {
			return ErrorDimensionMismatch
		}

		for index, value := range B.elements {
			i, j := B.GetRowColIndex(index);
			A.Set(i, j, value*A.Get(i, j));
		}
	}
	return A.ScaleMatrix(B);
}