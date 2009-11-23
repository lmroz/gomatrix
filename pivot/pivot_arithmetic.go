package matrix


func (P *PivotMatrix) Times(A MatrixRO) (*DenseMatrix, *error) {
	if P.Cols() != A.Rows() {
		return nil, NewError(ErrorBadInput, "P.Times(A): P.Cols() != A.Rows()");
	}
	B := Zeros(P.rows, A.Cols());
	for i:=0; i<P.rows; i++ {
		k := 0;
		for ; i!=P.pivots[k]; k++ {}
		for j:=0; j<P.cols; j++ {
			B.Set(i, j, A.Get(k, j));
		}
	}
	return B, nil
}