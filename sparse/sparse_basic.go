package matrix

import "math"

func (A *SparseMatrix) SwapRows(r1 int, r2 int) {
	js := map[int]bool {};
	for index := range A.elements {
		i, j := A.getRowColIndex(index);
		if i == r1 || i == r2 {
			js[j] = true;
		}
	}
	for j := range js {
		tmp := A.Get(r1, j);
		A.Set(r1, j, A.Get(r2, j));
		A.Set(r2, j, tmp);
	}
}

func (A *SparseMatrix) ScaleRow(r int, f float64) {
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		if i == r {
			A.Set(i, j, value*f);
		}
	}
}

func (A *SparseMatrix) ScaleAddRow(rd int, rs int, f float64) {
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		if i == rs {
			A.Set(rd, j, A.Get(rd, j)+value*f);
		}
	}
}

func (A *SparseMatrix) Symmetric() bool {
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		if i != j && value != A.Get(j, i) {
			return false;
		}
	}
	return true;
}

func (A *SparseMatrix) Transpose() *SparseMatrix {
	B := ZerosSparse(A.cols, A.rows);
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		B.Set(j, i, value);
	}
	return B;
}

//TODO: this function - not sure of the best way to do this for sparse matrices
func (A *SparseMatrix) Det() float64 {
	return 0;
}

func (A *SparseMatrix) Trace() (res float64) {
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		if i == j {
			res += value;
		}
	}
	return;
}

func (A *SparseMatrix) OneNorm() (res float64) {
	for _, value := range A.elements {
		res += math.Fabs(value);	
	}
	return;
}

func (A *SparseMatrix) InfinityNorm() (res float64) {
	for _, value := range A.elements {
		res = max(res, math.Fabs(value));
	}
	return;
}
