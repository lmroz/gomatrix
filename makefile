include $(GOROOT)/src/Make.$(GOARCH)

TARG=matrix
GOFILES=\
	matrix.go\
	arithmetic.go\
	error.go\
	util.go\
	dense/dense.go\
	dense/dense_arithmetic.go\
	dense/dense_basic.go\
	dense/dense_data.go\
	dense/dense_decomp.go\
	dense/dense_eigen.go\
	sparse/sparse.go\
	sparse/sparse_arithmetic.go\
	sparse/sparse_basic.go\
	pivot/pivot.go\
	pivot/pivot_arithmetic.go\
	pivot/pivot_basic.go\

include $(GOROOT)/src/Make.pkg

examples: examples.6
	6l -o examples examples.6
	
examples.6: examples.go install
	6g examples.go
