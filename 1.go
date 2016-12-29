package main

import (
	"fmt"
	"math"
)

func determinant(A [][]float64) float64 {
	d := 0.0
	if len(A) == 2 && cap(A) == 2 {
		d = A[0][0]*A[1][1] - A[1][0]*A[0][1]
	} else {
		d += A[0][0] * (A[1][1]*A[2][2] - A[1][2]*A[2][1])
		d -= A[0][1] * (A[1][0]*A[2][2] - A[1][2]*A[2][0])
		d += A[0][2] * (A[1][0]*A[2][1] - A[1][1]*A[2][0])
	}
	return d
}

func CoFactor(A [][]float64) [][]float64 {
	var C [][]float64
	var resMatrice [][]float64
	for j := 0; j < len(A); j++ {
		for i := 0; i < len(A); i++ {

			/* Form the adjoint a_ij */
			i1 := 0
			for ii := 0; ii < len(A); ii++ {
				if ii == len(A) {
					continue
				}
				j1 := 0
				for jj := 0; jj < len(A); jj++ {
					if jj == len(A) {
						continue
					}
					C[i1][j1] = A[ii][jj]
					j1++
				}
				i1++
			}

			/* Fill in the elements of the cofactor */
			resMatrice[i][j] = math.Pow(-1, float64(i+j+2)) * determinant(C)
		}
	}
	return resMatrice
}
func Transpose(A [][]float64) [][]float64 {
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			tmp := A[i][j]
			A[i][j] = A[j][i]
			A[j][i] = tmp
		}
	}
	return A
}

func main() {
	A := [][]float64{{2, -4, 3}, {1, -2, 4}, {3, -1, 5}}
	//b := [3]float64{1, 3, 2}
	fmt.Println(determinant(A), "\r\n")
	fmt.Println(CoFactor(A))
	fmt.Println(A)
	fmt.Println(Transpose(A))
}
