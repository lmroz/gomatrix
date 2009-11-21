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
	
	//copyMatrix() Matrix;//a bug has been accepted to address the inability to do this
	
	//this might not be a useful function
	isReadOnly() bool;
	
	//make a printable string
	String() string;
}

type Matrix interface {
	MatrixRO;

	//just here to do a v-table lookup for Copy(Matrix). will be moved to MatrixRO
	copyMatrix() Matrix;

	//arithmetic
	Add(B Matrix) *error;
	Subtract(B Matrix) *error;
	Scale(f float64);

	//	SwapRows(i1 int, i2 int);
	//	ScaleRow(i int, f float64);
	//	ScaleAddRow(i1 int, i2 int, f float64);
	//return x such that this*x = b
	//	Solve(b Matrix) Matrix;

	//	Transpose() Matrix;
	//	TransposeInPlace();
	//	Inverse() Matrix;

	//	OneNorm() float64;
	//	TwoNorm() float64;
	//	InfinityNorm() float64;

	Set(i int, j int, v float64);
}

func Copy(A Matrix) Matrix {
	return A.copyMatrix()
}


