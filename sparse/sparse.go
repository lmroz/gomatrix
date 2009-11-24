
package matrix

import (
	"fmt";
	"rand";
)

type SparseMatrix struct {
	matrix;
	elements	map[int]float64;
	// offset to start of matrix s.t. idx = i*cols + j + offset
	// offset = starting row * step + starting col
	offset	int;
	// analogous to dense step
	step	int;
}

func (A *SparseMatrix) Get(i int, j int) float64 {

	x, ok := A.elements[i*A.step+j+A.offset];

	if !ok {
		return 0
	}
	
	return x;
}

func (A *SparseMatrix) getRowIndex(index int) int {
	return (index - A.offset) / A.cols;
}

func (A *SparseMatrix) getColIndex(index int) int {
	return (index - A.offset) % A.cols;
}

func (A *SparseMatrix) getRowColIndex(index int) (i int, j int) {
	i = (index - A.offset) / A.step;
	j = (index - A.offset) % A.step;
	return;
}

// v == 0 results in removal of key from underlying map
func (A *SparseMatrix) Set(i int, j int, v float64) {
	A.elements[i*A.step+j+A.offset] = v, v != 0
}

//returns a channel that contains the indexes of non zero entries in this matrix
func (A *SparseMatrix) Indices() (out chan int) {
	//maybe thread the populating?
	for index := range A.elements {
		out <- index;
	}
	return;
}


func (A *SparseMatrix) GetMatrix(i int, j int, rows int, cols int) *SparseMatrix {
	B := new(SparseMatrix);
	B.rows = rows;
	B.cols = cols;
	B.offset = (i + A.offset / A.step) * A.step + (j + A.offset % A.step);
	B.step = A.step;
	B.elements = A.elements;
	return B;
}

func (A *SparseMatrix) GetColVector(j int) *SparseMatrix {
	return A.GetMatrix(0, j, A.rows, j+1)
}

func (A *SparseMatrix) GetRowVector(i int) *SparseMatrix {
	return A.GetMatrix(i, 0, i+1, A.cols)
}

func (A *SparseMatrix) Augment(B *SparseMatrix) (*SparseMatrix, *error) {
	if A.rows != B.rows {
		return nil, NewError(ErrorDimensionMismatch);
	}
	C := ZerosSparse(A.rows, A.cols+B.cols);
	
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		C.Set(i, j, value);	
	}
	
	for index, value := range B.elements {
		i, j := B.getRowColIndex(index);
		C.Set(i, j+A.cols, value);	
	}
	
	return C, nil;
}

func (A *SparseMatrix) Stack(B *SparseMatrix) (*SparseMatrix, *error) {
	if A.cols != B.cols {
		return nil, NewError(ErrorDimensionMismatch);
	}
	C := ZerosSparse(A.rows+B.rows, A.cols);
	
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		C.Set(i, j, value);	
	}
	
	for index, value := range B.elements {
		i, j := B.getRowColIndex(index);
		C.Set(i+A.rows, j, value);	
	}
	
	return C, nil;
}

func (A *SparseMatrix) L() *SparseMatrix {
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		if i >= j {
			B.Set(i, j, value);
		}
	}
	return B;
}

func (A *SparseMatrix) U() *SparseMatrix {
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		if i <= j {
			B.Set(i, j, value);
		}
	}
	return B;
}

func (A *SparseMatrix) Copy() *SparseMatrix {
	B := ZerosSparse(A.rows, A.cols);
	for index, value := range A.elements {
		B.elements[index] = value;
	}
	return B;
}

func ZerosSparse(rows int, cols int) *SparseMatrix {
	A := new(SparseMatrix);
	A.rows = rows;
	A.cols = cols;
	A.offset = 0;
	A.step = cols;
	A.elements = map[int] float64 {};
	return A;
}

func NormalsSparse(rows int, cols int, n int) *SparseMatrix {
	A := ZerosSparse(rows, cols);
	for k:=0; k<n; k++ {
		i := rand.Intn(rows);
		j := rand.Intn(cols);
		A.Set(i, j, rand.NormFloat64());
	}
	return A;
}

func MakeSparseMatrix(elements map[int]float64, rows int, cols int) *SparseMatrix {
	A := ZerosSparse(rows, cols);
	A.elements = elements;
	return A;
}

func (A *SparseMatrix) DenseMatrix() *DenseMatrix {
	B := Zeros(A.rows, A.cols);
	for index, value := range A.elements {
		i, j := A.getRowColIndex(index);
		B.Set(i, j, value);
	}
	return B;
}

func (A *SparseMatrix) String() string {

	s := "{";
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			s += fmt.Sprintf("%f", A.Get(i, j));
			if i != A.Rows()-1 || j != A.Cols()-1 {
				s += ","
			}
			if j != A.cols-1 {
				s += " "
			}
		}
		if i != A.Rows()-1 {
			s += "\n"
		}
	}
	s += "}";
	return s;
}
