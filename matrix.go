//Copyright John Asmuth 2009
package matrix

// Interfaces are more like contracts than classes; only the functions
// that absolutely need to be here should be here.  The interface should
// represent the set of functions on which all other matrix functions are based.

//The MatrixRO interface defines matrix operations that do not change the
//underlying data, such as information requests or the creation of transforms
type MatrixRO interface {
	//returns true if the underlying object is nil
	Nil() bool;

	Rows() int;
	Cols() int;
	// number of elements in the matrix
	NumElements() int;
	GetSize() (int, int);

	Get(i int, j int) float64;

	Det() float64;
	Trace() float64;

	//make a printable string
	String() string;
}

type Matrix interface {
	MatrixRO;

	Set(i int, j int, v float64);
}

type matrix struct {
	rows	int;
	cols	int;
}

func (A *matrix) Nil() bool	{ return A == nil }

func (A *matrix) Rows() int	{ return A.rows }

func (A *matrix) Cols() int	{ return A.cols }

func (A *matrix) NumElements() int	{ return A.rows * A.cols }

func (A *matrix) GetSize() (int, int)	{ return A.rows, A.cols }
