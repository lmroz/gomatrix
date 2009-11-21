package matrix

import (
	"rand";
)

type DenseMatrix struct {
	matrix;
	// flattened matrix data. elements[i*step+j] is row i, col j
	elements	[]float64;
	// actual offset between rows
	step	int;
}

//This returns an array of slices referencing the matrix data. Changes to
//the slices effect changes to the matrix
func (A *DenseMatrix) Arrays() [][]float64 {
	a := make([][]float64, A.rows);
	for i := 0; i < A.rows; i++ {
		a[i] = A.elements[i*A.step : i*A.step+A.cols]
	}
	return a;
}

func (A *DenseMatrix) Get(i int, j int) float64 {
	return A.elements[i*A.step+j]
}

func (A *DenseMatrix) Set(i int, j int, v float64) {
	A.elements[i*A.step+j] = v
}


func (A *DenseMatrix) getMatrix(i int, j int, rows int, cols int) *DenseMatrix {
	B := new(DenseMatrix);
	B.elements = A.elements[i*A.step+j : (i+rows)*A.step];
	B.rows = rows;
	B.cols = cols;
	B.step = A.step;
	return B;
}

func (A *DenseMatrix) GetMatrix(i int, j int, rows int, cols int) *DenseMatrix {
	return A.getMatrix(i, j, rows, cols)
}

func (A *DenseMatrix) GetColVector(j int) *DenseMatrix {
	return A.GetMatrix(0, j, A.rows, j+1)
}

func (A *DenseMatrix) GetRowVector(i int) *DenseMatrix {
	return A.GetMatrix(i, 0, i+1, A.cols)
}


func (A *DenseMatrix) L() *DenseMatrix {
	B := A.Copy();
	for i := 0; i < A.rows; i++ {
		for j := i + 1; j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func (A *DenseMatrix) U() *DenseMatrix {
	B := A.Copy();
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i && j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func (A *DenseMatrix) Copy() *DenseMatrix {
	B := Zeros(A.rows, A.cols);
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	return B;
}
func (A *DenseMatrix) copyMatrix() Matrix {
	return A.Copy();
}
func (A *DenseMatrix) copyMatrixReadOnly() MatrixRO {
	return A.Copy();
}


func (A *DenseMatrix) Augment(B *DenseMatrix) (*DenseMatrix, Error) {
	if A.Rows() != B.Rows() {
		return nil, NewError(ErrorBadInput, "Augment(A,B): A and B don't have the same number of rows")
	}
	C := Zeros(A.Rows(), A.Cols()+B.Cols());
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j))
		}
		for j := 0; j < B.Cols(); j++ {
			C.Set(i, j+A.Cols(), B.Get(i, j))
		}
	}
	return C, nil;
}

func (A *DenseMatrix) Stack(B *DenseMatrix) (*DenseMatrix, Error) {
	if A.Cols() != B.Cols() {
		return nil, NewError(ErrorBadInput, "Stack(A,B): A and B don't have the same number of columns")
	}
	C := Zeros(A.Rows()+B.Rows(), A.Cols());
	for j := 0; j < A.Cols(); j++ {
		for i := 0; i < A.Rows(); i++ {
			C.Set(i, j, A.Get(i, j))
		}
		for i := 0; i < B.Cols(); i++ {
			C.Set(i+A.Rows(), j, B.Get(i, j))
		}
	}
	return C, nil;
}

func Zeros(rows int, cols int) *DenseMatrix {
	A := new(DenseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;
	return A;
}

func Ones(rows int, cols int) *DenseMatrix {
	A := new(DenseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;

	for i := 0; i < len(A.elements); i++ {
		A.elements[i] = 1
	}

	return A;
}

func Numbers(rows int, cols int, num float64) *DenseMatrix {
	A := Zeros(rows, cols);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, num)
		}
	}

	return A;
}

func Eye(span int) *DenseMatrix {
	A := Zeros(span, span);
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A;
}

func Normals(rows int, cols int) *DenseMatrix {
	A := Zeros(rows, cols);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, rand.NormFloat64())
		}
	}

	return A;
}

func Diagonal(d []float64) *DenseMatrix {
	n := len(d);
	A := Zeros(n, n);
	for i := 0; i < n; i++ {
		A.Set(i, i, d[i])
	}
	return A;
}

func MakeDenseCopy(A MatrixRO) *DenseMatrix {
	B := Zeros(A.Rows(), A.Cols());
	for i := 0; i < B.rows; i++ {
		for j := 0; j < B.cols; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	return B
}

func MakeDenseMatrix(elements []float64, rows int, cols int) *DenseMatrix {

	A := Zeros(rows, cols);
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			A.Set(i, j, elements[i*cols+j])
		}
	}
	return A;
}

func MakeDenseMatrixStacked(data [][]float64) *DenseMatrix {
	rows := len(data);
	cols := len(data[0]);
	elements := make([]float64, rows*cols);
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elements[i*cols+j] = data[i][j]
		}
	}
	return MakeDenseMatrix(elements, rows, cols);
}

func (A *DenseMatrix) SparseMatrix() *SparseMatrix {
	B := ZerosSparse(A.rows, A.cols);
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			v := A.Get(i, j);
			if v != 0 {
				B.Set(i, j, v);
			}
		}
	}
	return B;
}


