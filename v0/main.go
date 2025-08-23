package main

import "fmt"

func main() {

	// defer statements execute just before func exit when all statements done LIFO
	defer fmt.Println("main 1")

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

	var mp2 map[string][]string
	fmt.Println(mp2)

	mp2 = map[string][]string{
		"k1": {"v1"},
	}
	mp2["k2"] = []string{"v2"}
	fmt.Println(mp2["k2"])

	delete(mp2, "k2")
     
	// ok can be any name return boolean
	v3, ok := mp2["k1"]
	fmt.Println(v3, ok)

	var s1 struct{
		nm string
		id int
	}
	fmt.Println(s1)
	s1.nm = "Z"
	fmt.Println(s1.nm)

	type s4 struct{
		nm string
		id int
	}
	var v4 s4
	fmt.Println(v4)

	v4 = s4{
		nm: "Z2",
		id: 2,
	}

	// structs creates new copy on assignment (copy by value)
	v5 := v4
	v4.nm = "Z3"
	fmt.Println(v4)
	fmt.Println(v5)

	// all loops in go are for with different forms

	i := 1
	for {
		i += 2
		fmt.Println(i)
		break
	}

	for i<10 {
		fmt.Println(i)
		i++
	}

	for i2 :=1; i2<10; i2++ {
		fmt.Println(i2)
	}

	// arrays, slices, maps
	// can't loop struct if needed use reflection
	// _, v _ is blanket identifier
	a2 := []int{1,2,3}
	for a,b := range a2{
		fmt.Println(a,b)
	}

	// i3:=10 can even be declared in separate line
	if i3:=10;i3<15 {
		fmt.Println("less")
	} else if i3>15 {
		fmt.Println("more")
	} else {
		fmt.Println("na")
	}

	// true can be skipped by default it is true
	// switch i:=8;true
	switch i4:=8; {
	case i4 < 10:
		fmt.Println("less")
	default:
		fmt.Println("na")
	}

	defer fmt.Println("main 2")

	// panic() built in func in go
	// defer recover is use case to prevent prog from crashing

	func1()

}

func func1() {
    
	fmt.Println("func1")

	i := 10
	if i < 20 {
		goto zlabel
	}
	zlabel:
	fmt.Println("label in func1")

	// register recover before panic
	// anonymous func
	// if needed function named as func2 define it and then call
	defer func() {
		fmt.Println(recover())
	}()

	// here anything can be passed
	panic("panic mode")
	fmt.Println("func1 2")
}