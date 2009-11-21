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
	isReadOnly() bool;
	
	//make a printable string
	String() string;
}

type Matrix interface {
	MatrixRO;


	
	//just here to do a v-table lookup for Copy(Matrix). will be moved to MatrixRO
	copyMatrix() Matrix;

	/* matrix.go */

	//get a sub matrix whose upper left corner is at i, j and has rows rows and cols cols
	//TODO: make these return reference matrices
	//	GetMatrix(i int, j int, rows int, cols int) Matrix;
	//	GetColVector(i int) Matrix;
	//	GetRowVector(j int) Matrix;


	/* arithmetic.go */

	//arithmetic
	Add(B Matrix) Error;
	Subtract(B Matrix) Error;
	Scale(f float64);

	//	Plus(B Matrix) Matrix;
	//	Minus(B Matrix) Matrix;
	//	Times(B Matrix) Matrix;
	//	ElementMult(B Matrix) Matrix;

	/* basic.go */

	//	Symmetric() bool;
	//check element-wise equality
	//	Equals(B Matrix) bool;
	//check that each element is within ε
	//	Approximates(B Matrix, ε float64) bool;

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

	/* decomp.go */

	//returns C such that C*C' = A
	//	Cholesky() Matrix;
	//returns L,U,P such that P*L*U = A
	//	LU() (Matrix, Matrix, Matrix);
	//puts [L\U] in the matrix, L's diagonal defined to be 1s. returns the pivot
	//	LUInPlace() Matrix;
	//	QR() (Matrix, Matrix);
	//returns V,D such that V*D*inv(V) = A
	//	Eigen() (Matrix, Matrix);

	/* data.go */
	Set(i int, j int, v float64);
}

func Copy(A Matrix) Matrix {
	return A.copyMatrix()
}


