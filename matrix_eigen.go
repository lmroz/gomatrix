package matrix

import "math"

//returns D,V st V*D*V' = A and D is diagonal
//code translated/ripped off from Jama
func (A *matrix) Eigen() (Matrix, Matrix) {
	if A.Symmetric() {
		n := A.rows;
		V := A.Copy().Arrays();
		d := make([]float64, n);
		e := make([]float64, n);

		tred2(V[0:n], d[0:n], e[0:n]);	//pass slices so they're references
		tql2(V[0:n], d[0:n], e[0:n]);

		D := zeros(n, n);
		for i := 0; i < n; i++ {
			D.Set(i, i, d[i])
		}

		return D, MakeMatrix(V);
	} else {
		//other stuff
	}
	return nil, nil;
}

func tred2(V [][]float64, d []float64, e []float64) {
	n := len(V);
	for j := 0; j < n; j++ {
		d[j] = V[n-1][j]
	}
	for i := n - 1; i > 0; i-- {
		scale := float64(0);
		h := float64(0);
		for k := 0; k < i; k++ {
			scale += math.Fabs(d[k])
		}
		if scale == 0 {
			e[i] = d[i-1];
			for j := 0; j < i; j++ {
				d[j] = V[i-1][j];
				V[i][j] = 0;
				V[j][i] = 0;
			}
		} else {
			for k := 0; k < i; k++ {
				d[k] /= scale;
				h += d[k] * d[k];
			}
			f := d[i-1];
			g := math.Sqrt(h);
			if f > 0 {
				g = -g
			}
			e[i] = scale * g;
			h -= f * g;
			d[i-1] = f - g;
			for j := 0; j < i; j++ {
				e[j] = 0
			}

			for j := 0; j < i; j++ {
				f = d[j];
				V[j][i] = f;
				g = e[j] + V[j][j]*f;
				for k := j + 1; k <= i-1; k++ {
					g += V[k][j] * d[k];
					e[k] += V[k][k] * f;
				}
				e[j] = g;
			}
			f = 0.0;
			for j := 0; j < i; j++ {
				e[j] /= h;
				f += e[j] * d[j];
			}
			hh := f / (h + h);
			for j := 0; j < i; j++ {
				e[j] -= hh * d[j]
			}
			for j := 0; j < i; j++ {
				f = d[j];
				g = e[j];
				for k := j; k <= i-1; k++ {
					V[k][j] -= (f*e[k] + g*d[k])
				}
				d[j] = V[i-1][j];
				V[i][j] = 0.0;
			}
		}
		d[i] = h;
	}

	for i := 0; i < n-1; i++ {
		V[n-1][i] = V[i][i];
		V[i][i] = 1.0;
		h := d[i+1];
		if h != 0.0 {
			for k := 0; k <= i; k++ {
				d[k] = V[k][i+1] / h
			}
			for j := 0; j <= i; j++ {
				g := float64(0);
				for k := 0; k <= i; k++ {
					g += V[k][i+1] * V[k][j]
				}
				for k := 0; k <= i; k++ {
					V[k][j] -= g * d[k]
				}
			}
		}
		for k := 0; k <= i; k++ {
			V[k][i+1] = 0.0
		}
	}
	for j := 0; j < n; j++ {
		d[j] = V[n-1][j];
		V[n-1][j] = 0.0;
	}
	V[n-1][n-1] = 1.0;
	e[0] = 0.0;

}


func tql2(V [][]float64, d []float64, e []float64) {

	//  This is derived from the Algol procedures tql2, by
	//  Bowdler, Martin, Reinsch, and Wilkinson, Handbook for
	//  Auto. Comp., Vol.ii-Linear Algebra, and the corresponding
	//  Fortran subroutine in EISPACK.

	n := len(V);

	for i := 1; i < n; i++ {
		e[i-1] = e[i]
	}
	e[n-1] = 0.0;

	f := float64(0);
	tst1 := float64(0);
	eps := math.Pow(2.0, -52.0);
	for l := 0; l < n; l++ {

		// Find small subdiagonal element

		tst1 = max(tst1, math.Fabs(d[l])+math.Fabs(e[l]));
		m := l;
		for m < n {
			if math.Fabs(e[m]) <= eps*tst1 {
				break
			}
			m++;
		}

		// If m == l, d[l] is an eigenvalue,
		// otherwise, iterate.

		if m > l {
			iter := 0;
			for true {
				iter = iter + 1;	// (Could check iteration count here.)

				// Compute implicit shift

				g := d[l];
				p := (d[l+1] - g) / (2.0 * e[l]);
				r := math.Sqrt(p*p + 1.0);
				if p < 0 {
					r = -r
				}
				d[l] = e[l] / (p + r);
				d[l+1] = e[l] * (p + r);
				dl1 := d[l+1];
				h := g - d[l];
				for i := l + 2; i < n; i++ {
					d[i] -= h
				}
				f = f + h;

				// Implicit QL transformation.

				p = d[m];
				c := float64(1);
				c2 := c;
				c3 := c;
				el1 := e[l+1];
				s := float64(0);
				s2 := float64(0);
				for i := m - 1; i >= l; i-- {
					c3 = c2;
					c2 = c;
					s2 = s;
					g = c * e[i];
					h = c * p;
					r = math.Sqrt(p*p + e[i]*e[i]);
					e[i+1] = s * r;
					s = e[i] / r;
					c = p / r;
					p = c*d[i] - s*g;
					d[i+1] = h + s*(c*g+s*d[i]);

					// Accumulate transformation.

					for k := 0; k < n; k++ {
						h = V[k][i+1];
						V[k][i+1] = s*V[k][i] + c*h;
						V[k][i] = c*V[k][i] - s*h;
					}
				}
				p = -s * s2 * c3 * el1 * e[l] / dl1;
				e[l] = s * p;
				d[l] = c * p;

				// Check for convergence.
				if !(math.Fabs(e[l]) > eps*tst1) {
					break
				}
			}
		}
		d[l] = d[l] + f;
		e[l] = 0.0;
	}

	// Sort eigenvalues and corresponding vectors.

	for i := 0; i < n-1; i++ {
		k := i;
		p := d[i];
		for j := i + 1; j < n; j++ {
			if d[j] < p {
				k = j;
				p = d[j];
			}
		}
		if k != i {
			d[k] = d[i];
			d[i] = p;
			for j := 0; j < n; j++ {
				p = V[j][i];
				V[j][i] = V[j][k];
				V[j][k] = p;
			}
		}
	}
}
