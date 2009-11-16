//Copyright John Asmuth 2009
package matrix

func max(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y;
}
func min(x float64, y float64) float64 {
	if x < y {
		return x
	}
	return y;
}

func (A *matrix) Elements() []float64	{ return A.elements }
func (A *matrix) Arrays() [][]float64 {
	a := make([][]float64, A.rows);
	for i:=0; i<A.rows; i++ {
		a[i] = A.elements[i*A.cols:(i+1)*A.cols]
	}
	return a
}


func (A *matrix) Rows() int	{ return A.rows }

func (A *matrix) Cols() int	{ return A.cols }

func (A *matrix) Get(i int, j int) float64	{ return A.elements[i*A.cols+j] }

func (A *matrix) Set(i int, j int, v float64) {
	if A.matrixType == lower && j > i || A.matrixType == upper && j <= i || A.matrixType == pivot {
		A.matrixType = 0
	}
	A.elements[i*A.cols+j] = v;
}

func (A *matrix) GetRow(i int) []float64	{ return A.elements[i*A.cols : (i+1)*A.cols] }

func (A *matrix) GetCol(j int) []float64 {
	col := make([]float64, A.rows);
	for i := 0; i < A.rows; i++ {
		col[i] = A.Get(i, j)
	}
	return col;
}

func (A *matrix) GetDiagonal() []float64 {
	span := A.rows;
	if A.cols < span {
		span = A.cols
	}
	diag := make([]float64, span);
	for i := 0; i < span; i++ {
		diag[i] = A.Get(i, i)
	}
	return diag;
}

func (A *matrix) BufferRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		buf[j] = A.Get(i, j)
	}
}

func (A *matrix) BufferCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		buf[i] = A.Get(i, j)
	}
}

func (A *matrix) BufferDiagonal(buf []float64) {
	for i := 0; i < A.rows && i < A.cols; i++ {
		buf[i] = A.Get(i, i)
	}
}

func (A *matrix) Copy() Matrix {
	return MakeMatrixFlat(A.elements, A.rows, A.cols)
}

func (A *matrix) copy() Matrix {
	B := new(matrix);
	B.elements = A.elements[0:len(A.elements)];
	B.rows = A.rows;
	B.cols = A.cols;
	return B;
}

func MakeMatrixFlat(elements []float64, rows int, cols int) Matrix {
	A := new(matrix);
	A.elements = elements[0:len(elements)];
	A.rows = rows;
	A.cols = cols;
	return A;
}

func MakeMatrix(data [][]float64) Matrix {
	rows := len(data);
	cols := len(data[0]);
	elements := make([]float64, rows*cols);
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elements[i*cols+j] = data[i][j]
		}
	}
	return MakeMatrixFlat(elements, rows, cols);
}
