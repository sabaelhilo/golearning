package main

import (
	"fmt"
	"math/rand"

	"math"

	"time"

	"github.com/user/stringutil"

	"strings"
)

var (
	variable, python bool = false, false
	java             bool = true
)

var i, j int = 1, 2

const MyAge = 29

const (
	//shift a 1 bit left 100 places
	Big = 1 << 100
	//shift it right 99 places
	Small = Big >> 99
)

type Vertex struct {
	X int
	Y int
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func newtonSqrt(x float64) float64 {

	var start float64 = 1
	var next float64 = 2
	for math.Abs(next-start) > needFloat(Small) {
		start = next
		next = start - ((math.Pow(start, 2) - x) / (2 * start))
	}

	return next
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for index := range picture {
		row := make([]uint8, dx)
		for rowX := range row {
			row[rowX] = uint8((rowX ^ index))
		}
		picture[index] = row
	}
	return picture
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X) * float64(v.X) + float64(v.Y) * float64(v.Y))
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	position := 0
	valueArray := make([]int, 2)
	return func() int {
		value := 0
		if position == 0 || position == 1 {
			value = position
			valueArray[position] = value
		} else {
			for _, v := range valueArray {
				value = value + v
			}
			valueArray[0] = valueArray[1]
			valueArray[1] = value
		}
		position = position + 1
		return value
	}
}

func WordCount(s string) map[string]int {
	//return a map of the counts of each word in s - s is a sentence
	fields := strings.Fields(s)
	m := make(map[string]int)
	for _, value := range fields {
		if _, ok := m[value]; ok {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(stringutil.Reverse("hello, world\n"))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Println()
	fmt.Println(add(10, 12))

	// implicit declaration
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	c, d := split(17)
	fmt.Println(c, d)

	fmt.Println(i, j, variable, python, java)

	var v1, v2 int = 3, 4
	var f float64 = math.Sqrt(float64(v1*v1 + v2*v2))
	var v3 uint = uint(f)
	fmt.Println(v1, v2, v3)

	v := 42.2
	fmt.Println("v is of type %T\n", v)
	fmt.Println("My age is %d", MyAge)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}

	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	defer fmt.Println("Deffering this")

	fmt.Printf("newton sqrt implementation is %g\n", newtonSqrt(64))

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")

	}

	//pointers - pointer holds memory address of a variable
	k, m := 42, 2701

	p := &k         // p now points to k
	fmt.Println(*p) //42
	*p = 21
	fmt.Println(k) // 21

	p = &m
	*p = *p / 37
	fmt.Println(m)  // 73
	fmt.Println(*p) // 73
	fmt.Println(k)  // stays 21

	vert := Vertex{1, 2}
	fmt.Println(vert.X)
	fmt.Println(vert.Y)

	var a1 [2]string
	a1[0] = "hello"
	a1[1] = "world"

	primes := [6]int{2, 3, 5, 7, 11, 13}

	for index, value := range primes {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	fmt.Println(primes)

	//slices dont store any data - just describes a section of an underlying array
	//changing elements in slice change the array
	var s []int = primes[1:4]
	fmt.Println(s)

	//dynamically sized arrays - creates array w/ 5 zeros (whatever the zero type is)
	a3 := make([]int, 5)
	printSlice("a", a3)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"

	m1 := make(map[string]Vertex)
	m1["bell labs"] = Vertex{10, 12}

	fmt.Println(m1["Bell Labs"])

	hypot := func(x, y float64) float64 {
		return math.Sqrt(y*x + y*y)
	}

	hypot(2, 2)

}
