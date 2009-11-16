package matrix

import (
	//"fmt";
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
		{48, 14, -56, -46,
		66, -21, -10, -108,
		-240, 68, 101, 356,
		114, -122, -56, -203}, 4, 4);
		
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
    	{0.6, -0.2, -0.4, 0.4,
    	30, -30, -60, 10,
    	-1200, 800, 2100, -800,
    	-6000, 0, -10000, 7000}, 4, 4);

    if !C.Equals(Ctrue) {
    	t.Fail()
    }
}

func TestSolve(t *testing.T) {
	A := MakeMatrixFlat([]float64
		{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,}, 4, 4);
	b := MakeMatrixFlat([]float64{1, 1, 1, 1}, 4, 1);
	x := A.Solve(b);
	
	xtrue := MakeMatrixFlat([]float64{-0.906250, -3.393750, 1.275000, 1.187500}, 4, 1);
	
	if !x.Equals(xtrue) {
		t.Fail()
	}
}

func TestLU(t *testing.T) {
	A := MakeMatrixFlat([]float64
		{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,},
		4, 4);
	L,U,P := A.LU();
	
	Ltrue := MakeMatrixFlat([]float64
		{1, 0, 0, 0,
		0.5, 1, 0, 0,
		-2, -2, 1, 0,
		-1, 1, -2, 1},
		4, 4);
		
	Utrue := MakeMatrixFlat([]float64
		{6, -2, -4, 4,
		0, -2, -4, -1,
		0, 0, 5, -2,
		0, 0, 0, 8},
		4, 4);
		
	Ptrue := MakeMatrixFlat([]float64
		{1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1},
		4, 4);
		
	PLU := P.Times(L.Times(U));
	
	if !L.Equals(Ltrue) || !U.Equals(Utrue) || !P.Equals(Ptrue) || !A.Equals(PLU) {
		t.Fail()
	}
	
}

func TestQR(t *testing.T) {
	A := MakeMatrixFlat([]float64
		{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,},
		4, 4);
	Q,R := A.QR();
	
	Qtrue := MakeMatrixFlat([]float64
		{-0.4, 0.278610, 0.543792, -0.683130,
		-0.2, -0.358213, -0.699161, -0.585540,
		0.8, 0.437816, -0.126237, -0.390360,
		0.4, -0.776129, 0.446686, -0.195180},
		4, 4);
		
	Rtrue := MakeMatrixFlat([]float64
		{-15, 7.8, 15.6, -5.4,
		0, 4.019950, 17.990272, -8.179206,
		0, 0, -5.098049, 5.612709,
		0, 0, 0, -1.561440},
		4, 4);
	
	QR := Q.Times(R);
		
	if !Q.Equals(Qtrue) || !R.Equals(Rtrue) || !A.Equals(QR) {
		t.Fail()
	}
}
