package matrix

import (
	"fmt";
	"testing";
	"time";
	"rand";
)

const ε = 0.000001

func TestTimes(t *testing.T) {
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	B := MakeMatrixFlat([]float64{1, 7, -4, 4,
		3, -2, -6, 1,
		-12, 8, 1, 20,
		0, 0, -10, 3,
	},
		4, 4);

	C := A.Times(B);

	Ctrue := MakeMatrixFlat([]float64{48, 14, -56, -46,
		66, -21, -10, -108,
		-240, 68, 101, 356,
		114, -122, -56, -203,
	},
		4, 4);

	if !C.Equals(Ctrue) {
		t.Fail()
	}

}

func TestSymmetric(t *testing.T) {
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	if A.Symmetric() {
		t.Fail()
	}
	B := MakeMatrixFlat([]float64{6, 3, -12, -6,
		3, -3, 8, 0,
		-12, 8, 21, -10,
		-6, 0, -10, 7,
	},
		4, 4);
	if !B.Symmetric() {
		t.Fail()
	}
}

func TestEigen(t *testing.T) {
	A := MakeMatrixFlat([]float64{2, 1,
		1, 2,
	},
		2, 2);
	D, V := A.Eigen();

	Aguess := V.Times(D).Times(V.Transpose());

	if !A.Approximates(Aguess, ε) {
		t.Fail()
	}
}

func TestParallelTimes(t *testing.T) {
	w := 100;
	h := 10;

	threads := 1;

	A := zeros(h, w);
	B := zeros(w, h);
	rand.Seed(time.Nanoseconds());
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			A.Set(i, j, rand.NormFloat64());
			B.Set(j, i, rand.NormFloat64());
		}
	}

	var C Matrix;
	//start := time.Nanoseconds();
	C = A.ParallelTimes(B, threads);
	//end := time.Nanoseconds();
	//fmt.Printf("%fns for parallel\n", float(end-start)/1000000000);

	//start = time.Nanoseconds();
	Ctrue := A.Times(B);
	//end = time.Nanoseconds();
	//fmt.Printf("%fns for synchronous\n", float(end-start)/1000000000);

	if !C.Equals(Ctrue) {
		t.Fail()
	}

}

func TestElementMult(t *testing.T) {

	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	T := MakeMatrixFlat([]float64{0.1, 0.1, 0.1, 0.1,
		10, 10, 10, 10,
		100, 100, 100, 100,
		1000, 1000, 1000, 1000,
	},
		4, 4);
	C := A.ElementMult(T);

	Ctrue := MakeMatrixFlat([]float64{0.6, -0.2, -0.4, 0.4,
		30, -30, -60, 10,
		-1200, 800, 2100, -800,
		-6000, 0, -10000, 7000,
	},
		4, 4);

	if !C.Approximates(Ctrue, ε) {
		t.Fail()
	}
}

func TestSolve(t *testing.T) {
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	b := MakeMatrixFlat([]float64{1, 1, 1, 1}, 4, 1);
	x := A.Solve(b);

	xtrue := MakeMatrixFlat([]float64{-0.906250, -3.393750, 1.275000, 1.187500}, 4, 1);

	if !x.Equals(xtrue) {
		t.Fail()
	}
}

func TestInverse(t *testing.T) {
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	Ainv := A.Inverse();

	Ainvtrue := MakeMatrixFlat([]float64{-0.114583, -0.479167, -0.208333, -0.104167,
		-0.718750, -1.787500, -0.725000, -0.162500,
		0.375000, 0.550000, 0.300000, 0.050000,
		0.437500, 0.375000, 0.250000, 0.125000,
	},
		4, 4);

	if !Ainv.Approximates(Ainvtrue, ε) {
		t.Fail()
	}
}

func TestLU(t *testing.T) {
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	L, U, P := A.LU();

	Ltrue := MakeMatrixFlat([]float64{1, 0, 0, 0,
		0.5, 1, 0, 0,
		-2, -2, 1, 0,
		-1, 1, -2, 1,
	},
		4, 4);

	Utrue := MakeMatrixFlat([]float64{6, -2, -4, 4,
		0, -2, -4, -1,
		0, 0, 5, -2,
		0, 0, 0, 8,
	},
		4, 4);

	Ptrue := MakeMatrixFlat([]float64{1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	},
		4, 4);

	PLU := P.Times(L.Times(U));

	if !L.Equals(Ltrue) || !U.Equals(Utrue) || !P.Equals(Ptrue) || !A.Equals(PLU) {
		t.Fail()
	}

}

func TestQR(t *testing.T) {
	A := MakeMatrixFlat([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	Q, R := A.QR();

	Qtrue := MakeMatrixFlat([]float64{-0.4, 0.278610, 0.543792, -0.683130,
		-0.2, -0.358213, -0.699161, -0.585540,
		0.8, 0.437816, -0.126237, -0.390360,
		0.4, -0.776129, 0.446686, -0.195180,
	},
		4, 4);

	Rtrue := MakeMatrixFlat([]float64{-15, 7.8, 15.6, -5.4,
		0, 4.019950, 17.990272, -8.179206,
		0, 0, -5.098049, 5.612709,
		0, 0, 0, -1.561440,
	},
		4, 4);

	QR := Q.Times(R);

	if !Q.Approximates(Qtrue, ε) ||
		!R.Approximates(Rtrue, ε) ||
		!A.Approximates(QR, ε) {
		t.Fail()
	}
}
