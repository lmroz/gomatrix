
package matrix


type SparseMatrix struct {
	matrix;
	elements	map[int]float64;
	// offset to start of matrix s.t. idx = i*cols + j + offset
	offset	int;
}

func (A *SparseMatrix) Get(i int, j int) float64 {

	x, ok := A.elements[i*A.cols+j];

	if !ok {
		return 0
	}
	
	return x;
}

// v == 0 results in removal of key from underlying map
func (A *SparseMatrix) Set(i int, j int, v float64) {
	A.elements[i*A.cols+j] = v, v == 0
}

func ZerosSparse(rows int, cols int) *SparseMatrix {
	A := new(SparseMatrix);
	A.rows = rows;
	A.cols = cols;
	A.offset = 0;//? not sure how offset fits in. for referencing, presumably, but wouldn't you need an offset for each dimension?
	return A;
}

func (A *SparseMatrix) DenseMatrix() *DenseMatrix {
	B := Zeros(A.rows, A.cols);
	//TODO: don't off the top of my head know how to iterate a map
	return B;
}