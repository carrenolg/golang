package main

import (
	"fmt"
	"image/color"

	"carrenolg.io/books/gopl/ch6/geometry"
	"carrenolg.io/books/gopl/ch6/intlist"
	"carrenolg.io/books/gopl/ch6/urlvalues"
)

func main() {
	fmt.Println("6 .Methods")
	fmt.Println("\n\n6.1 Methods Declarations")
	// type Point
	p := geometry.Point{
		X: 1,
		Y: 2,
	}
	q := geometry.Point{
		X: 2,
		Y: 4,
	}
	fmt.Println(geometry.Distance(p, q)) // function call
	fmt.Println(p.Distance(q))           // method call

	// type Path
	perim := geometry.Path{
		geometry.Point{
			X: 1,
			Y: 1},
		geometry.Point{
			X: 5,
			Y: 1},
		geometry.Point{
			X: 5,
			Y: 4,
		},
		geometry.Point{
			X: 1,
			Y: 1,
		},
	}
	fmt.Println(geometry.Path.Distance(perim)) // standalone function
	fmt.Println(perim.Distance())              // method of geometry.Path

	fmt.Println("\n\n6.2 - Methods with a Pointer Receiver")
	// method calls
	r := &geometry.Point{
		X: 1,
		Y: 2,
	}
	r.ScaleBy(2)
	fmt.Println(*r)
	// method calls
	r2 := geometry.Point{
		X: 1,
		Y: 2,
	}
	pptr := &r2
	pptr.ScaleBy(2)
	fmt.Println(r2)
	// method calls
	r3 := geometry.Point{
		X: 1,
		Y: 2,
	}
	(&r3).ScaleBy(2)
	fmt.Println(r3)

	// Receiver types (variable and pointer)

	// x is a variable not a pointer
	x := geometry.Point{
		X: 5,
		Y: 5,
	}
	x.ScaleBy(2) // compiler adds (&x) implicitly
	fmt.Println(x)

	// y is a pointer not a variable
	y := &geometry.Point{
		X: 3,
		Y: 3,
	}
	y.ScaleBy(2)
	fmt.Println(*y)

	// compile error
	/*
		geometry.Point{
			X: 1,
			Y: 2,
		}.ScaleBy(2)
	*/

	// Receiver Arguments vs. Receiver Parameters
	a := geometry.Point{
		X: 15,
		Y: 25,
	}
	b := geometry.Point{
		X: 1,
		Y: 2,
	}

	// Parameter T, Argument *T
	aptr := &a // &a = return a memory address (pointer)
	//dis := aptr.Distance(b)   // implicit adds (*aptr) Dereference, load values
	dis := (*aptr).Distance(b) // Check
	fmt.Println(dis)

	// Parameter *T, Argument T
	bval := b          // variable
	bval.ScaleBy(2)    // implicit adds (&bval) reference, load memory address
	(&bval).ScaleBy(2) // check
	fmt.Println(bval)

	// Check types
	fmt.Printf("%T\n", aptr)
	fmt.Printf("%T\n", &bval)

	fmt.Println("\n\n6.2.1 - Nil is a Valid Receiver Value")
	// Wow
	l := intlist.IntList{
		Value: 5,
	}
	k := intlist.IntList{
		Value: 20,
		Tail:  &l,
	}
	j := intlist.IntList{
		Value: 75,
		Tail:  &k,
	}
	fmt.Println(j.Sum())

	// urlvalues
	urlvalues.Run()

	// Composing Types
	fmt.Println("\n\n6.3 - Composing Types by Struct Embedding")

	//
	var cp geometry.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Y = 2
	fmt.Println(cp.Y) // select field (Y) without mentioning Point*/

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	cp1 := geometry.ColoredPoint{
		Point: geometry.Point{
			X: 1,
			Y: 1,
		},
		Color: red,
	}
	cp2 := geometry.ColoredPoint{
		Point: geometry.Point{
			X: 5,
			Y: 4,
		},
		Color: blue,
	}

	// fmt.Println(cp1, cp2)
	// Check methods promoted
	fmt.Println("Distance:", cp1.Distance(cp2.Point))

	// share common structures and
	// vary relationships between objects dynamically
	cp2.Point = cp1.Point
	cp1.ScaleBy(2)
	fmt.Println(cp1.Point, cp2.Point)
}
