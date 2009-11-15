package main

import (
	"fmt";
	. "matrix";
)

func main() {
	b := MakeMatrixFlat([]float64{1, 1, 1, 1}, 4, 1);
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	fmt.Printf("Hello, world!\n");
	fmt.Printf("A = %v\n\n", A);
	fmt.Printf("b = %v\n\n", b);
	x := A.Solve(b);
	fmt.Printf("x = %v\n\n", x);
	fmt.Printf("Ax = %v\n\n", A.Times(x));
	fmt.Printf("A.Inverse() = %v\n\n", A.Inverse());
	
	Q, R := A.QR();
	fmt.Printf("Q = %v\n\n", Q);
	fmt.Printf("R = %v\n\n", R);

	fmt.Printf("Q.Times(R) = %v\n\n", Q.Times(R));



	T := MakeMatrixFlat([]float64{0.1, 0.1, 0.1, 0.1,
					  10, 10, 10, 10,
					  100, 100, 100, 100,
					  1000, 1000, 1000, 1000,
								},
					4, 4);
	fmt.Printf("T = %v\n\n", T);

	fmt.Printf("A.ElementMult(T) = %v\n\n", A.ElementMult(T));
}
