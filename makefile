include $(GOROOT)/src/Make.$(GOARCH)

TARG=matrix
GOFILES=\
	matrix.go\
	arithmetic.go\
	data.go\
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
	pivot/pivot.go\
	pivot/pivot_arithmetic.go\
	pivot/pivot_basic.go\

include $(GOROOT)/src/Make.pkg
