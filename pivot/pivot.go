package matrix

type PivotMatrix struct {
	matrix;
	pivots		[]int;
	pivotSign	float64;
}

//for pivots we can speed this up a bit
func (A *PivotMatrix) Inverse() *PivotMatrix {
	return A.Transpose();
}

func (P *PivotMatrix) Transpose() *PivotMatrix {
	newPivots := make([]int, P.rows);
	for i:=0; i<P.rows; i++ {
		newPivots[P.pivots[i]] = i;
	}
	return MakePivotMatrix(newPivots, P.pivotSign);
}

func (P *PivotMatrix) Det() float64	{ return P.pivotSign }

func (P *PivotMatrix) Get(i int, j int) float64 {
	if P.pivots[j] == i {
		return 1;
	}
	return 0;
}
	
func (P *PivotMatrix) Times(A MatrixRO) (*DenseMatrix, *error) {
	if P.Cols() != A.Rows() {
		return nil, NewError(ErrorBadInput, "P.Times(A): P.Cols() != A.Rows()");
	}
	B := Zeros(P.rows, A.Cols());
	for i:=0; i<P.rows; i++ {
		k := 0;
		for ; i!=P.pivots[k]; k++ {}
		for j:=0; j<P.cols; j++ {
			B.Set(i, j, A.Get(k, j));
		}
	}
	return B, nil
}
	
func (P *PivotMatrix) Trace() (r float64) {
	for i := 0; i < len(P.pivots); i++ {
		if P.pivots[i] == i {
			r += 1;
		}
	}
	return;
}
	
func (P *PivotMatrix) isReadOnly() bool {
	return true
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
