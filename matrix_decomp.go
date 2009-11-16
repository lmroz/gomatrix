//Copyright John Asmuth 2009

package matrix

import (
	"math";
)

func (A *matrix) Cholesky() Matrix {
	n := A.rows;
	L := zeros(n, n);
	isspd := A.cols == n;

	for j := 0; j < n; j++ {
		Lrowj := L.GetRow(j);
		d := float64(0);
		for k := 0; k < j; k++ {
			Lrowk := L.GetRow(k);
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

	return L;
}


//returns L,U,P such that PLU=A. I realize that it's supposed to be LUP.
func (A *matrix) LU() (Matrix, Matrix, Matrix) {
	m := A.rows;
	n := A.cols;
	LU := A.getMatrix(0, 0, m, n);
	piv := make([]int, m);
	for i := 0; i < m; i++ {
		piv[i] = i
	}
	pivsign := float64(1.0);
	LUcolj := make([]float64, m);
	LUrowi := make([]float64, n);

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
		if false && p != j {
			LU.swapRows(p, j);
			k := piv[p];
			piv[p] = piv[j];
			piv[j] = k;
			pivsign = -pivsign;
		}

		if j < m && LU.elements[j*n+j] != 0.0 {
			for i := j + 1; i < m; i++ {
				LU.elements[i*n+j] /= LU.elements[j*n+j]
			}
		}
	}

	P := zeros(LU.rows, LU.cols);
	for i := 0; i < LU.rows; i++ {
		P.Set(piv[i], i, 1)
	}
	P.matrixType = pivot;
	P.pivotSign = pivsign;

	L := LU.L();
	for i := 0; i < m; i++ {
		L.Set(i, i, 1)
	}
	U := LU.U();

	return L, U, P;
}


func (A *matrix) QR() (Matrix, Matrix) {
	m := A.rows;
	n := A.cols;
	QR := A.getMatrix(0, 0, m, n);
	Q := zeros(m,n);
	R := zeros(m,n);
	i,j,k := 0,0,0;
	norm := float64(0.0);
	s := float64(0.0);

       for k = 0; k < n; k++ {
          norm = 0;
          for i = k; i < m; i++ {
             norm = math.Hypot(norm,QR.Get(i,k));
          }
 
          if norm != 0.0 {
             if QR.Get(k,k) < 0 {
                norm = -norm;
             }
 
            for i = k; i < m; i++ {
                QR.Set(i,k, QR.Get(i,k)/norm);
             }
             QR.Set(k,k,QR.Get(k,k)+1.0);
 
             for j = k+1; j < n; j++ {
                s = 0.0; 
                for i = k; i < m; i++ {
                   s += QR.Get(i,k)*QR.Get(i,j);
                }
                s = -s/QR.Get(k,k);
                for i = k; i < m; i++ {
                   QR.Set(i,j,QR.Get(i,j)+s*QR.Get(i,k));

			if i < j {
				R.Set(i,j,QR.Get(i,j));
			}
                }

             }
          }
          
	R.Set(k,k,-norm);

       }

	//Q Matrix:
	i,j,k = 0,0,0;

       for k = n-1; k >= 0; k-- {
          Q.Set(k,k,1.0);
          for j = k; j < n; j++ {
             if QR.Get(k,k) != 0 {
                s = 0.0;
                for i = k; i < m; i++ {
                   s += QR.Get(i,k)*Q.Get(i,j);
                }
                s = -s/QR.Get(k,k);
                for i = k; i < m; i++ {
                   Q.Set(i,j,Q.Get(i,j) + s*QR.Get(i,k));
                }
             }
          }
       }



	return Q,R;
}

