package matrix

func (A *SparseMatrix) Plus(B MatrixRO) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.Add(B);
	return C, err;
}

func (A *SparseMatrix) Minus(B MatrixRO) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.Subtract(B);
	return C, err;
}

func (A *SparseMatrix) Add(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		//error
		return nil;
	}
	
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		A.Set(i, j, value + B.Get(i, j))
	}
	
	return nil;
}

func (A *SparseMatrix) Subtract(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		//error
		return nil;
	}
	
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		A.Set(i, j, value - B.Get(i, j))
	}
	
	return nil;
}

func (A *SparseMatrix) Times(B MatrixRO) (*SparseMatrix, *error) {
	if A.cols != B.Rows() {
		//error
		return nil, nil
	}
	
	C := ZerosSparse(A.rows, B.Cols());
	
	for index, value := range A.elements {
		i, k := A.getRowColIndex(index);
		//not sure if there is a more efficient way to do this without using
		//a different data structure
		for j := 0; j < B.Cols(); j++ {
			v := B.Get(k, j);
			if v != 0 {
				C.Set(i, j, C.Get(i, j) + value*v);
			}			
		}
	}
	
	return C, nil
}

func (A *SparseMatrix) Scale(f float64) *error {
	for index, value := range A.elements {
		A.elements[index] = value*f;
	}
	
	return nil;
}

func (A *SparseMatrix) ScaleMatrix(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		//error
		return nil;
	}
	
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		A.Set(i, j, value * B.Get(i, j))
	}
	
	return nil;
}