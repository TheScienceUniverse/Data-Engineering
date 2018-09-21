package main
import "fmt"
type Num struct{
	n, d int	// numerator, denominator
}
func iniV(x *Num){
	x.n = 0; x.d = 1
}
func gcd(a, b int) int {
	var c int
	for ; ; {
		if c = b % a; c == 0{ break }
		b = a; a = c
	}
	return a;
}
func calc(a, b Num, s byte) Num {
	var n Num
	var x int
	if s == '+'{
		n.n = b.d * a.n + a.d * b.n
		n.d = a.d * b.d
	} else if s == '-' {
		n.n = b.d * a.n - a.d * b.n
		n.d = a.d * b.d
	} else if s == '*' {
		n.n = a.n * b.n
		n.d = a.d * b.d
	} else if s == '/' {
		n.n = a.n * b.d
		n.d = a.d * b.n
	}
	if n.n == 0 { n.d = 1 }
	if (n.n & n.d) != 0{
		x = gcd(n.n, n.d)
		n.n /= x; n.d /= x
	}
	if n.d < 0 {
		n.n *= -1
		n.d *= -1
	}
	return n
}
func minI(A []Num, n int) int {
	s := A[0]
	var p int = 0
	for i := 1; i < n; i++ {
		if s.n*A[i].d > s.d*A[i].n {
			s.n = A[i].n; s.d = A[i].d
			p = i
		}
	}
	return p
}
func main(){
	N := 10						// MAX_LEN(nx+nq)
	fmt.Println("Query for \"Z = CX\"...")
	var nx int
	fmt.Print("# Independant Variables: ")
	fmt.Scanf("%d", &nx)
	fmt.Println("Query for \"AX = B\"...")
	var nq int
	fmt.Print("# Inequality Equations: ")
	fmt.Scanf("%d", &nq)
//	fmt.Println(nx, "", nq);
	fmt.Println("Creating Equation \"Z = CX\"...")
	var C = make([]Num, nx+nq, N)			// C[nx+nq]
	var i int
	for i = 0; i < nx; i++ {
		fmt.Print("Coefficient of x", i, ": ")
		fmt.Scanf("%d", &C[i].n)
		C[i].d = 1
	}
	for i = nx; i < nq+nx; i++ { iniV(&C[i]) }
//	fmt.Println("C: ",C)

	fmt.Println("Creating Equation \"AX = B\"...")
	A := make([][]Num, nq)				// A[nq][nx+nq]
	B := make([]Num, nq)				// B[nq]
	for i = 0; i < nq; i++ { A[i] = make([]Num, nx+nq) }
	var j int
	for i = 0; i < nq; i++ {
		fmt.Println("Equation Number:",i)
		for j = 0; j < nx; j++ {
			fmt.Print("Coefficient of x", j, ": ")
			fmt.Scanf("%d", &A[i][j].n)
			A[i][j].d = 1
		}
		for j = 0; j < nq; j++ {
			A[i][j+nx].n = 0
			if i == j { A[i][j+nx].n = 1 }
			A[i][j+nx].d = 1
		}
		fmt.Print("Solution Value: ")
		fmt.Scanf("%d", &B[i].n)
		B[i].d = 1
	}
	fmt.Println("A: ",A,"\nB: ",B)
	fmt.Println("Creating CB...")
	var CB = make([]int, nq)			// CB[nq]
	for i = 0; i < nq; i++ { CB[i] = nx+i }
	fmt.Println("CB: ",CB)

	var d string
	Z := make([]Num, nx+nq)				// Z[nx+nq]
	i = 0
	for i = 0; i < nq+nx; i++ { iniV(&Z[i]) }
	var ev, lv int
	var s bool
	R := make([]Num, nq)				// R[nq]
	var rc int = 0

	T := make([][]Num, nq)				// T[nq][nx+nq]
	for i = 0; i < nq; i++ { T[i] = make([]Num, nx+nq) }

	for ; ; {
		for i = 0; i < nq+nx; i++ { iniV(&Z[i]) }
		for i = 0; i < nq+nx; i++ {
			for j = 0; j < nq; j++ {
				Z[i] = calc(Z[i],calc(C[CB[j]],A[j][i],'*'),'+')
			}
		}
		fmt.Println("Z: ",Z)
		for i = 0; i < nq+nx; i++ { Z[i] = calc(Z[i],C[i],'-') }
		fmt.Println("Z-C: ",Z)

		s = true			// Hope for Luck
		for i = 0; i < nq+nx; i++ {
			if Z[i].n < 0 { s = false; break }
		}

		if s == true {
			break			// Calculation Done
		} else {			// Better Luck Next Time
			ev = minI(Z, nx+nq)	// Entering Vector
			fmt.Println("eV: ",ev)
			for i = 0; i < nq; i++ { R[i] = calc(B[i], A[i][ev], '/') }
			fmt.Println("R: ",R)
			lv = minI(R, nq)	// Leaving Vector
			fmt.Println("lV: ",lv)
			CB[lv] = ev

			T = A
			for i = 0; i < nq; i++ {
				if i == lv {
					for j = 0; j < nx; j++ {
						T[lv][j] = calc(A[lv][j], A[lv][ev], '/')
					}
					B[i] = calc(B[i], A[lv][ev], '/')
				} else {
					for j = 0; j < nx; j++ {
						if j == ev {
							iniV(&T[i][j])
						} else {
							T[i][j] = calc(A[i][j], calc(calc(A[lv][j], A[i][ev], '*'), A[lv][ev], '/'), '-')
						}
					}
					B[i] = calc(B[i], calc(calc(A[i][ev], B[lv], '*'), A[lv][ev], '/'), '-')
				}
				
			}
			A = T
		}

		fmt.Println("A: ",A,"\nB:",B);
		fmt.Print("Enter Y to continue or N to stop: ")
		fmt.Scanf("%s",&d)
		rc += 1
		if (rc > 10) || (d[0] == 'N') { break }
	}
	fmt.Println("CB:",CB)
	X := make([]Num, nx)
	for i = 0; i < nx; i++ { iniV(&X[i]) }
	for i = 0; i < nq; i++ {
		if CB[i] < nx { X[CB[i]] = B[i] }
	}
	fmt.Println("X:",X)

	fmt.Println("")
}
