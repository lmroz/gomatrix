//Copyright John Asmuth 2009
package matrix

import (
	"fmt";
	"rand";
)

const (
	_ = iota;
	ErrorNilMatrix;
	ErrorBadInput;
)

type Matrix interface {
	/* matrix.go */

	//get a sub matrix whose upper left corner is at i, j and has rows rows and cols cols
	//TODO: make these return reference matrices
	GetMatrix(i int, j int, rows int, cols int) Matrix;
	GetColVector(i int) Matrix;
	GetRowVector(j int) Matrix;

	//get the lower portion of this matrix
	L() Matrix;
	//get the upper portion of this matrix
	U() Matrix;

	//make a printable string
	String() string;
	
	/* arithmetic.go */

	//arithmetic
	Add(B Matrix);
	Subtract(B Matrix);
	Scale(f float64);

	Plus(B Matrix) Matrix;
	Minus(B Matrix) Matrix;
	Times(B Matrix) Matrix;
	ElementMult(B Matrix) Matrix;

	/* basic.go */

	Symmetric() bool;
	//check element-wise equality
	Equals(B Matrix) bool;
	//check that each element is within ε
	Approximates(B Matrix, ε float64) bool;

	swapRows(i1 int, i2 int);
	scaleRow(i int, f float64);
	scaleAddRow(i1 int, i2 int, f float64);
	//return x such that this*x = b
	Solve(b Matrix) Matrix;

	Transpose() Matrix;
	TransposeInPlace();
	Inverse() Matrix;
	Det() float64;
	Trace() float64;

	OneNorm() float64;
	TwoNorm() float64;
	InfinityNorm() float64;

	/* decomp.go */

	//returns C such that C*C' = A
	Cholesky() Matrix;
	//returns L,U,P such that P*L*U = A
	LU() (Matrix, Matrix, Matrix);
	//puts [L\U] in the matrix, L's diagonal defined to be 1s. returns the pivot
	LUInPlace() Matrix;
	QR() (Matrix, Matrix);
	//returns V,D such that V*D*inv(V) = A
	Eigen() (Matrix, Matrix);
	
	/* data.go */
	
	ErrorCode() int;
	ErrorString() string;

	// get at the raw data - returns slices so it's a reference
	Elements() []float64;
	Arrays() [][]float64;

	Rows() int;
	Cols() int;

	Get(i int, j int) float64;
	Set(i int, j int, v float64);

	RowCopy(i int) []float64;
	ColCopy(j int) []float64;
	DiagonalCopy() []float64;

	//fill a pre-allocated buffer with row i
	BufferRow(i int, buf []float64);
	//fill a pre-allocated buffer with column j
	BufferCol(j int, buf []float64);
	//fill a pre-allocated buffer with the diagonal
	BufferDiagonal(buf []float64);
	//copy a buffer into row i
	FillRow(i int, buf []float64);
	//copy a buffer into column j
	FillCol(j int, buf []float64);
	//copy a buffer into the diagonal
	FillDiagonal(buf []float64);
	
	//get a copy of this matrix
	Copy() Matrix;
}

func (A *matrix) getMatrix(i int, j int, rows int, cols int) *matrix {
	B := zeros(rows, cols);
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			B.Set(y, x, A.Get(y+i, x+j))
		}
	}
	return B;
}

//TODO: modify for reference matrices
func (A *matrix) GetMatrix(i int, j int, rows int, cols int) Matrix {
	return A.getMatrix(i, j, rows, cols)
}

func (A *matrix) GetColVector(j int) Matrix	{ return A.GetMatrix(0, j, A.rows, j+1) }

func (A *matrix) GetRowVector(i int) Matrix	{ return A.GetMatrix(i, 0, i+1, A.cols) }


func (A *matrix) L() Matrix {
	B := A.Copy();
	for i := 0; i < A.rows; i++ {
		for j := i+1; j < A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func (A *matrix) U() Matrix {
	B := A.Copy();
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i && j<A.cols; j++ {
			B.Set(i, j, 0)
		}
	}
	return B;
}

func Augment(A Matrix, B Matrix) Matrix {
	if A.Rows() != B.Rows() {
		return Error(ErrorBadInput, "Augment(A,B): A and B don't have the same number of rows");
	}
	C := zeros(A.Rows(), A.Cols()+B.Cols());
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j))
		}
		for j := 0; j < B.Cols(); j++ {
			C.Set(i, j+A.Cols(), B.Get(i, j))
		}
	}
	return C;
}

func Stack(A Matrix, B Matrix) Matrix {
	if A.Cols() != B.Cols() {
		return Error(ErrorBadInput, "Stack(A,B): A and B don't have the same number of columns");
	}
	C := zeros(A.Rows()+B.Rows(), A.Cols());
	for j := 0; j < A.Cols(); j++ {
		for i := 0; i < A.Rows(); i++ {
			C.Set(i, j, A.Get(i, j))
		}
		for i := 0; i < B.Cols(); i++ {
			C.Set(i+A.Rows(), j, B.Get(i, j))
		}
	}
	return C;
}

//TODO: modify for reference matrices
func zeros(rows int, cols int) *matrix {
	A := new(matrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	return A;
}
func Zeros(rows int, cols int) Matrix	{ return zeros(rows, cols) }

//TODO: modify for reference matrices
func numbers(rows int, cols int, num float64) *matrix {
	A := new(matrix);
	A.elements = make([]float64, rows*cols);
	for i := 0; i < rows*cols; i++ {
		A.elements[i] = num
	}
	A.rows = rows;
	A.cols = cols;
	return A;
}
func Numbers(rows int, cols int, num float64) Matrix {
	return numbers(rows, cols, num)
}

func Ones(rows int, cols int) Matrix {
	return numbers(rows, cols, 1)
}

//TODO: modify for reference matrices
func eye(span int) *matrix {
	A := zeros(span, span);
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A;
}
func Eye(span int) Matrix	{ return eye(span) }

func normals(rows int, cols int) *matrix {
	A := zeros(rows, cols);
	
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, rand.NormFloat64());	
		}
	}
	
	return A
}
func Normals(rows int, cols int) Matrix {
	return normals(rows, cols)
}

func diagonal(d []float64) *matrix {
	n := len(d);
	A := zeros(n, n);
	for i := 0; i < n; i++ {
		A.Set(i, i, d[i])
	}
	return A;
}
func Diagonal(d []float64) Matrix	{ return diagonal(d) }

//TODO: modify for reference matrices
func PivotMatrix(pivots []int, pivotSign float64) Matrix {
	n := len(pivots);
	P := new(pivotMatrix);
	P.matrix = new(matrix);
	P.elements = make([]float64, n*n);
	P.rows = n;
	P.cols = n;
	for i:=0; i<n; i++ {
		P.Set(pivots[i], i, 1)
	}
	P.pivotSign = pivotSign;
	return P
}

func Error(errorCode int, errorString string) Matrix {
	E := new(errorMatrix);
	E.errorCode = errorCode;
	E.errorString = errorString;
	return E
}

func (A *matrix) String() string {
	s := "{";
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			s += fmt.Sprintf("%f", A.Get(i, j));
			if i != A.rows-1 || j != A.cols-1 {
				s += ","
			}
			if j != A.cols-1 {
				s += " "
			}
		}
		if i != A.rows-1 {
			s += "\n"
		}
	}
	s += "}";
	return s;
}
