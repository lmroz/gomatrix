//Copyright John Asmuth 2009
package matrix

import (
	"fmt"
)

const (
	_ = iota;
	upper;
	lower;
	pivot;
)

type matrix struct {
	elements	[]float64;
	rows		int;
	cols		int;
	
	matrixType	int;
	pivotSign	float64;
}

type Matrix interface {
	swapRows(int, int);
	scaleRow(int, float64);
	scaleAddRow(int, int, float64);

	Elements() []float64;

	Rows() int;
	Cols() int;
	
	Get(i int, j int) float64;
	Set(i int, j int, v float64);
	
	GetRow(i int) []float64;
	GetCol(j int) []float64;
	GetDiagonal() []float64;
	BufferRow(i int, buf []float64);
	BufferCol(j int, buf []float64);
	BufferDiagonal(buf []float64);
	
	StackHorizontal(B Matrix) Matrix;
	StackVertical(B Matrix) Matrix;
	
	GetMatrix(i int, j int, rows int, cols int) Matrix;
	GetColVector(i int) Matrix;
	GetRowVector(j int) Matrix;
	
	Plus(B Matrix) Matrix;
	Times(B Matrix) Matrix;
	Scale(f float64) Matrix;
	
	Transpose() Matrix;
	Inverse() Matrix;
	Cholesky() Matrix;
	LU() (Matrix, Matrix, Matrix);
	
	L() Matrix;
	U() Matrix;
	
	Det() float64;
	Trace() float64;
	
	Solve(b Matrix) Matrix;
	
	Copy() Matrix;
	
	String() string;
}

func (m *matrix) swapRows(r1 int, r2 int) {
	for i:=0; i<m.cols; i++ {
		tmp := m.elements[r1*m.cols+i];
		m.elements[r1*m.cols+i] = m.elements[r2*m.cols+i];
		m.elements[r2*m.cols+i] = tmp
	}
}

func (m *matrix) scaleRow(r int, f float64) {
	for i:=0; i<m.cols; i++ {
		m.elements[r*m.cols+i] *= f
	}
}

func (m *matrix) scaleAddRow(rd int, rs int, f float64) {
	for i:=0; i<m.cols; i++ {
		m.elements[rd*m.cols+i] += m.elements[rs*m.cols+i]*f
	}
}

func (A *matrix) StackHorizontal(B Matrix) Matrix {
	if A.Rows() != B.Rows() {
		return nil
	}
	C := zeros(A.Rows(), A.Cols()+B.Cols());
	for i:=0; i<C.Rows(); i++ {
		for j:=0; j<A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j))
		}
		for j:=0; j<B.Cols(); j++ {
			C.Set(i, j+A.Cols(), B.Get(i, j))
		}
	}
	return C
}

func (A *matrix) StackVertical(B Matrix) Matrix {
	if A.Cols() != B.Cols() {
		return nil
	}
	C := zeros(A.Rows()+B.Rows(), A.Cols());
	for j:=0; j<A.Cols(); j++ {
		for i:=0; i<A.Rows(); i++ {
			C.Set(i, j, A.Get(i, j))
		}
		for i:=0; i<B.Rows(); i++ {
			C.Set(i, j+A.Rows(), B.Get(i, j))
		}
	}
	return C
}

func (A *matrix) getMatrix(i int, j int, rows int, cols int) *matrix {
	B := zeros(rows, cols);
	for y:=0; y<rows; y++ {
		for x:=0; x<cols; x++ {
			B.Set(y, x, A.Get(y+i, x+j))
		}
	}
	return B
}

func (A *matrix) GetMatrix(i int, j int, rows int, cols int) Matrix {
	return A.getMatrix(i, j, rows, cols)
}
func (A *matrix) GetColVector(j int) Matrix {
	return A.GetMatrix(0, j, A.rows, j+1)
}

func (A *matrix) GetRowVector(i int) Matrix {
	return A.GetMatrix(i, 0, i+1, A.cols)
}

func (A *matrix) Plus(B Matrix) Matrix {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return nil
	}

	C := zeros(A.Rows(), A.Cols());
	
	for i:=0; i<A.Rows(); i++ {
		for j:=0; j<A.Cols(); j++ {
			C.Set(i, j, A.Get(i, j)+B.Get(i, j));
		}
	}
	return C;
}

func (A *matrix) Times(B Matrix) Matrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	C := zeros(A.Rows(), B.Cols());
	for i:=0; i<A.Rows(); i++ {
		for j:=0; j<B.Cols(); j++ {
			sum := float64(0);
			for k:=0; k<A.Cols(); k++ {
				sum += A.Get(i, k)*B.Get(k, j)
			}
			C.Set(i, j, sum)
		}
	}
	return C
}
func (A *matrix) Scale(f float64) Matrix {
	B := zeros(A.Rows(), A.Cols());
	for i:=0; i<A.Rows(); i++ {
		for j:=0; j<A.Cols(); j++ {
			B.Set(i, j, f*A.Get(i, j))
		}
	}
	return B;
}

func (A *matrix) Transpose() Matrix {
	B := zeros(A.Cols(), A.Rows());
	for i:=0; i<A.Rows(); i++ {
		for j:=0; j<A.Cols(); j++ {
			B.Set(j, i, A.Get(i, j))
		}
	}
	return B;
}

func (A *matrix) L() Matrix {
	B := zeros(A.rows, A.cols);
	for i:=0; i<A.rows; i++ {
		for j:=0; j<=i; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	B.matrixType = lower;
	return B
}

func (A *matrix) U() Matrix {
	B := zeros(A.rows, A.cols);
	for i:=0; i<A.rows; i++ {
		for j:=i; j<A.cols; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	B.matrixType = upper;
	return B
}

func (A *matrix) Copy() Matrix {
	return MakeMatrixFlat(A.elements, A.rows, A.cols)
}

func zeros(rows int, cols int) *matrix {
	A := new(matrix);
	A.elements = make([]float64, rows*cols);
	A.rows = rows;
	A.cols = cols;
	return A;
}

func Zeros(rows int, cols int) Matrix	{ return zeros(rows, cols) }

func eye(span int) *matrix {
	A := zeros(span, span);
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A;
}

func Eye(span int) Matrix { return eye(span) }

func (A *matrix) String() string {
	s := "[";
	for i:=0; i<A.rows; i++ {
		for j:=0; j<A.cols; j++ {
			s += fmt.Sprintf("%f", A.Get(i, j));
			if j!=A.cols-1 {
				s += " "
			}
		}
		if i!=A.rows-1 {
			s += "; "
		}
	}
	s += "]";
	return s
}

