package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows int
	columns int
	values [][]int
}

// get the number of rows
func (matrix Matrix) getRows() int {
	return matrix.rows
}

// get the number of columns
func (matrix Matrix) getColumns() int {
	return matrix.columns
}

// initialise matrix
func (matrix *Matrix) initialiseMatrix() {
	r, c, v := matrix.rows, matrix.columns, matrix.values
	for i:= 0; i<r; i++ {
		v[i] = make([]int, c)
	}
}

// set the elements of the matrix at a given position (i,j)
func (matrix *Matrix) setValues(i, j, val int) {
	matrix.values[i][j] = val
}


// adding two matrices
func (matrix *Matrix) addMatrix(matrix2 Matrix) Matrix {
	r, c := matrix.rows, matrix.columns
	result := Matrix{r,c,make([][]int, r)}

	for i:=0; i<r; i++ {
		result.values[i] = make([]int, c)
		for j:=0; j<c; j++ {
			result.values[i][j] = matrix2.values[i][j] + matrix.values[i][j]
		}
	}
	return result
}

// print matrix structure as json using inbuilt function
func (matrix Matrix) printJSON() {
	j, err := json.Marshal(matrix.values)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(j))
	}
}

// print matrix structure as json using custom function
func (matrix Matrix) printJSONCustom() {
	r, c, v := matrix.rows, matrix.columns, matrix.values
	fmt.Println("[")
	for i:=0; i<r; i++ {
		fmt.Println(" [")
		for j:=0; j<c; j++ {
			fmt.Printf("  %v",v[i][j])
			if j != c-1 {
				fmt.Printf(",")
			}
			fmt.Println()
		}
		fmt.Printf(" ]")
		if i != r-1 {
			fmt.Printf(",")
		}
		fmt.Println()
	}
	fmt.Println("]")
}

func main() {
	m := Matrix{2,2, make([][]int, 2)}
	m.initialiseMatrix()

	fmt.Println(m.getRows())

	fmt.Println(m.getColumns())

	m.setValues(0, 0, 4)
	m.setValues(0, 1, 2)
	m.setValues(1, 0, 3)
	m.setValues(1, 1, 9)

	m2 := Matrix{2, 2, make([][]int, 2)}
	m2.initialiseMatrix()
	m2.setValues(0, 0, 2)
	m2.setValues(0, 1, 3)
	m2.setValues(1, 0, 7)
	m2.setValues(1, 1, 5)

	fmt.Println(m.addMatrix(m2))

	m.printJSON()
	m2.printJSONCustom()

}