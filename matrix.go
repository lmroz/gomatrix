//Copyright John Asmuth 2009
package matrix

import (
	"fmt";
	"math";
)

/*
 * For matrix.matrixType. Some things take less time if it's known
 * which one of these types the matrix is. nil for a normal matrix.
 */
const (
	_	= iota;
	// upper matrices have only zeros below the diagonal
	upper;
	// lower matrices have only zeros above the diagonal
	lower;
	// pivot matrices are permutations of the identity
	pivot;
)

type matrix struct {
	// flattened matrix data. elements[i*cols+j] is row i, col j
	elements	[]float64;
	// the number of rows
	rows	int;
	// the number of columns
	cols	int;

	// the type of matrix {_, upper, lower, pivot}
	matrixType	int;
	// if this is a pivot matrix, the determinant goes here
	pivotSign	float64;
}

type Matrix interface {
	// exchange two rows in this matrix
	swapRows(i1 int, i2 int);
	// multiply all elements in this row by a constant
	scaleRow(i int, f float64);
	// i1 = i1+f*i2
	scaleAddRow(i1 int, i2 int, f float64);

	// get at the raw data
	Elements() []float64;
	Arrays() [][]float64;

	Symmetric() bool;

	Rows() int;
	Cols() int;

	Get(i int, j int) float64;
	Set(i int, j int, v float64);

	GetRow(i int) []float64;
	GetCol(j int) []float64;
	GetDiagonal() []float64;

	//fill a pre-allocated buffer with row i
	BufferRow(i int, buf []float64);
	//fill a pre-allocated buffer with column j
	BufferCol(j int, buf []float64);
	//fill a pre-allocated buffer with the diagonal
	BufferDiagonal(buf []float64);

	//append B to the right of this matrix
	StackHorizontal(B Matrix) Matrix;
	//append B below this matrix
	StackVertical(B Matrix) Matrix;

	//get a sub matrix whose upper left corner is at i, j and has rows rows and cols cols
	GetMatrix(i int, j int, rows int, cols int) Matrix;
	//get a column in matrix form
	GetColVector(i int) Matrix;
	//get a row in matrix form
	GetRowVector(j int) Matrix;

	//add this matrix to another
	Plus(B Matrix) Matrix;
	//subtract the other matrix from this one
	Minus(B Matrix) Matrix;
	//multiply this matrix by another
	Times(B Matrix) Matrix;
	//multiply every element in this matrix by a scalar
	Scale(f float64) Matrix;
	//multiply each element in this matrix by the corresponding element in another
	ElementMult(B Matrix) Matrix;

	//add this matrix to another in place, return self
	PlusInPlace(B Matrix) Matrix;
	//subtract the other matrix from this one in place, return self
	MinusInPlace(B Matrix) Matrix;
	//multiply every element in this matrix by a scalar in place, return self
	ScaleInPlace(f float64) Matrix;

	//run the multiplication with each dot product done in parallel
	ParallelTimes(B Matrix, threads int) Matrix;

	//check element-wise equality
	Equals(B Matrix) bool;
	//check that each element is within ε
	Approximates(B Matrix, ε float64) bool;

	OneNorm() float64;
	TwoNorm() float64;
	InfinityNorm() float64;

	Transpose() Matrix;
	Inverse() Matrix;
	Cholesky() Matrix;
	LU() (Matrix, Matrix, Matrix);
	QR() (Matrix, Matrix);
	Eigen() (Matrix, Matrix);
	
	TransposeInPlace() Matrix;
	//puts [L\U] in the matrix, L's diagonal defined to be 1s. returns the pivot
	LUInPlace() Matrix;

	//get the lower portion of this matrix
	L() Matrix;
	//get the upper portion of this matrix
	U() Matrix;

	Det() float64;
	Trace() float64;

	//return x such that this*x = b
	Solve(b Matrix) Matrix;

	//get a copy of this matrix
	Copy() Matrix;

	//format for printing
	String() string;
}

func (m *matrix) swapRows(r1 int, r2 int) {
	for i := 0; i < m.cols; i++ {
		i1, i2 := r1*m.cols+i, r2*m.cols+i;
		tmp := m.elements[i1];
		m.elements[i1] = m.elements[i2];
		m.elements[i2] = tmp;
	}
}

func (m *matrix) scaleRow(r int, f float64) {
	for i := 0; i < m.cols; i++ {
		m.elements[r*m.cols+i] *= f
	}
}

func (m *matrix) scaleAddRow(rd int, rs int, f float64) {
	for i := 0; i < m.cols; i++ {
		m.elements[rd*m.cols+i] += m.elements[rs*m.cols+i] * f
	}
}

func (A *matrix) Symmetric() bool {
	if A.rows != A.cols {
		return false
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i; j++ {
			if A.elements[i*A.cols+j] != A.elements[j*A.cols+i] {
				return false
			}
		}
	}
	return true;
}

func (A *matrix) StackHorizontal(B Matrix) Matrix {
	if A.Rows() != B.Rows() {
		return nil
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

func (A *matrix) StackVertical(B Matrix) Matrix {
	if A.Cols() != B.Cols() {
		return nil
	}
	C := zeros(A.Rows()+B.Rows(), A.Cols());
	for j := 0; j < A.Cols(); j++ {
		for i := 0; i < A.Rows(); i++ {
			C.Set(i, j, A.Get(i, j))
		}
		for i := 0; i < B.Rows(); i++ {
			C.Set(i, j+A.Rows(), B.Get(i, j))
		}
	}
	return C;
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

func (A *matrix) GetMatrix(i int, j int, rows int, cols int) Matrix {
	return A.getMatrix(i, j, rows, cols)
}
func (A *matrix) GetColVector(j int) Matrix	{ return A.GetMatrix(0, j, A.rows, j+1) }

func (A *matrix) GetRowVector(i int) Matrix	{ return A.GetMatrix(i, 0, i+1, A.cols) }

func (A *matrix) Plus(B Matrix) Matrix {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return nil
	}

	C := zeros(A.Rows(), A.Cols());

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}
	return C;
}

func (A *matrix) Minus(B Matrix) Matrix {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return nil
	}

	C := zeros(A.Rows(), A.Cols());

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j)-B.Get(i, j))
		}
	}
	return C;
}

func (A *matrix) Times(B Matrix) Matrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	C := zeros(A.Rows(), B.Cols());

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0);
			for k := 0; k < A.Cols(); k++ {
				sum += A.Get(i, k) * B.Get(k, j)
			}
			C.Set(i, j, sum);
		}
	}

	return C;
}


type ij struct {
	i	int;
	j	int;
}

func dotRowCol(A Matrix, B Matrix, C Matrix, in chan ij, quit chan bool) {
	for true {
		select {
		case ij_ := <-in:
			sum := float64(0);
			for k := 0; k < A.Cols(); k++ {
				sum += A.Get(ij_.i, k) * B.Get(k, ij_.j)
			}
			C.Set(ij_.i, ij_.j, sum);
		case <-quit:
			return
		}
	}
}

func (A *matrix) ParallelTimes(B Matrix, threads int) Matrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	C := zeros(A.Rows(), B.Cols());

	in := make(chan ij);
	quit := make(chan bool);

	for i := 0; i < threads; i++ {
		go dotRowCol(A, B, C, in, quit)
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			ij_ := ij{i, j};
			in <- ij_;
		}
	}

	for i := 0; i < threads; i++ {
		quit <- true
	}

	return C;
}

func (A *matrix) ElementMult(B Matrix) Matrix {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return nil
	}
	C := zeros(A.rows, A.cols);
	Belements := B.Elements();
	for i := 0; i < len(C.elements); i++ {
		C.elements[i] = A.elements[i] * Belements[i]
	}
	return C;
}

func (A *matrix) Scale(f float64) Matrix {
	B := zeros(A.Rows(), A.Cols());
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			B.Set(i, j, f*A.Get(i, j))
		}
	}
	return B;
}

func (A *matrix) PlusInPlace(B Matrix) Matrix {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return nil
	}
	Belements := B.Elements();
	for i := 0; i < len(A.elements); i++ {
		A.elements[i] += Belements[i]
	}
	return A;
}

func (A *matrix) MinusInPlace(B Matrix) Matrix {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return nil
	}
	Belements := B.Elements();
	for i := 0; i < len(A.elements); i++ {
		A.elements[i] -= Belements[i]
	}
	return A;
}

func (A *matrix) ScaleInPlace(f float64) Matrix {
	for i := 0; i < len(A.elements); i++ {
		A.elements[i] *= f
	}
	return A;
}

func (A *matrix) Equals(B Matrix) bool {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return false
	}
	Belements := B.Elements();
	for i := 0; i < len(A.elements); i++ {
		if A.elements[i] != Belements[i] {
			return false
		}
	}
	return true;
}

func (A *matrix) Approximates(B Matrix, ε float64) bool {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return false
	}
	Belements := B.Elements();
	for i := 0; i < len(A.elements); i++ {
		if math.Fabs(A.elements[i]-Belements[i]) > ε {
			return false
		}
	}
	return true;
}

func (A *matrix) OneNorm() (ε float64) {
	for i := 0; i < len(A.elements); i++ {
		if A.elements[i] > ε {
			ε = A.elements[i]
		}
	}
	return;
}

func (A *matrix) TwoNorm() float64 {
	//requires computing of eigenvalues
	return 0
}

func (A *matrix) InfinityNorm() (ε float64) {
	for i := 0; i < len(A.elements); i++ {
		ε += A.elements[i]
	}
	return;
}

func (A *matrix) Transpose() Matrix {
	B := zeros(A.Cols(), A.Rows());
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			B.Set(j, i, A.Get(i, j))
		}
	}
	return B;
}

func (A *matrix) TransposeInPlace() Matrix {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			tmp := A.elements[i*A.cols+j];
			A.elements[i*A.cols+j] = A.elements[j*A.cols+i];
			A.elements[j*A.cols+i] = tmp
		}
	}
	return A
}

func (A *matrix) L() Matrix {
	B := zeros(A.rows, A.cols);
	for i := 0; i < A.rows; i++ {
		for j := 0; j <= i; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	B.matrixType = lower;
	return B;
}

func (A *matrix) U() Matrix {
	B := zeros(A.rows, A.cols);
	for i := 0; i < A.rows; i++ {
		for j := i; j < A.cols; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	B.matrixType = upper;
	return B;
}

func zeros(rows int, cols int) *matrix {
	A := new(matrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	return A;
}

func Zeros(rows int, cols int) Matrix	{ return zeros(rows, cols) }

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

func Numbers(rows int, cols int, num float64) Matrix	{ return numbers(rows, cols, num) }

func eye(span int) *matrix {
	A := zeros(span, span);
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A;
}

func Eye(span int) Matrix	{ return eye(span) }

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
