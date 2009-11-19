include $(GOROOT)/src/Make.$(GOARCH)

TARG=matrix
GOFILES=\
	matrix.go\
	arithmetic.go\
	basic.go\
	decomp.go\
	eigen.go\
	data.go\
	pivot.go\
	error.go\
	util.go\

include $(GOROOT)/src/Make.pkg
