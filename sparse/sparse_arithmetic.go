package matrix

func (A *SparseMatrix) Plus(B MatrixRO) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.Add(B);
	return C, err;
}
func (A *SparseMatrix) PlusSparse(B *SparseMatrix) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.AddSparse(B);
	return C, err;
}

func (A *SparseMatrix) Minus(B MatrixRO) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.Subtract(B);
	return C, err;
}
func (A *SparseMatrix) MinusSparse(B *SparseMatrix) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.SubtractSparse(B);
	return C, err;
}

func (A *SparseMatrix) Add(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorBadInput, "A.Add(B): A and B dimensions don't match");
	}
	
	for i:=0; i<A.rows; i++ {
		for j:=0; j<A.cols; j++ {
			A.Set(i, j, A.Get(i, j) + B.Get(i, j))
		}
	}
	
	return nil;
}

func (A *SparseMatrix) AddSparse(B *SparseMatrix) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorBadInput, "A.Add(B): A and B dimensions don't match");
	}
	
	for index, value := range B.elements {
		i, j := A.getRowColIndex(index);
		A.Set(i, j, A.Get(i, j) + value)
	} 
	
	return nil;
}

func (A *SparseMatrix) Subtract(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorBadInput, "A.Add(B): A and B dimensions don't match");
	}
	
	for i:=0; i<A.rows; i++ {
		for j:=0; j<A.cols; j++ {
			A.Set(i, j, A.Get(i, j) - B.Get(i, j))
		}
	}
	
	return nil;
}

func (A *SparseMatrix) SubtractSparse(B *SparseMatrix) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorBadInput, "A.Subtract(B): A and B dimensions don't match");
	}
	
	for index, value := range B.elements {
		i, j := A.getRowColIndex(index);
		A.Set(i, j, A.Get(i, j) - value)
	}
	
	return nil;
}

func (A *SparseMatrix) Times(B MatrixRO) (*SparseMatrix, *error) {
	if A.cols != B.Rows() {
		return nil, NewError(ErrorBadInput, "A.Times(B): A.Cols() != B.Rows()");
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

func (A *SparseMatrix) TimesSparse(B *SparseMatrix) (*SparseMatrix, *error) {
	return A.Times(B);//nothing clever yet
}

func (A *SparseMatrix) Scale(f float64) *error {
	for index, value := range A.elements {
		A.elements[index] = value*f;
	}
	
	return nil;
}

func (A *SparseMatrix) ElementMult(B MatrixRO) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.ScaleMatrix(B);
	return C, err;
}

func (A *SparseMatrix) ElementMultSparse(B *SparseMatrix) (*SparseMatrix, *error) {
	C := A.Copy();
	err := C.ScaleMatrixSparse(B);
	return C, err;
}

func (A *SparseMatrix) ScaleMatrix(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorBadInput, "A.ScaleMatrix(B): A and B dimensions don't match");
	}
	
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		A.Set(i, j, value * B.Get(i, j))
	}
	
	return nil;
}

func (A *SparseMatrix) ScaleMatrixSparse(B *SparseMatrix) *error {
	if len(B.elements) > len(A.elements) {
		if A.rows != B.Rows() || A.cols != B.Cols() {
			return NewError(ErrorBadInput, "A.ScaleMatrix(B): A and B dimensions don't match");
		}
		
		for index, value := range B.elements {
			i, j := B.getRowColIndex(index);
			A.Set(i, j, value * A.Get(i, j))
		}
	}
	return A.ScaleMatrix(B);
}
