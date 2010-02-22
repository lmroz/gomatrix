all: demo1 demo2 demo3

demo1: demo1.6
	6l -o demo1 demo1.6
demo1.6: demo1.go
	6g demo1.go
demo2: demo2.6
	6l -o demo2 demo2.6
demo2.6: demo2.go
	6g demo2.go
demo3: demo3.6
	6l -o demo3 demo3.6
demo3.6: demo3.go
	6g demo3.go
clean:
	rm -f demo1 demo1.6 demo2 demo2.6 demo3 demo3.6
