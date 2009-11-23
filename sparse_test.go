package matrix

import (
	"testing";
	"rand";
	//"fmt";
)

func TestGetMatrix_Sparse(t *testing.T) {
	A := ZerosSparse(6, 6);
	for i:=0; i<36; i++ {
		x := rand.Intn(6);
		y := rand.Intn(6);
		A.Set(y, x, 1);
	}
	B := A.GetMatrix(1, 1, 4, 4);
	
	for i:=0; i<4; i++ {
		for j:=0; j<4; j++ {
			if B.Get(i, j) != A.Get(i+1, j+1) {
				t.Fail();
			}
		}
	}
	
}
