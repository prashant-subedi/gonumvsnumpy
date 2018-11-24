package main

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main() {
	u := mat.NewDense(3, 3, []float64{1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	v := mat.NewDense(3, 3, []float64{1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	c := mat.NewDense(3, 3, nil)
	a := time.Now()
	c.Mul(u, v)
	d := time.Now()
	duration := d.Sub(a)
	fmt.Print("Duration is ", duration)
	fmt.Println(mat.Formatted(c, mat.Prefix(" "), mat.Squeeze()))
}
