// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

/*
Multiply this pivot matrix by another.
*/
func (P *PivotMatrix) Times(A MatrixRO) (*DenseMatrix, Error) {
	if P.Cols() != A.Rows() {
		return nil, ErrorDimensionMismatch
	}
	B := Zeros(P.rows, A.Cols());
	for i := 0; i < P.rows; i++ {
		k := 0;
		for ; i != P.pivots[k]; k++ {
		}
		for j := 0; j < A.Cols(); j++ {
			B.Set(i, j, A.Get(k, j))
		}
	}
	return B, NoError;
}

/*
Multiplication optimized for when two pivots are the operands.
*/
func (P *PivotMatrix) TimesPivot(A *PivotMatrix) (*PivotMatrix, Error) {
	if P.rows != A.rows {
		return nil, ErrorDimensionMismatch
	}

	newPivots := make([]int, P.rows);
	newSign := P.pivotSign * A.pivotSign;

	for i := 0; i < A.rows; i++ {
		newPivots[i] = P.pivots[A.pivots[i]]
	}

	return MakePivotMatrix(newPivots, newSign), NoError;
}

/*
Equivalent to PxA, but streamlined to take advantage of the datastructures.
*/
func (P *PivotMatrix) RowPivotDense(A *DenseMatrix) (*DenseMatrix, Error) {
	if P.rows != A.rows {
		return nil, ErrorDimensionMismatch
	}
	B := Zeros(A.rows, A.cols);
	for si := 0; si < A.rows; si++ {
		di := P.pivots[si];
		Astart := si * A.step;
		Bstart := di * B.step;
		for j := 0; j < A.cols; j++ {
			B.elements[Bstart+j] = A.elements[Astart+j]
		}
	}
	return B, NoError;
}

/*
Equivalent to AxP, but streamlined to take advantage of the datastructures.
*/
func (P *PivotMatrix) ColPivotDense(A *DenseMatrix) (*DenseMatrix, Error) {
	if P.rows != A.cols {
		return nil, ErrorDimensionMismatch
	}
	B := Zeros(A.rows, A.cols);
	for i := 0; i < B.rows; i++ {
		Astart := i * A.step;
		Bstart := i * B.step;
		for sj := 0; sj < B.cols; sj++ {
			dj := P.pivots[sj];
			B.elements[Bstart+dj] = A.elements[Astart+sj];
		}
	}
	return B, NoError;
}

/*
Equivalent to PxA, but streamlined to take advantage of the datastructures.
*/
func (P *PivotMatrix) RowPivotSparse(A *SparseMatrix) (*SparseMatrix, Error) {
	if P.rows != A.rows {
		return nil, ErrorDimensionMismatch
	}
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		si, j := A.GetRowColIndex(index);
		di := P.pivots[si];
		B.Set(di, j, value);
	}

	return B, NoError;
}

/*
Equivalent to AxP, but streamlined to take advantage of the datastructures.
*/
func (P *PivotMatrix) ColPivotSparse(A *SparseMatrix) (*SparseMatrix, Error) {
	if P.rows != A.cols {
		return nil, ErrorDimensionMismatch
	}
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		i, sj := A.GetRowColIndex(index);
		dj := P.pivots[sj];
		B.Set(i, dj, value);
	}

	return B, NoError;
}
