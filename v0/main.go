package main

import "fmt"

func main() {
	//println("go v0")
	fmt.Println("go v0")

    // vars in go strongly typed
    // declare
    var v string

    // declare and init
    var v1 string = "z"
	fmt.Println(v,v1)

	//initialize with inferred type
	var v2 = "z2"
	fmt.Println(v2)

	// short declare syntax
	z3 := "z3"
	fmt.Println(z3)

	// go doesn't support implicit conversions
	var f int = 1
	var f2 float32
	// error
	//f2 = f
	f2 = float32(f) 
	fmt.Println(f2)

	// multiple var declaration in 1 line
	a,b := 1,2 
	c := a + b
	d := a == b
	fmt.Println(a,b,c,d)
 
	// const can go unused
	const x = 10

	var y int = x
	fmt.Println(y)

	// iota starts at 0 within each const block increment by 1 with each constant
	const (
		m = 2 / 2
		n = iota
		l
		p = 5
		q
	)
	fmt.Println(m,n,l,p,q)

	s := "S"

	// creates a copy
	s2 := s

	s = "A"

	// pointer
	s3 := &s

	fmt.Println(s,s2,*s3,s3)

	var arr [3]int
	fmt.Println(arr)

	arr = [3]int{1,2}
	fmt.Println(arr)

	// slice doesn't need size unlike array
	var sl []int
	fmt.Println(sl)
	sl = append(sl,1)
	fmt.Println(sl)

	// array - fixed size, slice/map - dynamic
	var mp map[string]int
	fmt.Println(mp)
	mp = map[string]int{"a":1}
	fmt.Println(mp["a"])
}