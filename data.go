//Copyright John Asmuth 2009
package matrix

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

