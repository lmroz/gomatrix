package main

import (
	"fmt";
	. "matrix";
)

func main() {
	A := Zeros(3, 3);
	fmt.Printf("You can make a 3x3 matrix of zeros:\n%v\n\n", A);

	A = Normals(3, 3);
	fmt.Printf("Or one filled with random data, A:\n%v\n\n", A);

	v := A.Get(1, 2);
	fmt.Printf("You can access a particular element A_{1,2}: %f\n\n", v);

	A.Set(1, 2, 5);
	fmt.Printf("Or change it to a different scalar:\n%v\n\n", A);

	A.SwapRows(0, 2);
	fmt.Printf("Swapping rows is easy:\n%v\n\n", A);

	Ainv, _ := A.Inverse();
	fmt.Printf("You can find A's inverse:\n%v\n\n", Ainv);

	b := Normals(3, 1);
	fmt.Printf("And a column vector, b:\n%v\n\n", b);

	x, _ := A.Solve(b);
	fmt.Printf("And find x st Ax=b:\n%v\n\n", x);

	V, D, _ := A.Eigen();
	fmt.Printf("Perhaps you want its eigenvectors and eigenvalues V,D st VDV'=A:\n%v\n%v\n\n", V, D);

	B, _ := A.TimesDense(A.Transpose());
	fmt.Printf("We can make a symmetric matrix by finding B st B=A*A':\n%v\n\n", B);

	C, _ := B.Cholesky();
	fmt.Printf("And find its cholesky decomposition C st CC'=B:\n%v\n\n", C);

	Q, R := A.QR();
	fmt.Printf("We can find QR decompositions Q,R st QR=A:\n%v\n%v\n\n", Q, R);

	L, U, P := A.LU();
	fmt.Printf("Or LU decompositions L,U,P st PLU=A:\n%v\n%v\n%v\n\n", L, U, P);

	fmt.Printf("Or A's trace (%f) and determinant (%f)\n\n", A.Trace(), A.Det());
}
