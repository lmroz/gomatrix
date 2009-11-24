package matrix

func (P *PivotMatrix) Times(A MatrixRO) (*DenseMatrix, *error) {
	if P.Cols() != A.Rows() {
		return nil, NewError(ErrorDimensionMismatch);
	}
	B := Zeros(P.rows, A.Cols());
	for i:=0; i<P.rows; i++ {
		k := 0;
		for ; i!=P.pivots[k]; k++ {}
		for j:=0; j<A.Cols(); j++ {
			B.Set(i, j, A.Get(k, j));
		}
	}
	return B, nil
}
func (P *PivotMatrix) TimesPivot(A *PivotMatrix) (*PivotMatrix, *error) {
	//TODO: this method
	return nil, nil;
}