package matrix

func (A *DenseMatrix) Plus(B *DenseMatrix) *DenseMatrix {
	res, _ := Sum(A, B);
	return res.(*DenseMatrix);
}

func (A *DenseMatrix) Minus(B *DenseMatrix) *DenseMatrix {
	res, _ := Difference(A, B);
	return res.(*DenseMatrix);
}

func (A *DenseMatrix) Add(B Matrix) Error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return NewError(ErrorBadInput, "A.Subtract(B): A and B dimensions don't match");
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}
	
	return nil
}

func (A *DenseMatrix) Subtract(B Matrix) Error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return NewError(ErrorBadInput, "A.Subtract(B): A and B dimensions don't match");
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, A.Get(i, j)-B.Get(i, j))
		}
	}
	
	return nil
}

func (A *DenseMatrix) Times(B *DenseMatrix) *DenseMatrix {
	res, _ := Product(A, B);
	return res;
}


func (A *DenseMatrix) ElementMult(B Matrix) (*DenseMatrix, Error) {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return nil, NewError(ErrorBadInput, "ElementMult(A, B):A and B have different dimensions")
	}
	C := Zeros(A.rows, A.cols);
	for i := 0; i < C.rows; i++ {
		for j := 0; j < C.cols; j++ {
			C.Set(i, j, A.Get(i, j)*B.Get(i, j))
		}
	}
	return C, nil;
}

func (A *DenseMatrix) Scale(f float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.Set(i, j, A.Get(i, j)*f)
		}
	}
}

func (A *DenseMatrix) ScaleMatrix(B Matrix) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.Set(i, j, A.Get(i, j)*B.Get(i, j))
		}
	}
}

