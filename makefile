include $(GOROOT)/src/Make.$(GOARCH)

TARG=matrix
GOFILES=\
	matrix.go\
	arithmetic.go\
	error.go\
	util.go\
	dense.go\
	dense_arithmetic.go\
	dense_basic.go\
	dense_data.go\
	dense_decomp.go\
	dense_eigen.go\
	dense_svd.go\
	sparse.go\
	sparse_arithmetic.go\
	sparse_basic.go\
	pivot.go\
	pivot_arithmetic.go\
	pivot_basic.go\

include $(GOROOT)/src/Make.pkg

src: cleansrc
	cp -r ../gomatrix $(GOROOT)/src/pkg/matrix
cleansrc:
	rm -r $(GOROOT)/src/pkg/matrix

examples: examples.6
	6l -o examples examples.6
	
examples.6: examples.go install
	6g examples.go
