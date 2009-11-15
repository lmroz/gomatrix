include $(GOROOT)/src/Make.$(GOARCH)

TARG=matrix
GOFILES=\
	matrix.go\
	matrix_basic.go\
	matrix_decomp.go\
	matrix_data.go\

include $(GOROOT)/src/Make.pkg