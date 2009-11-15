//copyright John Asmuth 2009

package matrix

import (
	"math";
)

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
