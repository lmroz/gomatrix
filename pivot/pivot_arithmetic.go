package matrix

func (P *PivotMatrix) Times(A MatrixRO) (*DenseMatrix, *error) {
	if P.Cols() != A.Rows() {
		return nil, NewError(ErrorDimensionMismatch)
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
	return B, nil;
}

func (P *PivotMatrix) TimesPivot(A *PivotMatrix) (*PivotMatrix, *error) {
	if P.rows != A.rows {
		return nil, NewError(ErrorDimensionMismatch)
	}

	newPivots := make([]int, P.rows);
	newSign := P.pivotSign * A.pivotSign;

	for i := 0; i < A.rows; i++ {
		newPivots[i] = P.pivots[A.pivots[i]]
	}

	return MakePivotMatrix(newPivots, newSign), nil;
}

func (P *PivotMatrix) RowPivotDense(A *DenseMatrix) (*DenseMatrix, *error) {
	if P.rows != A.rows {
		return nil, NewError(ErrorDimensionMismatch)
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
	return B, nil;
}

func (P *PivotMatrix) ColPivotDense(A *DenseMatrix) (*DenseMatrix, *error) {
	if P.rows != A.cols {
		return nil, NewError(ErrorDimensionMismatch)
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
	return B, nil;
}

func (P *PivotMatrix) RowPivotSparse(A *SparseMatrix) (*SparseMatrix, *error) {
	if P.rows != A.rows {
		return nil, NewError(ErrorDimensionMismatch)
	}
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		si, j := A.GetRowColIndex(index);
		di := P.pivots[si];
		B.Set(di, j, value);
	}

	return B, nil;
}

func (P *PivotMatrix) ColPivotSparse(A *SparseMatrix) (*SparseMatrix, *error) {
	if P.rows != A.cols {
		return nil, NewError(ErrorDimensionMismatch)
	}
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		i, sj := A.GetRowColIndex(index);
		dj := P.pivots[sj];
		B.Set(i, dj, value);
	}

	return B, nil;
}
