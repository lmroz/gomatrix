//Copyright 2009 John Asmuth
package matrix

import "math"

func (A *matrix) Symmetric() bool {
	if A.rows != A.cols {
		return false
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < i; j++ {
			if A.Get(i, j) != A.Get(j, i) {
				return false
			}
		}
	}
	return true;
}

func (m *matrix) swapRows(r1 int, r2 int) {
	for j := 0; j < m.cols; j++ {
		tmp := m.Get(r1, j);
		m.Set(r1, j, m.Get(r2, j));
		m.Set(r2, j, tmp);
	}
}

func (m *matrix) scaleRow(r int, f float64) {
	for j := 0; j < m.cols; j++ {
		m.Set(r, j, m.Get(r, j)*f)
	}
}

func (m *matrix) scaleAddRow(rd int, rs int, f float64) {
	for j := 0; j < m.cols; j++ {
		m.Set(rd, j, m.Get(rd, j)+m.Get(rs, j)*f)
	}
}

func (A *matrix) Inverse() Matrix {
	if A.Rows() != A.Cols() {
		return Error(ErrorBadInput, "A.Inverse(): A is not square")
	}
	aug := Augment(A, Eye(A.Rows()));
	for i := 0; i < aug.Rows(); i++ {
		j := i;
		for k := i; k < aug.Rows(); k++ {
			if math.Fabs(aug.Get(k, i)) > math.Fabs(aug.Get(j, i)) {
				j = k
			}
		}
		if j != i {
			aug.swapRows(i, j)
		}
		if aug.Get(i, i) == 0 {
			return Error(ErrorBadInput, "A.Inverse(): A has no inverse")
		}
		aug.scaleRow(i, 1.0/aug.Get(i, i));
		for k := 0; k < aug.Rows(); k++ {
			if k == i {
				continue
			}
			aug.scaleAddRow(k, i, -aug.Get(k, i));
		}
	}
	inv := aug.GetMatrix(0, A.Cols(), A.Rows(), A.Cols());
	return inv;
}

func (A *matrix) Det() float64 {
	_, U, P := A.LU();
	return product(U.DiagonalCopy()) * P.Det();
}

func (A *matrix) Trace() float64	{ return sum(A.DiagonalCopy()) }

func (A *matrix) OneNorm() (ε float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			ε += A.Get(i, j)
		}
	}
	return;
}

func (A *matrix) InfinityNorm() (ε float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			ε = max(ε, A.Get(i, j))
		}
	}
	return;
}

func (A *matrix) Norm(p float64) float64 {
	sum := float64(0);
	
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			sum += math.Pow(A.Get(i, j), p);
		}
	}
	
	return math.Pow(sum, 1/p)
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

func (A *matrix) TransposeInPlace() {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			tmp := A.Get(i, j);
			A.Set(i, j, A.Get(j, i));
			A.Set(j, i, tmp);
		}
	}
}

func solveLower(A Matrix, b Matrix) Matrix {
	x := make([]float64, A.Cols());
	for i := 0; i < A.Rows(); i++ {
		x[i] = b.Get(i, 0);
		for j := 0; j < i; j++ {
			x[i] -= x[j] * A.Get(i, j)
		}
		//the diagonal defined to be ones
		//x[i] /= A.Get(i, i);
	}
	return MakeMatrixFlat(x, A.Cols(), 1);
}

func solveUpper(A Matrix, b Matrix) Matrix {
	x := make([]float64, A.Cols());
	for i := A.Rows() - 1; i >= 0; i-- {
		x[i] = b.Get(i, 0);
		for j := i + 1; j < A.Cols(); j++ {
			x[i] -= x[j] * A.Get(i, j)
		}
		x[i] /= A.Get(i, i);
	}
	return MakeMatrixFlat(x, A.Cols(), 1);
}

func (A *matrix) Solve(b Matrix) Matrix {
	Acopy := A.Copy();
	P := Acopy.LUInPlace();
	Pinv := P.Inverse();
	pb := Pinv.Times(b);
	y := solveLower(Acopy, pb);
	x := solveUpper(Acopy, y);
	return x;
}
