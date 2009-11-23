package matrix

type PivotMatrix struct {
	matrix;
	pivots		[]int;
	pivotSign	float64;
}

func (P *PivotMatrix) Get(i int, j int) float64 {
	if P.pivots[j] == i {
		return 1;
	}
	return 0;
}

	
func (P *PivotMatrix) isReadOnly() bool {
	return true
}

func (P *PivotMatrix) DenseMatrix() *DenseMatrix {
	A := Zeros(P.rows, P.cols);
	for j:=0; j<P.rows; j++ {
		A.Set(P.pivots[j], j, 1);
	}
	return A;
}

func (P *PivotMatrix) SparseMatrix() *SparseMatrix {
	A := ZerosSparse(P.rows, P.cols);
	for j:=0; j<P.rows; j++ {
		A.Set(P.pivots[j], j, 1);
	}
	return A;
}

func (P *PivotMatrix) Copy() *PivotMatrix {
	return MakePivotMatrix(P.pivots, P.pivotSign);
}

func MakePivotMatrix(pivots []int, pivotSign float64) *PivotMatrix {
	n := len(pivots);
	P := new(PivotMatrix);
	P.rows = n;
	P.cols = n;
	P.pivots = pivots;
	P.pivotSign = pivotSign;
	return P;
}
