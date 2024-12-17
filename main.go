package main

import (
	"fmt"
	"strconv"
	s "strings"
)

type point struct {
	x, y int
}

func main() {
	a, b := 1, 2

	// raw string
	str := `A "raw" 
        string`

	fmt.Println("golang test", a, b, str)

	// array
	nums := [...]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[3:])

	var towDimention [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			towDimention[i][j] = i + j
		}
	}
	fmt.Println("2D Array:", towDimention)

	// pointer
	ptr := getpointer()
	*ptr = 2
	fmt.Println("outside addr a", ptr)
	fmt.Println("outside addr ptr", &ptr)

	dynamicptr := new(int)
	fmt.Println("dynamic alloc ptr", dynamicptr)

	interger := 90
	sstr := string(interger)
	fmt.Println("interger", interger, "string", sstr)

	str = "90"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("conv failed")
	} else {
		fmt.Println(num)
	}

	fmt.Println(s.Contains(string("abc"), "bc"))

	pt := point{1, 2}
	fmt.Printf("%v\n", pt)
	fmt.Printf("%+v\n", pt)
	fmt.Printf("%#v\n", pt)
	fmt.Printf("%T\n", pt)

	xx, yy := autoreturn()
	fmt.Println("auto return", xx, yy)

	nnums := []int{1, 2, 3, 4}
	fmt.Println("sum", sums(nnums...))
	fmt.Println("sum", sums(1, 2, 3, 4))

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map", m)

	addfunc := func(a, b int) int {
		return a + b
	}
	fmt.Println("1 + 2 =", addfunc(1, 2))

	inner, val := outer()
	fmt.Println(val)
	fmt.Println(inner())
	
}

// => 200，这里涉及到golang中闭包和内存逃逸的概念，inner()实际上执行了两次，outer()中一次，fmt又一次，
//但为什么是200呢，编译器不能确定outer_var在后续会不会使用，
//所以outer_var不会随着outer()结束而释放它的栈（Stack）空间，
//而会‘逃逸到’堆（Heap）上，那么第二次的inner()中outer_var就会是101。
func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99
		return outer_var
	}
	inner()

	go func(from string){
		for i := 0; i < 3; i++{
			fmt.Println(from, ":", i)
		}
	}("hello")

	go func(from string){
		fmt.Println(from)
	}("going")


	return inner, outer_var
}

func autoreturn() (x int, y int) {
	return
}

func sums(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func getpointer() (_ *int) {
	a := 1
	fmt.Println("local a addr", &a)
	return &a
}
