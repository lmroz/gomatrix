package matrix

func (A *DenseMatrix) Plus(B MatrixRO) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.Add(B);
	return C, err;
}

func (A *DenseMatrix) Minus(B MatrixRO) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.Subtract(B);
	return C, err;
}

func (A *DenseMatrix) Add(B MatrixRO) *error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return NewError(ErrorBadInput, "A.Subtract(B): A and B dimensions don't match");
	}

	for i := 0; i < A.rows; i++ {
		index := i*A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] += B.Get(i, j);
			index++;
		}
	}
	
	return nil
}

func (A *DenseMatrix) Subtract(B MatrixRO) *error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return NewError(ErrorBadInput, "A.Subtract(B): A and B dimensions don't match");
	}

	for i := 0; i < A.rows; i++ {
		index := i*A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] -= B.Get(i, j);
			index++;
		}
	}
	
	return nil
}

func (A *DenseMatrix) Times(B MatrixRO) (*DenseMatrix, *error) {
	if A.Cols() != B.Rows() {
		return nil, NewError(ErrorBadInput, "A.Times(B): A.Cols() != B.Rows()");
	}
	C := Zeros(A.Rows(), B.Cols());

	for i := 0; i < A.rows; i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0);
			for k := 0; k < A.cols; k++ {
				sum += A.Get(i, k) * B.Get(k, j)
			}
			C.Set(i, j, sum);
		}
	}

	return C, nil;
}


func (A *DenseMatrix) ElementMult(B MatrixRO) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.ScaleMatrix(B);
	return C, err;
}

func (A *DenseMatrix) Scale(f float64) {
	for i := 0; i < A.rows; i++ {
		index := i*A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] *= f;
			index++;
		}
	}
}

func (A *DenseMatrix) ScaleMatrix(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorBadInput, "A.ScaleMatrix(B):A and B have different dimensions")
	}
	for i := 0; i < A.rows; i++ {
		index := i*A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] *= B.Get(i, j);
			index++;
		}
	}
	return nil;
}

