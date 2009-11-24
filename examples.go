package main


import (
	"fmt";
	. "./_obj/matrix";
)

func main() {
	A := Zeros(3, 3);
	fmt.Printf("You can make a 3x3 matrix of zeros:\n%v\n\n", A);
	
	A = Normals(3, 3);
	fmt.Printf("Or one filled with random data, A:\n%v\n\n", A);
	
	b := Normals(3, 1);
	fmt.Printf("And a column vector, b:\n%v\n\n", b);
	
	x, _ := A.Solve(b);
	fmt.Printf("And solve x st Ax=b:\n%v\n\n", x);
	
	V, D := A.Eigen();
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