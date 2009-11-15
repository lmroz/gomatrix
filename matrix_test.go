package main

import (
	"fmt";
	. "matrix"
)

func main() {
	b := MakeMatrixFlat([]float64{1, 1, 1, 1}, 4, 1);
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
								  3, -3, -6, 1,
								  -12, 8, 21, -8,
								  -6, 0, -10, 7}, 4, 4);
	fmt.Printf("Hello, world!\n");
	fmt.Printf("A = %v\n\n", A);
	fmt.Printf("b = %v\n\n", b);
	x := A.Solve(b);
	fmt.Printf("x = %v\n\n", x);
	fmt.Printf("Ax = %v\n\n", A.Times(x));
	fmt.Printf("A.Inverse() = %v\n\n", A.Inverse());
}
