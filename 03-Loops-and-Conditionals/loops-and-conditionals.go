package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v is out-of-scope here
	return lim
}

func sqrt2(x float64) float64 {
	z := 1.0;
	for i := 0; i < 10; i++ {
		// Similar to Newton's method
		z -= (z * z - x) / (2 * z);
		fmt.Printf("Iteration %v: z = %v\n", i, z);
	}
	return z;
}

func main() {
	sum := 0;
	for i := 0; i < 10; i++ {
		sum += i;
	}
	fmt.Println(sum);
	sum = 1;
	// Identical to "while" loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// Infinite loop
	// for {
	// }
	
	fmt.Println(sqrt(2), sqrt(-4));
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	);
	fmt.Println(sqrt2(float64(425)));

	fmt.Print("Go runs on ... ");
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.");
	case "linux":
		fmt.Println("Linux.");
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os);
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world");

	fmt.Println("hello");

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
