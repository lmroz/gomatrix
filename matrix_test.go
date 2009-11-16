package matrix

import (
	"fmt";
	"testing"
)

func TestTimes(t *testing.T) {
	A := MakeMatrixFlat([]float64
		{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,}, 4, 4);
	B := MakeMatrixFlat([]float64
		{1, 7, -4, 4,
		3, -2, -6, 1,
		-12, 8, 1, 20,
		0, 0, -10, 3,}, 4, 4);
	C := A.Times(B);
	
	Ctrue := MakeMatrixFlat([]float64
		{48.000000, 14.000000, -56.000000, -46.000000,
		66.000000, -21.000000, -10.000000, -108.000000,
		-240.000000, 68.000000, 101.000000, 356.000000,
		114.000000, -122.000000, -56.000000, -203.000000}, 4, 4);
		
    if !C.Equals(Ctrue) {
    	t.Fail()
    }

}

func TestElementMult(t *testing.T) {

	A := MakeMatrixFlat([]float64
		{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,}, 4, 4);
	T := MakeMatrixFlat([]float64
		{0.1, 0.1, 0.1, 0.1,
    	10, 10, 10, 10,
    	100, 100, 100, 100,
    	1000, 1000, 1000, 1000}, 4, 4);
    C := A.ElementMult(T);
    
    Ctrue := MakeMatrixFlat([]float64
    	{0.600000, -0.200000, -0.400000, 0.400000,
    	30.000000, -30.000000, -60.000000, 10.000000,
    	-1200.000000, 800.000000, 2100.000000, -800.000000,
    	-6000.000000, 0.000000, -10000.000000, 7000.000000}, 4, 4);

	fmt.Printf("%v\n%v\n", C, Ctrue);
    if !C.Equals(Ctrue) {
    	t.Fail()
    }
}


func foo() {
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

	fmt.Printf("A.ElementMult(T) = %v\n\n", A.ElementMult(T) );
}
