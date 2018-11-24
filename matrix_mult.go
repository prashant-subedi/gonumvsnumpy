package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

var x = 500
var k = 500
var y = 500

func randFloats(n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = rand.Float64()
	}
	return res
}
func main() {
	u_raw := randFloats(x * k)
	v_raw := randFloats(k * y)

	u := mat.NewDense(x, k, u_raw)

	v := mat.NewDense(k, y, v_raw)

	// Start Timing
	a := time.Now()
	c := mat.NewDense(x, y, nil)
	c.Mul(u, v)
	d := time.Now()
	// End Timing

	duration := d.Sub(a)

	// fmt.Println("u = ", mat.Formatted(u, mat.Prefix(" "), mat.Squeeze()))
	// fmt.Println("v = ", mat.Formatted(v, mat.Prefix(" "), mat.Squeeze()))
	// fmt.Println("result=", mat.Formatted(c, mat.Prefix(" "), mat.Squeeze()))

	fmt.Println("Duration is ", duration.Seconds(), "s")

}
