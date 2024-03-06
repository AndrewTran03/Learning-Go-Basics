package main

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { 
	return x*10 + 1 
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// var c, python, java bool;
var i, j int = 1, 2;
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
);
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
);



func main() {
	fmt.Println("My favorite number is", rand.Intn(10));
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7));
	fmt.Println(math.Pi);
	fmt.Println(add(42, 13));
	a, b := swap("hello", "world");
	fmt.Println(a, b);
	fmt.Println(split(17));
	// var i int;
	// fmt.Println(i, c, python, java);
	var c, python, java = true, false, "no!";
	fmt.Println(i, j, c, python, java);
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe);
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt);
	fmt.Printf("Type: %T Value: %v\n", z, z);
	var i int;
	var f float64;
	var bl bool;
	var s string;
	fmt.Printf("%v %v %v %q\n", i, f, bl, s);
	var x, y int = 3, 4;
	var fl float64 = math.Sqrt(float64(x*x + y*y));
	var z uint = uint(fl);
	fmt.Println(x, y, z);
	v := '4'; // change me!
	fmt.Printf("v is of type %T\n", v);
	const World = "世界";
	fmt.Println("Hello", World);
	const Pi = math.Pi;
	fmt.Println("Happy", Pi, "Day");
	const Truth = true
	fmt.Println("Go rules?", Truth);
	fmt.Println(needInt(Small));
	fmt.Println(needFloat(Small));
	fmt.Println(needFloat(Big));
}

