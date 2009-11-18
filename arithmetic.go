package matrix

import "math"

func (A *matrix) Equals(B Matrix) bool {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return false
	}
	
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			if A.Get(i, j) != B.Get(i, j) {
				return false
			}
		}
	}
	return true;
}

func (A *matrix) Approximates(B Matrix, ε float64) bool {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return false
	}
	
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			if math.Fabs(A.Get(i, j) - B.Get(i, j)) > ε {
				return false
			}
		}
	}
	return true;
}

func Sum(A Matrix, B Matrix) Matrix {
	if A.Cols() != B.Cols() || A.Rows() != B.Rows() {
		return nil
	}

	C := A.Copy();
	C.Add(B);
	return C;
}
func (A *matrix) Plus(B Matrix) Matrix {
	return Sum(A, B)
}

func Difference(A Matrix, B Matrix) Matrix {
	if A.Cols() != B.Cols() || A.Rows() != B.Rows() {
		return nil
	}

	C := A.Copy();
	C.Subtract(B);
	return C;
}
func (A *matrix) Minus(B Matrix) Matrix {
	return Difference(A, B)
}

func (A *matrix) Add(B Matrix) {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, A.Get(i, j)+B.Get(i, j))
		}
	}
}

func (A *matrix) Subtract(B Matrix) {
	if A.cols != B.Cols() || A.rows != B.Rows() {
		return
	}

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			A.Set(i, j, A.Get(i, j)-B.Get(i, j))
		}
	}
}

func Product(A Matrix, B Matrix) Matrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	C := zeros(A.Rows(), B.Cols());

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
func (A *matrix) Times(B Matrix) Matrix {
	return Product(A, B)
}

func ParallelProduct(A Matrix, B Matrix, threads int) Matrix {
	if A.Cols() != B.Rows() {
		return nil
	}

	C := zeros(A.Rows(), B.Cols());

	in := make(chan int);
	quit := make(chan bool);
	finish := make(chan bool);

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
				finish <- true;
				return
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

	for i := 0; i < threads; i++ {
		<- finish
	}

	return C;
}

func (A *matrix) ElementMult(B Matrix) Matrix {
	if A.rows != B.Rows() || A.cols != B.Cols() {
		return nil
	}
	C := zeros(A.rows, A.cols);
	for i := 0; i < C.rows; i++ {
		for j := 0; j < C.cols; j++ {
			C.Set(i, j, A.Get(i, j)*B.Get(i, j))
		}
	}
	return C;
}

func Scaled(A Matrix, f float64) Matrix {
	B := A.Copy();
	B.Scale(f);
	return B;
}

func (A *matrix) Scale(f float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.Set(i, j, A.Get(i, j)*f)
		}
	}
}
