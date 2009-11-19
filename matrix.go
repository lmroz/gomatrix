//Copyright John Asmuth 2009
package matrix

import (
	"fmt";
	"rand";
)


// Interfaces are more like contracts than classes; only the functions
// that absolutely need to be here should be here.  The interface should
// represent the set of functions on which all other matrix functions are based.
type Matrix interface {
	/* matrix.go */

	//get a sub matrix whose upper left corner is at i, j and has rows rows and cols cols
	//TODO: make these return reference matrices
	//	GetMatrix(i int, j int, rows int, cols int) Matrix;
	//	GetColVector(i int) Matrix;
	//	GetRowVector(j int) Matrix;

	//make a printable string
	String() string;

	/* arithmetic.go */

	//arithmetic
	Add(B Matrix);
	Subtract(B Matrix);
	Scale(f float64);

	//	Plus(B Matrix) Matrix;
	//	Minus(B Matrix) Matrix;
	//	Times(B Matrix) Matrix;
	//	ElementMult(B Matrix) Matrix;

	/* basic.go */

	//	Symmetric() bool;
	//check element-wise equality
	//	Equals(B Matrix) bool;
	//check that each element is within ε
	//	Approximates(B Matrix, ε float64) bool;

	//	SwapRows(i1 int, i2 int);
	//	ScaleRow(i int, f float64);
	//	ScaleAddRow(i1 int, i2 int, f float64);
	//return x such that this*x = b
	//	Solve(b Matrix) Matrix;

	//	Transpose() Matrix;
	//	TransposeInPlace();
	//	Inverse() Matrix;
	Det() float64;
	Trace() float64;

	//	OneNorm() float64;
	//	TwoNorm() float64;
	//	InfinityNorm() float64;

	/* decomp.go */

	//returns C such that C*C' = A
	//	Cholesky() Matrix;
	//returns L,U,P such that P*L*U = A
	//	LU() (Matrix, Matrix, Matrix);
	//puts [L\U] in the matrix, L's diagonal defined to be 1s. returns the pivot
	//	LUInPlace() Matrix;
	//	QR() (Matrix, Matrix);
	//returns V,D such that V*D*inv(V) = A
	//	Eigen() (Matrix, Matrix);

	/* data.go */

	//	ErrorCode() int;
	//	ErrorString() string;

	// get at the raw data - returns slices so it's a reference
	//	Elements() []float64;
	//	Arrays() [][]float64;

	Rows() int;
	Cols() int;

	// number of elements in the matrix
	NumElements() int;

	Get(i int, j int) float64;
	Set(i int, j int, v float64);
	GetSize() (int, int);

	//	RowCopy(i int) []float64;
	//	ColCopy(j int) []float64;
	//	DiagonalCopy() []float64;

	//fill a pre-allocated buffer with row i
	//	BufferRow(i int, buf []float64);
	//fill a pre-allocated buffer with column j
	//	BufferCol(j int, buf []float64);
	//fill a pre-allocated buffer with the diagonal
	//	BufferDiagonal(buf []float64);
	//copy a buffer into row i
	//	FillRow(i int, buf []float64);
	//copy a buffer into column j
	//	FillCol(j int, buf []float64);
	//copy a buffer into the diagonal
	//	FillDiagonal(buf []float64);

	//get a copy of this matrix
	Copy() Matrix;
}

func (A *denseMatrix) getMatrix(i int, j int, rows int, cols int) *denseMatrix {
	/*
		B := zeros(rows, cols);
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				B.Set(y, x, A.Get(y+i, x+j))
			}
		}
		return B;
	*/
	B := new(denseMatrix);
	B.elements = A.elements[i*A.step+j : (i+rows)*A.step];
	B.rows = rows;
	B.cols = cols;
	B.step = A.step;
	return B;
}

func (A *denseMatrix) GetMatrix(i int, j int, rows int, cols int) *denseMatrix {
	return A.getMatrix(i, j, rows, cols)
}

func (A *denseMatrix) GetColVector(j int) *denseMatrix {
	return A.GetMatrix(0, j, A.rows, j+1)
}

func (A *denseMatrix) GetRowVector(i int) *denseMatrix {
	return A.GetMatrix(i, 0, i+1, A.cols)
}


func (A *denseMatrix) L() *denseMatrix {
	B := A.copy();
	for i := 0; i < A.rows; i++ {
		for j := i + 1; j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func (A *denseMatrix) U() *denseMatrix {
	B := A.copy();
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i && j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func Augment(A Matrix, B Matrix) (*denseMatrix, Error) {
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

func Stack(A Matrix, B Matrix) (*denseMatrix, Error) {
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

func Zeros(rows int, cols int) *denseMatrix {
	A := new(denseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;
	return A;
}

func Ones(rows int, cols int) *denseMatrix {
	A := new(denseMatrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	A.step = cols;

	for i := 0; i < len(A.elements); i++ {
		A.elements[i] = 1
	}

	return A;
}

func NewMatrix(rows int, cols int) *denseMatrix {
	return Zeros(rows, cols)
}

func Numbers(rows int, cols int, num float64) *denseMatrix {
	A := new(denseMatrix);
	A.elements = make([]float64, rows*cols);
	for i := 0; i < rows*cols; i++ {
		A.elements[i] = num
	}
	A.rows = rows;
	A.cols = cols;
	A.step = cols;
	return A;
}

func Eye(span int) *denseMatrix {
	A := Zeros(span, span);
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A;
}

func Normals(rows int, cols int) *denseMatrix {
	A := Zeros(rows, cols);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, rand.NormFloat64())
		}
	}

	return A;
}

func Diagonal(d []float64) *denseMatrix {
	n := len(d);
	A := Zeros(n, n);
	for i := 0; i < n; i++ {
		A.Set(i, i, d[i])
	}
	return A;
}

/*func PivotMatrix(pivots []int, pivotSign float64) *pivotMatrix {
	n := len(pivots);
	P := new(pivotMatrix);
	P.denseMatrix = new(denseMatrix);
	P.elements = make([]float64, n*n);
	P.rows = n;
	P.cols = n;
	P.step = n;
	for i := 0; i < n; i++ {
		P.Set(pivots[i], i, 1)
	}
	P.pivotSign = pivotSign;
	return P;
}*/

func (A *denseMatrix) String() string {
	s := "{";
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			s += fmt.Sprintf("%f", A.Get(i, j));
			if i != A.Rows()-1 || j != A.Cols()-1 {
				s += ","
			}
			if j != A.cols-1 {
				s += " "
			}
		}
		if i != A.Rows()-1 {
			s += "\n"
		}
	}
	s += "}";
	return s;
}

