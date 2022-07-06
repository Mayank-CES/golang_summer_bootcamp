package main

import (
	"fmt"
)

type Matrix struct {
	rownum, colnum int
	mat            [][]float64
}

func (m Matrix) rows() int {
	return m.rownum
}

func (m Matrix) cols() int {
	return m.colnum
}

func (m *Matrix) setval(i, j int, val float64) {
	m.mat[i][j] = val
}

func (m *Matrix) addmat(matr [][]float64) {
	var i, j int
	for i = 0; i < m.rownum; i++ {
		for j = 0; j < m.colnum; j++ {
			m.mat[i][j] = m.mat[i][j] + matr[i][j]
		}
	}
}

func main() {
	// var v [][]float64
	v := [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	m := Matrix{2, 2, v}
	fmt.Println(m.mat)
	fmt.Println(m.rows())
	fmt.Println(m.cols())

	m.setval(0, 1, 5.0)
	fmt.Println(m.mat)

	m.addmat(v)
	fmt.Println(m.mat)

}
