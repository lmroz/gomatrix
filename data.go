//Copyright John Asmuth 2009
package matrix

import "fmt"

type matrix struct {
	rows	int;
	cols	int;
}

func (A *matrix) Nil() bool {
	return A == nil
}

func (A *matrix) Rows() int	{ return A.rows }

func (A *matrix) Cols() int	{ return A.cols }

func (A *matrix) NumElements() int	{ return A.rows * A.cols }

func (A *matrix) GetSize() (int, int)	{ return A.rows, A.cols }

func (m *matrix) isReadOnly() bool {
	return false
}


func (A *PivotMatrix) String() string {
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