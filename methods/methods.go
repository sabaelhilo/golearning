package methods

import (
	"math"
	"fmt"
	"strconv"
)

type IPAddr [4]byte

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

func (ip IPAddr) String() string {
	var s string

	for i, v := range ip {
		vs := strconv.Itoa(int(v))
		if i == 0 {
			s = vs
		} else {
			s = s + "." + vs
		}
	}
	return s
}


