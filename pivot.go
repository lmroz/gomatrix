package matrix

type pivotMatrix struct {
	*denseMatrix;
	//pivots		[]int;
	pivotSign	float64;
}

//for pivots we can speed this up a bit
func (A *pivotMatrix) Inverse() Matrix {
	return A.Transpose();
}

func (A *pivotMatrix) Det() float64	{ return A.pivotSign }

func PivotMatrix(pivots []int, pivotSign float64) *pivotMatrix {
	n := len(pivots);
	P := new(pivotMatrix);
	P.denseMatrix = new(denseMatrix);
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