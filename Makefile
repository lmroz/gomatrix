all:
	cd src/matrix;make
clean:
	cd src/matrix;make clean;cd ../..;\
	cd demo;make clean
install:
	cd src/matrix;make install
demo:
	cd demo;make