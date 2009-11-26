// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

//The MatrixRO interface defines matrix operations that do not change the
//underlying data, such as information requests or the creation of transforms
/*
Read-only matrix types (at the moment, PivotMatrix).
*/
type MatrixRO interface {
	//Returns true if the underlying object is nil.
	Nil() bool;
	
	//The number of rows in this matrix.
	Rows() int;
	//The number of columns in this matrix.
	Cols() int;
	
	//The number of elements in this matrix.
	NumElements() int;
	//The size pair, (Rows(), Cols())
	GetSize() (int, int);

	//The element in the ith row and jth column.
	Get(i int, j int) float64;

	//The determinant of this matrix.
	Det() float64;
	//The trace of this matrix.
	Trace() float64;

	//A pretty-print string.
	String() string;
}

/*
A mutable matrix.
*/
type Matrix interface {
	MatrixRO;

	//Set the element at the ith row and jth column to v.
	Set(i int, j int, v float64);
}

type matrix struct {
	rows	int;
	cols	int;
}

func (A *matrix) Nil() bool	{ return A == nil }

func (A *matrix) Rows() int	{ return A.rows }

func (A *matrix) Cols() int	{ return A.cols }

func (A *matrix) NumElements() int	{ return A.rows * A.cols }

func (A *matrix) GetSize() (rows, cols int)	{
	rows = A.rows;
	cols = A.cols;
	return;
}
