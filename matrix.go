package matrix

import (
	"fmt";
	"math";
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

func (m *matrix) contains(i int, j int) bool {
	return !(i < 0 || j < 0 || i >= m.rows || j >= m.cols)
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

func (A *matrix) Elements() []float64 { return A.elements }

func (A *matrix) Rows() int	{ return A.rows }

func (A *matrix) Cols() int	{ return A.cols }

func (A *matrix) Get(i int, j int) float64 {
	return A.elements[i*A.cols+j];
}

func (A *matrix) Set(i int, j int, v float64) {
	if A.matrixType == lower && j>i || A.matrixType == upper && j<=i || A.matrixType == pivot {
		A.matrixType = 0;
	}
	A.elements[i*A.cols+j] = v;
}

func (A *matrix) GetRow(i int) []float64 {
	return A.elements[i*A.cols:(i+1)*A.cols]
}

func (A *matrix) GetCol(j int) []float64 {
	col := make([]float64, A.rows);
	for i:=0; i<A.rows; i++ {
		col[i] = A.Get(i, j)
	}
	return col
}

func (A *matrix) GetDiagonal() []float64 {
	span := A.rows;
	if A.cols < span {
		span = A.cols
	}
	diag := make([]float64, span);
	for i:=0; i<span; i++ {
		diag[i] = A.Get(i, i)
	}
	return diag
}

func (A *matrix) BufferRow(i int, buf []float64) {
	for j:=0; j<A.cols; j++ {
		buf[j] = A.Get(i, j)
	}
}

func (A *matrix) BufferCol(j int, buf []float64) {
	for i:=0; i<A.rows; i++ {
		buf[i] = A.Get(i, j)
	}
}

func (A *matrix) BufferDiagonal(buf []float64) {
	for i:=0; i<A.rows && i<A.cols; i++ {
		buf[i] = A.Get(i, i)
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

func (A *matrix) Inverse() Matrix {
	if A.Rows() != A.Cols() {
		return nil
	}
	aug := A.StackHorizontal(Eye(A.Rows()));
	for i:=0; i<aug.Rows(); i++ {
		j:=i;
		for k:=i; k<aug.Rows(); k++ {
			if math.Fabs(aug.Get(k, i)) > math.Fabs(aug.Get(j, i)) {
				j = k
			}
		}
		if j != i {
			aug.swapRows(i, j)
		}
		if aug.Get(i, i) == 0 {
			//no inverse
			return nil
		}
		aug.scaleRow(i, 1.0/aug.Get(i, i));
		for k:=i+1; k<aug.Rows(); k++ {
			aug.scaleAddRow(k, i, -aug.Get(k, i))
		}
	}
	return aug.GetMatrix(0, A.Cols(), A.Rows(), A.Cols())
}

func max(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
func min(x float64, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func (A *matrix) Cholesky() Matrix {
	n := A.rows;
	L := zeros(n, n);
	isspd := A.cols == n;
	
	for j:=0; j<n; j++ {
		Lrowj := L.GetRow(j);
		d := float64(0);
		for k:=0; k<j; k++ {
			Lrowk := L.GetRow(k);
			s := float64(0);
			for i:=0; i<k; i++ {
				s += Lrowk[i]*Lrowj[i]
			}
			s = (A.Get(j, k)-s)/Lrowk[k];
			Lrowj[k] = s;
			L.Set(j, k, s);
			d += s*s;
			isspd = isspd && (A.Get(k, j) == A.Get(j, k))
		}
		d = A.Get(j, j) - d;
		isspd = isspd && (d > 0.0);
		L.Set(j, j, math.Sqrt(max(d, float64(0))));
		for k:=j+1; k<n; k++ {
			L.Set(j, k, 0)
		}
	}
	
	return L
}

func (A *matrix) LU() (Matrix, Matrix, Matrix) {
	m := A.rows;
	n := A.cols;
	LU := A.getMatrix(0, 0, m, n);
	piv := make([]int, m);
	for i:=0; i<m; i++ {
		piv[i] = i
	}
	pivsign := float64(1.0);
	LUcolj := make([]float64, m);
	LUrowi := make([]float64, n);
	
	for j:=0; j<n; j++ {
		LU.BufferCol(j, LUcolj);
		for i:=0; i<m; i++ {
			LU.BufferRow(i, LUrowi);
			kmax := i;
			if j<i {
				kmax = j
			}
			s := float64(0);
			for k:=0; k<kmax; k++ {
				s += LUrowi[k]*LUcolj[k]
			}
			LUcolj[i] -= s;
			LUrowi[j] = LUcolj[i];
			LU.Set(i, j, LUrowi[j])
		}
		
		p := j;
		for i:=j+1; i<m; i++ {
			if math.Fabs(LUcolj[i]) > math.Fabs(LUcolj[p]) {
				p = i
			}
		}
		if false && p != j {
			LU.swapRows(p, j);
			k := piv[p];
			piv[p] = piv[j];
			piv[j] = k;
			pivsign = -pivsign;
		}
		
		if j<m && LU.elements[j*n+j] != 0.0 {
			for i:=j+1; i<m; i++ {
				LU.elements[i*n+j] /= LU.elements[j*n+j]
			}
		}
	}
	
	P := zeros(LU.rows, LU.cols);
	for i:=0; i<LU.rows; i++ {
		P.Set(piv[i], i, 1)
	}
	P.matrixType = pivot;
	P.pivotSign = pivsign;
	
	L := LU.L();
	for i:=0; i<m; i++ {
		L.Set(i, i, 1)
	}
	U := LU.U();
	
	
	return L, U, P
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

func (A *matrix) Det() float64 {
	if A.matrixType == upper || A.matrixType == lower {
		result := float64(1);
		diag := A.GetDiagonal();
		for i:=0; i<len(diag); i++ {
			result *= diag[i]
		}
		return result
	}
	else if A.matrixType == pivot {
		return A.pivotSign
	}
	_,U,P := A.LU();
	return U.Det()*P.Det()
}

func (A *matrix) Trace() (r float64) {
	for i:=0; i<A.rows; i++ {
		r += A.elements[i*A.cols+i]
	}
	return
}

func (A *matrix) Copy() Matrix {
	return MakeMatrixFlat(A.elements, A.rows, A.cols)
}

func (A *matrix) Solve(b Matrix) Matrix {
	if A.matrixType == lower {
		x := make([]float64, A.cols);
		for i:=0; i<A.rows; i++ {
			x[i] = b.Get(i,0);
			for j:=0; j<i; j++ {
				x[i] -= x[j]*A.Get(i, j)
			}
			x[i] /= A.Get(i, i)
		}
		return MakeMatrixFlat(x, A.cols, 1)
	}
	
	if A.matrixType == upper {
		x := make([]float64, A.cols);
		for i:=A.rows-1; i>=0; i-- {
			x[i] = b.Get(i,0);
			for j:=i+1; j<A.cols; j++ {
				x[i] -= x[j]*A.Get(i, j)
			}
			x[i] /= A.Get(i, i)
		}
		return MakeMatrixFlat(x, A.cols, 1)
	}
	
	L,U,P := A.LU();
	pb := P.Inverse().Times(b);
	y := L.Solve(pb);
	x := U.Solve(y);
	return x
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

func MakeMatrixFlat(elements []float64, rows int, cols int) Matrix {
	A := new(matrix);
	A.elements = elements[0:len(elements)];
	A.rows = rows;
	A.cols = cols;
	return A
}

func MakeMatrix(data [][]float64) Matrix {
	rows := len(data);
	cols := len(data[0]);
	elements := make([]float64, rows*cols);
	for i:=0; i<rows; i++ {
		for j:=0; j<cols; j++ {
			elements[i*cols+j] = data[i][j]
		}
	}
	return MakeMatrixFlat(elements, rows, cols)
}

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

