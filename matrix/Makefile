include $(GOROOT)/src/Make.inc

TARG=gomatrix.googlecode.com/hg/matrix
GOFILES=\
	arithmetic.go\
	dense.go\
	dense_arithmetic.go\
	dense_basic.go\
	dense_data.go\
	dense_decomp.go\
	dense_eigen.go\
	dense_svd.go\
	error.go\
	matrix.go\
	pivot.go\
	pivot_arithmetic.go\
	pivot_basic.go\
	settings.go\
	sparse.go\
	sparse_arithmetic.go\
	sparse_basic.go\
	util.go\

GC+= -I ../_obj
LD+= -L ../_obj

../_obj/$(TARG).a: _obj/$(TARG).a
	mkdir -p $(dir $@); cp -f $< $@
package: ../_obj/$(TARG).a

include $(GOROOT)/src/Make.pkg
