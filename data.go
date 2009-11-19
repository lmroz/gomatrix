//Copyright John Asmuth 2009
package matrix

type matrix struct {
	// flattened matrix data. elements[i*step+j] is row i, col j
	elements	[]float64;
	// the number of rows
	rows	int;
	// the number of columns
	cols	int;
	// actual offset between rows
	step	int;
}

func (A *matrix) ErrorCode() int {
	if A == nil {
		return ErrorNilMatrix
	}
	return 0;
}

func (A *matrix) ErrorString() string {
	if A == nil {
		return "Matrix is nil"
	}
	return "no error";
}

//TODO: this might not make sense with reference matrices

//This returns a slice referencing the matrix data. Changes to the slice
//effect changes to the matrix
func (A *matrix) Elements() []float64	{ return A.elements[0 : A.rows*A.cols] }

//This returns an array of slices referencing the matrix data. Changes to
//the slices effect changes to the matrix
func (A *matrix) Arrays() [][]float64 {
	a := make([][]float64, A.rows);
	for i := 0; i < A.rows; i++ {
		a[i] = A.elements[i*A.step : i*A.step+A.cols]
	}
	return a;
}

func (A *matrix) Rows() int	{ return A.rows }

func (A *matrix) Cols() int	{ return A.cols }

func (A *matrix) NumElements() int	{ return A.rows * A.cols }

func (A *matrix) Get(i int, j int) float64	{ return A.elements[i*A.step+j] }

func (A *matrix) Set(i int, j int, v float64)	{ A.elements[i*A.step+j] = v }

//returns a copy of the row (not a slice)
func (A *matrix) RowCopy(i int) []float64 {
	row := make([]float64, A.cols);
	for j := 0; j < A.cols; j++ {
		row[j] = A.Get(i, j)
	}
	return row;
}

//returns a copy of the column (not a slice)
func (A *matrix) ColCopy(j int) []float64 {
	col := make([]float64, A.rows);
	for i := 0; i < A.rows; i++ {
		col[i] = A.Get(i, j)
	}
	return col;
}

//returns a copy of the diagonal (not a slice)
func (A *matrix) DiagonalCopy() []float64 {
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

func (A *matrix) FillRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		A.Set(i, j, buf[j])
	}
}

func (A *matrix) FillCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		A.Set(i, j, buf[j])
	}
}

func (A *matrix) FillDiagonal(buf []float64) {
	for i := 0; i < A.rows && i < A.cols; i++ {
		A.Set(i, i, buf[i])
	}
}

func (A *matrix) Copy() Matrix	{ return MakeMatrixFlat(A.elements, A.rows, A.cols) }

func (A *matrix) copy() Matrix {
	B := NewMatrix(A.rows, A.cols);
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.Set(i, j, A.Get(i, j))
		}
	}
	return B;
}

func MakeMatrixFlat(elements []float64, rows int, cols int) Matrix {

	A := NewMatrix(rows, cols);
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			A.Set(i, j, elements[i*cols+j])
		}
	}
	return A;
}

func MakeMatrixReference(elements []float64, rows int, cols int) Matrix {
	A := new(matrix);
	A.elements = elements;
	A.rows = rows;
	A.cols = cols;
	A.step = cols;
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
