package matrix

import "math"

func (P *PivotMatrix) SwapRows(r1 int, r2 int) *error {
	tmp := P.pivots[r1];
	P.pivots[r1] = P.pivots[r2];
	P.pivots[r2] = tmp;
	P.pivotSign *= -1;

	return nil;
}

func (P *PivotMatrix) Symmetric() bool {
	for i:=0; i<P.rows; i++ {
		if P.pivots[P.pivots[i]] != i {
			return false;
		}
	}
	return true;
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

func (P *PivotMatrix) Det() float64	{ 
	return P.pivotSign;
}
	
func (P *PivotMatrix) Trace() (r float64) {
	for i := 0; i < len(P.pivots); i++ {
		if P.pivots[i] == i {
			r += 1;
		}
	}
	return;
}

func (P *PivotMatrix) Solve(b MatrixRO) (*DenseMatrix, *error) {
	return P.Transpose().Times(b); //error comes from times
}


func (A *PivotMatrix) OneNorm() float64 {
	return float64(A.rows);
}
func (A *PivotMatrix) TwoNorm() float64 {
	return math.Sqrt(float64(A.rows));
}
func (A *PivotMatrix) InfinityNorm() float64 {
	return 1;
}
