package matrix

import (
	"math";
	"reflect";
	)

func Sum(A MatrixRO, B MatrixRO) Matrix {
	C := MakeDenseCopy(A);
	err := C.Add(B);
	if err.OK() {
		return C;
	}
	return nil;
}

func Difference(A Matrix, B Matrix) Matrix {
	C := MakeDenseCopy(A);
	err := C.Subtract(B);
	if err.OK() {
		return C;
	}
	return nil;
}

func Product(A MatrixRO, B MatrixRO) *DenseMatrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	C := Zeros(A.Rows(), B.Cols());

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			sum := float64(0);
			for k := 0; k < A.Cols(); k++ {
				sum += A.Get(i, k) * B.Get(k, j)
			}
			C.Set(i, j, sum);
		}
	}

	return C;
}

func ParallelProduct(A Matrix, B Matrix, threads int) *DenseMatrix {
	if A.Cols() != B.Rows() {
		return nil
	}

	C := Zeros(A.Rows(), B.Cols());

	in := make(chan int);
	quit := make(chan bool);

	dotRowCol := func() {
		for true {
			select {
			case i := <-in:
				sums := make([]float64, B.Cols());
				for k := 0; k < A.Cols(); k++ {
					for j := 0; j < B.Cols(); j++ {
						sums[j] += A.Get(i, k) * B.Get(k, j)
					}
				}
				for j := 0; j < B.Cols(); j++ {
					C.Set(i, j, sums[j])
				}
			case <-quit:
				return;
			}
		}
	};

	for i := 0; i < threads; i++ {
		go dotRowCol()
	}

	for i := 0; i < A.Rows(); i++ {
		in <- i
	}

	for i := 0; i < threads; i++ {
		quit <- true
	}

	return C;
}

func Scaled(A MatrixRO, f float64) Matrix {
	B := MakeDenseCopy(A);
	B.Scale(f);
	return B;
}

func Equals(A MatrixRO, B MatrixRO) bool {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return false;
	}
	for i:=0; i<A.Rows(); i++ {
		for j:=0; j<A.Cols(); j++ {
			if A.Get(i, j) != B.Get(i, j) {
				return false;
			}
		}
	}
	return true;
}

func ApproxEquals(A MatrixRO, B MatrixRO, ε float64) bool {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return false;
	}
	for i:=0; i<A.Rows(); i++ {
		for j:=0; j<A.Cols(); j++ {
			if math.Fabs(A.Get(i, j)-B.Get(i, j)) > ε {
				return false;
			}
		}
	}
	return true;
}

func MultipleProduct(values ...) (Matrix){
	v := reflect.NewValue(values).(*reflect.StructValue);
	if v.NumField() < 2 {
		return nil;
	}

	inter := v.Field(0).Interface();
	if C, ok := inter.(Matrix); ok {
		for i:=1; i < v.NumField(); i++ {
			inter := v.Field(i).Interface();
			if A, ok := inter.(Matrix); ok {
				C = Product(C,A);
			}
		}
		return C;
	}

	return nil;
}


