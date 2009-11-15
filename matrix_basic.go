//Copyright 2009 John Asmuth
package matrix

import "math"

func (A *matrix) Inverse() Matrix {
	if A.Rows() != A.Cols() {
		return nil
	}
	aug := A.StackHorizontal(Eye(A.Rows()));
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
		for k := i + 1; k < aug.Rows(); k++ {
			aug.scaleAddRow(k, i, -aug.Get(k, i))
		}
	}
	return aug.GetMatrix(0, A.Cols(), A.Rows(), A.Cols());
}

func (A *matrix) Det() float64 {
	if A.matrixType == upper || A.matrixType == lower {
		result := float64(1);
		diag := A.GetDiagonal();
		for i := 0; i < len(diag); i++ {
			result *= diag[i]
		}
		return result;
	} else if A.matrixType == pivot {
		return A.pivotSign
	}
	_, U, P := A.LU();
	return U.Det() * P.Det();
}

func (A *matrix) Trace() (r float64) {
	for i := 0; i < A.rows; i++ {
		r += A.elements[i*A.cols+i]
	}
	return;
}

func (A *matrix) Solve(b Matrix) Matrix {
	if A.matrixType == lower {
		x := make([]float64, A.cols);
		for i := 0; i < A.rows; i++ {
			x[i] = b.Get(i, 0);
			for j := 0; j < i; j++ {
				x[i] -= x[j] * A.Get(i, j)
			}
			x[i] /= A.Get(i, i);
		}
		return MakeMatrixFlat(x, A.cols, 1);
	}

	if A.matrixType == upper {
		x := make([]float64, A.cols);
		for i := A.rows - 1; i >= 0; i-- {
			x[i] = b.Get(i, 0);
			for j := i + 1; j < A.cols; j++ {
				x[i] -= x[j] * A.Get(i, j)
			}
			x[i] /= A.Get(i, i);
		}
		return MakeMatrixFlat(x, A.cols, 1);
	}

	L, U, P := A.LU();
	pb := P.Inverse().Times(b);
	y := L.Solve(pb);
	x := U.Solve(y);
	return x;
}
