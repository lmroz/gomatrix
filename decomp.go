//Copyright John Asmuth 2009

package matrix

import (
	"math";
)

func (A *denseMatrix) Cholesky() (*denseMatrix, Error) {
	n := A.Rows();
	L := Zeros(n, n);
	isspd := A.Cols() == n;

	for j := 0; j < n; j++ {
		Lrowj := L.RowCopy(j);
		d := float64(0);
		for k := 0; k < j; k++ {
			Lrowk := L.RowCopy(k);
			s := float64(0);
			for i := 0; i < k; i++ {
				s += Lrowk[i] * Lrowj[i]
			}
			s = (A.Get(j, k) - s) / Lrowk[k];
			Lrowj[k] = s;
			L.Set(j, k, s);
			d += s * s;
			isspd = isspd && (A.Get(k, j) == A.Get(j, k));
		}
		d = A.Get(j, j) - d;
		isspd = isspd && (d > 0.0);
		L.Set(j, j, math.Sqrt(max(d, float64(0))));
		for k := j + 1; k < n; k++ {
			L.Set(j, k, 0)
		}
	}

	if isspd {
		return nil, NewError(ErrorBadInput, "A.Cholesky(): A is not semi-positive definite")
	}

	return L, nil;
}


//returns L,U,P such that PLU=A. I realize that it's supposed to be LUP.
func (A *denseMatrix) LU() (*denseMatrix, *denseMatrix, *pivotMatrix) {
	m := A.Rows();
	n := A.Cols();
	LU := A.copy();

	P := LU.LUInPlace();

	L := LU.L();
	for i := 0; i < m && i < n; i++ {
		L.Set(i, i, 1)
	}
	U := LU.U();

	return L, U, P;
}

func (LU *denseMatrix) LUInPlace() *pivotMatrix {
	m := LU.Rows();
	n := LU.Cols();
	LUcolj := make([]float64, m);
	LUrowi := make([]float64, n);
	piv := make([]int, m);
	for i := 0; i < m; i++ {
		piv[i] = i
	}
	pivsign := float64(1.0);

	for j := 0; j < n; j++ {
		LU.BufferCol(j, LUcolj);
		for i := 0; i < m; i++ {
			LU.BufferRow(i, LUrowi);
			kmax := i;
			if j < i {
				kmax = j
			}
			s := float64(0);
			for k := 0; k < kmax; k++ {
				s += LUrowi[k] * LUcolj[k]
			}
			LUcolj[i] -= s;
			LUrowi[j] = LUcolj[i];
			LU.Set(i, j, LUrowi[j]);
		}

		p := j;
		for i := j + 1; i < m; i++ {
			if math.Fabs(LUcolj[i]) > math.Fabs(LUcolj[p]) {
				p = i
			}
		}
		if p != j {
			LU.SwapRows(p, j);
			k := piv[p];
			piv[p] = piv[j];
			piv[j] = k;
			pivsign = -pivsign;
		}

		if j < m && LU.Get(j, j) != 0 {
			for i := j + 1; i < m; i++ {
				LU.Set(i, j, LU.Get(i, j)/LU.Get(j, j))
			}
		}
	}

	P := PivotMatrix(piv, pivsign);

	return P;
}


func (A *denseMatrix) QR() (*denseMatrix, *denseMatrix) {
	m := A.Rows();
	n := A.Cols();
	QR := A.Copy();
	Q := Zeros(m, n);
	R := Zeros(m, n);
	i, j, k := 0, 0, 0;
	norm := float64(0.0);
	s := float64(0.0);

	for k = 0; k < n; k++ {
		norm = 0;
		for i = k; i < m; i++ {
			norm = math.Hypot(norm, QR.Get(i, k))
		}

		if norm != 0.0 {
			if QR.Get(k, k) < 0 {
				norm = -norm
			}

			for i = k; i < m; i++ {
				QR.Set(i, k, QR.Get(i, k)/norm)
			}
			QR.Set(k, k, QR.Get(k, k)+1.0);

			for j = k + 1; j < n; j++ {
				s = 0.0;
				for i = k; i < m; i++ {
					s += QR.Get(i, k) * QR.Get(i, j)
				}
				s = -s / QR.Get(k, k);
				for i = k; i < m; i++ {
					QR.Set(i, j, QR.Get(i, j)+s*QR.Get(i, k));

					if i < j {
						R.Set(i, j, QR.Get(i, j))
					}
				}

			}
		}

		R.Set(k, k, -norm);

	}

	//Q Matrix:
	i, j, k = 0, 0, 0;

	for k = n - 1; k >= 0; k-- {
		Q.Set(k, k, 1.0);
		for j = k; j < n; j++ {
			if QR.Get(k, k) != 0 {
				s = 0.0;
				for i = k; i < m; i++ {
					s += QR.Get(i, k) * Q.Get(i, j)
				}
				s = -s / QR.Get(k, k);
				for i = k; i < m; i++ {
					Q.Set(i, j, Q.Get(i, j)+s*QR.Get(i, k))
				}
			}
		}
	}

	return Q, R;
}

