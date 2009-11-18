//Copyright 2009 John Asmuth
package matrix

import "math"

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

func (A *matrix) Inverse() Matrix {
	if A.Rows() != A.Cols() {
		return nil
	}
	aug := augment(A, Eye(A.Rows()));
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
			//no inverse
			return nil
		}
		aug.scaleRow(i, 1.0/aug.Get(i, i));
		for k := 0; k < aug.Rows(); k++ {
			if k == i {
				continue
			}
			aug.scaleAddRow(k, i, -aug.Get(k, i));
		}
	}
	return aug.GetMatrix(0, A.Cols(), A.Rows(), A.Cols());
}

func (A *matrix) Det() float64 {
	_, U, P := A.LU();
	return product(U.GetDiagonal()) * P.Det();
}

func (A *pivotMatrix) Det() float64 {
	return A.pivotSign
}

func (A *matrix) Trace() float64 {
	return sum(A.GetDiagonal());
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

func (A *matrix) TransposeInPlace() {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			tmp := A.elements[i*A.cols+j];
			A.elements[i*A.cols+j] = A.elements[j*A.cols+i];
			A.elements[j*A.cols+i] = tmp;
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
	pb := P.Inverse().Times(b);
	y := solveLower(Acopy, pb);
	x := solveUpper(Acopy, y);
	return x;
}
