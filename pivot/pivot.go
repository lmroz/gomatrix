package matrix

type PivotMatrix struct {
	DenseMatrix;
	//pivots		[]int;
	pivotSign	float64;
}

//for pivots we can speed this up a bit
func (A *PivotMatrix) Inverse() Matrix {
	return A.Transpose();
}

func (A *PivotMatrix) Det() float64	{ return A.pivotSign }

func MakePivotMatrix(pivots []int, pivotSign float64) *PivotMatrix {
	n := len(pivots);
	P := new(PivotMatrix);
	P.elements = make([]float64, n*n);
	P.rows = n;
	P.cols = n;
	P.step = n;
	for i := 0; i < n; i++ {
		P.Set(pivots[i], i, 1)
	}
	//P.pivots = pivots;
	P.pivotSign = pivotSign;
	return P;
}
