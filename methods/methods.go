package methods

import (
	"math"
	"fmt"
)

type Face interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) EmptyInterfaceExample() {
	var emptyI interface{} = "string"
	value, ok := emptyI.(string)
	fmt.Println("value and ok: ", value, ok)
}

