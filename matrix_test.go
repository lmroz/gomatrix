package matrix

import (
	"fmt";
	"testing";
	"time";
	"rand";
)


const ε = 0.000001
const verbose = false

/* TEST: arithmetic.go */

func TestEquals(t *testing.T) {
	if !Ones(5, 3).Equals(Ones(5, 3)) {
		t.Fail()
	}
	if Ones(3, 5).Equals(Ones(5, 3)) {
		t.Fail()
	}
	if Zeros(3, 3).Equals(Ones(3, 3)) {
		t.Fail()
	}
}

func TestApproximates(t *testing.T) {
	A := Numbers(3, 3, 6);
	B := Numbers(3, 3, .1);
	C := Numbers(3, 3, .6);
	D, err := A.ElementMult(B);
	if err != nil && !D.Approximates(C, .000001) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	A := Normals(3, 3);
	B := Normals(3, 3);
	C, err := Sum(A, B);
	if err != nil {
		t.Fail()
	}
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < C.Cols(); j++ {
			if A.Get(i, j)+B.Get(i, j) != C.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestSubtract(t *testing.T) {
	A := Normals(3, 3);
	B := Normals(3, 3);
	C, err := Difference(A, B);
	if err != nil {
		t.Fail()
	}
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < C.Cols(); j++ {
			if A.Get(i, j)-B.Get(i, j) != C.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestProduct(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	B := MakeDenseMatrix([]float64{1, 7, -4, 4,
		3, -2, -6, 1,
		-12, 8, 1, 20,
		0, 0, -10, 3,
	},
		4, 4);

	C, err := Product(A, B);

	if err != nil {
		t.Fail()
	}

	Ctrue := MakeDenseMatrix([]float64{48, 14, -56, -46,
		66, -21, -10, -108,
		-240, 68, 101, 356,
		114, -122, -56, -203,
	},
		4, 4);

	if !C.Equals(Ctrue) {
		t.Fail()
	}

}

func TestParallelProduct(t *testing.T) {
	w := 100;
	h := 4;

	threads := 2;

	rand.Seed(time.Nanoseconds());
	A := Normals(h, w);
	B := Normals(w, h);

	var C *DenseMatrix;
	var start, end int64;

	start = time.Nanoseconds();
	Ctrue, err := Product(A, B);
	if err != nil {
		t.Fail()
	}
	end = time.Nanoseconds();
	if verbose {
		fmt.Printf("%fs for synchronous\n", float(end-start)/1000000000)
	}

	start = time.Nanoseconds();
	C, err = ParallelProduct(A, B, threads);
	if err != nil {
		t.Fail()
	}
	end = time.Nanoseconds();
	if verbose {
		fmt.Printf("%fs for parallel\n", float(end-start)/1000000000)
	}

	if !C.Equals(Ctrue) {
		t.Fail()
	}
}

func TestElementMult(t *testing.T) {

	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	T := MakeDenseMatrix([]float64{0.1, 0.1, 0.1, 0.1,
		10, 10, 10, 10,
		100, 100, 100, 100,
		1000, 1000, 1000, 1000,
	},
		4, 4);
	C, err := A.ElementMult(T);

	if err != nil {
		t.Fail()
	}

	Ctrue := MakeDenseMatrix([]float64{0.6, -0.2, -0.4, 0.4,
		30, -30, -60, 10,
		-1200, 800, 2100, -800,
		-6000, 0, -10000, 7000,
	},
		4, 4);

	if !C.Approximates(Ctrue, ε) {
		t.Fail()
	}
}

func TestScale(t *testing.T) {
	A := Normals(3, 3);
	f := float64(5.3);
	B := A.Copy();
	B.Scale(f);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j)*f != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestScaleMatrix(t *testing.T) {

}

/* TEST: basic.go */


func TestSymmetric(t *testing.T) {
	A := MakeDenseMatrix([]float64{
		6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	if A.Symmetric() {
		t.Fail()
	}
	B := MakeDenseMatrix([]float64{
		6, 3, -12, -6,
		3, -3, 8, 0,
		-12, 8, 21, -10,
		-6, 0, -10, 7,
	},
		4, 4);
	if !B.Symmetric() {
		t.Fail()
	}
}

func TestSwapRows(t *testing.T)	{}

func TestScaleRow(t *testing.T)	{}

func TestScaleAddRow(t *testing.T)	{}

func TestInverse(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	Ainv, err := A.Inverse();

	if err != nil {
		t.Fail()
	}

	AAinv, err := Product(A, Ainv);

	if err != nil {
		t.Fail()
	}

	if !Eye(A.Rows()).Approximates(AAinv, ε) {
		if verbose {
			fmt.Printf("A\n%v\n\nAinv\n%v\n\nA*Ainv\n%v\n", A, Ainv, AAinv)
		}
		t.Fail();
	}
}

func TestDet(t *testing.T)	{}

func TestTrace(t *testing.T)	{}

func TestOneNorm(t *testing.T)	{}

func TestTwoNorm(t *testing.T)	{}

func TestInfinityNorm(t *testing.T)	{}

func TestTranspose(t *testing.T)	{}

func TestSolve(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	b := MakeDenseMatrix([]float64{1, 1, 1, 1}, 4, 1);
	x, err := A.Solve(b);

	if err != nil {
		t.Fail()
	}

	xtrue := MakeDenseMatrix([]float64{-0.906250, -3.393750, 1.275000, 1.187500}, 4, 1);

	if !x.Equals(xtrue) {
		t.Fail()
	}
}

/* TEST: data.go */

func TestElements(t *testing.T)	{}

func TestArrays(t *testing.T)	{}

func TestRows(t *testing.T)	{}

func TestCols(t *testing.T)	{}

func TestGet(t *testing.T)	{}

func TestSet(t *testing.T)	{}

func TestRowCopy(t *testing.T)	{}

func TestColCopy(t *testing.T)	{}

func TestDiagonalCopy(t *testing.T)	{}

func TestBufferRow(t *testing.T)	{}

func TestBufferCol(t *testing.T)	{}

func TestBufferDiagonal(t *testing.T)	{}

func TestFillRow(t *testing.T)	{}

func TestFillCol(t *testing.T)	{}

func TestFillDiagonal(t *testing.T)	{}

func TestCopy(t *testing.T)	{}

func TestMakeDenseMatrix(t *testing.T)	{}

func TestMakeMatrixReference(t *testing.T)	{}

func TestMakeMatrix(t *testing.T)	{}

/* TEST: decomp.go */

func TestCholesky(t *testing.T)	{
	A := MakeDenseMatrix([]float64{1, 0.2, 0,
		0.2, 1, 0.5,
		0, 0.5, 1}, 3, 3);
	B, err := A.Cholesky();
	if err != nil {
		t.Fail()
	}
	if !A.Approximates(B.Times(B.Transpose()), ε) {
		t.Fail()
	}
}

func TestLU(t *testing.T) {

	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	L, U, P := A.LU();

	LU, err := Product(L, U);
	PLU, err := Product(P, LU);

	if err != nil {
		t.Fail()
	}

	if !A.Equals(PLU) {
		t.Fail()
	}

	A = MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	Ltrue, Utrue, Ptrue := A.LU();

	P = A.LUInPlace();
	L = A.L();
	U = A.U();

	for i := 0; i < L.Rows(); i++ {
		L.Set(i, i, 1)
	}

	PL, _ := Product(P, L);
	PLU2, _ := Product(PL, U);
	PLtrue, _ := Product(Ptrue, Ltrue);
	PLUtrue, _ := Product(PLtrue, Utrue);

	if !PLU2.Equals(PLUtrue) {
		t.Fail()
	}

}

func TestQR(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);
	Q, R := A.QR();

	Qtrue := MakeDenseMatrix([]float64{-0.4, 0.278610, 0.543792, -0.683130,
		-0.2, -0.358213, -0.699161, -0.585540,
		0.8, 0.437816, -0.126237, -0.390360,
		0.4, -0.776129, 0.446686, -0.195180,
	},
		4, 4);

	Rtrue := MakeDenseMatrix([]float64{-15, 7.8, 15.6, -5.4,
		0, 4.019950, 17.990272, -8.179206,
		0, 0, -5.098049, 5.612709,
		0, 0, 0, -1.561440,
	},
		4, 4);

	QR, _ := Product(Q, R);

	if !Q.Approximates(Qtrue, ε) ||
		!R.Approximates(Rtrue, ε) ||
		!A.Approximates(QR, ε) {
		t.Fail()
	}
}

/* TEST: eigen.go */

func TestEigen(t *testing.T) {
	A := MakeDenseMatrix([]float64{
		2, 1,
		1, 2,
	},
		2, 2);
	V, D := A.Eigen();

	Vinv, _ := V.Inverse();
	Aguess := V.Times(D).Times(Vinv);

	if !A.Approximates(Aguess, ε) {
		t.Fail()
	}

	B := MakeDenseMatrix([]float64{
		6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4);

	V, D = B.Eigen();

	Vinv, _ = V.Inverse();

	if !B.Approximates(V.Times(D).Times(Vinv), ε) {
		if verbose {
			fmt.Printf("B =\n%v\nV=\n%v\nD=\n%v\n", B, V, D)
		}
		t.Fail();
	}

	B = B.Times(B.Transpose());
	V, D = B.Eigen();
	Vinv, _ = V.Inverse();

	if !B.Approximates(V.Times(D).Times(Vinv), ε) {
		if verbose {
			fmt.Printf("B =\n%v\nV=\n%v\nD=\n%v\n", B, V, D)
		}
		t.Fail();
	}
}

/* TEST: matrix.go */

func TestGetMatrix(t *testing.T) {
	//TODO: wait for reference matrices
}

func TestGetColVector(t *testing.T) {
	//TODO: wait for reference matrices
}

func TestGetRowVector(t *testing.T) {
	//TODO: wait for reference matrices
}

func TestL(t *testing.T) {
	A := Normals(4, 4);
	L := A.L();
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j > i && L.Get(i, j) != 0 {
				t.Fail()
			} else if j <= i && L.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(4, 2);
	L = A.L();
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j > i && L.Get(i, j) != 0 {
				t.Fail()
			} else if j <= i && L.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(2, 4);
	L = A.L();
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j > i && L.Get(i, j) != 0 {
				t.Fail()
			} else if j <= i && L.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestU(t *testing.T) {
	A := Normals(4, 4);
	U := A.U();
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j < i && U.Get(i, j) != 0 {
				t.Fail()
			} else if j >= i && U.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(2, 4);
	U = A.U();
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j < i && U.Get(i, j) != 0 {
				t.Fail()
			} else if j >= i && U.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(4, 2);
	U = A.U();
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j < i && U.Get(i, j) != 0 {
				t.Fail()
			} else if j >= i && U.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestAugment(t *testing.T) {
	var A, B, C *DenseMatrix;
	A = Normals(4, 4);
	B = Normals(4, 4);
	C, _ = A.Augment(B);
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i, j+A.Cols()) != B.Get(i, j) {
				t.Fail()
			}
		}
	}

	A = Normals(2, 2);
	B = Normals(4, 4);
	C, err := A.Augment(B);
	if err == nil {
		t.Fail()
	}

	A = Normals(4, 4);
	B = Normals(4, 2);
	C, _ = A.Augment(B);
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i, j+A.Cols()) != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestStack(t *testing.T) {

	var A, B, C *DenseMatrix;
	A = Normals(4, 4);
	B = Normals(4, 4);
	C, _ = A.Stack(B);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i+A.Rows(), j) != B.Get(i, j) {
				t.Fail()
			}
		}
	}

	A = Normals(4, 4);
	B = Normals(4, 2);
	C, err := A.Stack(B);
	if err == nil {
		t.Fail()
	}

	A = Normals(2, 4);
	B = Normals(4, 4);
	C, err = A.Stack(B);

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i+A.Rows(), j) != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestZeros(t *testing.T) {
	A := Zeros(4, 5);
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j) != 0 {
				t.Fail()
			}
		}
	}
}


func TestNumbers(t *testing.T) {
	n := float64(1.0);
	A := Numbers(3, 3, n);
	//	fmt.Printf("%v\n\n\n",A.String());

	Atrue := MakeDenseMatrix([]float64{n, n, n,
		n, n, n,
		n, n, n,
	},
		3, 3);
	if !A.Equals(Atrue) {
		t.Fail()
	}
}

func TestOnes(t *testing.T) {

	A := Ones(4, 5);
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j) != 1 {
				t.Fail()
			}
		}
	}
}

func TestEye(t *testing.T) {

	A := Eye(4);
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if (i != j && A.Get(i, j) != 0) || (i == j && A.Get(i, j) != 1) {
				t.Fail()
			}
		}
	}
}

func TestNormals(t *testing.T) {
	//test that it's filled with random data?
	A := Normals(3, 4);
	if A.Rows() != 3 || A.Cols() != 4 {
		t.Fail()
	}
}

/* TEST: util.go */

func Test_min(t *testing.T)	{}

func Test_max(t *testing.T)	{}

func Test_sum(t *testing.T)	{}

func Test_product(t *testing.T)	{}


func TestDiagonal(t *testing.T)	{}

func TestPivotMatrix(t *testing.T)	{}

func TestString(t *testing.T)	{}

