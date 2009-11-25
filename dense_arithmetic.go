package matrix

func (A *DenseMatrix) Plus(B *DenseMatrix) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.Add(B);
	return C, err;
}
func (A *DenseMatrix) PlusDense(B *DenseMatrix) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.AddDense(B);
	return C, err;
}

func (A *DenseMatrix) Minus(B *DenseMatrix) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.Subtract(B);
	return C, err;
}

func (A *DenseMatrix) MinusDense(B *DenseMatrix) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.SubtractDense(B);
	return C, err;
}

func (A *DenseMatrix) Add(B MatrixRO) *error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return NewError(ErrorDimensionMismatch)
	}

	for i := 0; i < A.rows; i++ {
		index := i * A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] += B.Get(i, j);
			index++;
		}
	}

	return nil;
}

func (A *DenseMatrix) AddDense(B *DenseMatrix) *error {
	return A.Add(B)
}

func (A *DenseMatrix) Subtract(B MatrixRO) *error {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return NewError(ErrorDimensionMismatch)
	}

	for i := 0; i < A.rows; i++ {
		index := i * A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] -= B.Get(i, j);
			index++;
		}
	}

	return nil;
}

func (A *DenseMatrix) SubtractDense(B *DenseMatrix) *error {
	return A.Subtract(B)
}

func (A *DenseMatrix) Times(B MatrixRO) (*DenseMatrix, *error) {
	if A.cols != B.Rows() {
		return nil, NewError(ErrorDimensionMismatch)
	}
	C := Zeros(A.rows, B.Cols());

	for i := 0; i < A.rows; i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0);
			for k := 0; k < A.cols; k++ {
				sum += A.elements[i*A.step+k] * B.Get(k, j)
			}
			C.elements[i*C.step+j] = sum;
		}
	}

	return C, nil;
}

func (A *DenseMatrix) TimesDense(B *DenseMatrix) (*DenseMatrix, *error) {
	if A.cols != B.rows {
		return nil, NewError(ErrorDimensionMismatch)
	}
	C := Zeros(A.rows, B.cols);

	//Astart := 0;
	for i := 0; i < A.rows; i++ {
		for j := 0; j < B.cols; j++ {
			//Bstart := j;
			sum := float64(0);
			for k := 0; k < A.cols; k++ {
				sum += A.elements[i*A.step+k] * B.elements[k*B.step+j]
				//sum += A.elements[i*A.step+k] * B.Get(k, j);

				//for some reason this next line is *slower*...
				//sum += A.elements[Astart+k] * B.elements[k*B.step+j];

				//slowest, though a more mature compiler might inline it to be
				//like the first version
				//sum += A.Get(i, k) * B.Get(k, j);
			}
			C.elements[i*C.step+j] = sum;
		}
		//Astart += A.step;
	}

	return C, nil;
}


func (A *DenseMatrix) ElementMult(B MatrixRO) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.ScaleMatrix(B);
	return C, err;
}

func (A *DenseMatrix) ElementMultDense(B *DenseMatrix) (*DenseMatrix, *error) {
	C := A.Copy();
	err := C.ScaleMatrixDense(B);
	return C, err;
}

func (A *DenseMatrix) Scale(f float64) {
	for i := 0; i < A.rows; i++ {
		index := i * A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[index] *= f;
			index++;
		}
	}
}

func (A *DenseMatrix) ScaleMatrix(B MatrixRO) *error {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return NewError(ErrorDimensionMismatch)
	}
	for i := 0; i < A.rows; i++ {
		indexA := i * A.step;
		for j := 0; j < A.cols; j++ {
			A.elements[indexA] *= B.Get(i, j);
			indexA++;
		}
	}
	return nil;
}

func (A *DenseMatrix) ScaleMatrixDense(B *DenseMatrix) *error {
	if A.rows != B.rows || A.cols != B.cols {
		return NewError(ErrorDimensionMismatch)
	}
	for i := 0; i < A.rows; i++ {
		indexA := i * A.step;
		indexB := i * B.step;
		for j := 0; j < A.cols; j++ {
			A.elements[indexA] *= B.elements[indexB];
			indexA++;
			indexB++;
		}
	}
	return nil;
}
