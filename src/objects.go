package main

import "fmt"

type ifigure interface {
	draw()
}

// rectangle
type rectangle struct {
	width, height float64
}

func (r rectangle) draw() {
	fmt.Printf("Printing rectangle width(%f) height(%f)\n", r.width, r.height)
}

func (r *rectangle) Width() float64 {
	return r.width
}

// circle
type circle struct {
	x, y, radii float64
}

func (c circle) draw() {
	fmt.Printf("Printing circle x(%f) y(%f) radii(%f)\n", c.x, c.y, c.radii)
}

func (c circle) Perimeter() float64 {
	return 2 * 3.141516 * c.radii
}

func MainObjects() {

	fmt.Println("==============================================================")
	fmt.Println("Objects")
	fmt.Println("==============================================================")

	var inst1 ifigure = rectangle{width: 3, height: 4}
	inst1.draw()

	var figures []ifigure

	figures = append(figures, rectangle{width: 31, height: 64})
	figures = append(figures, circle{x: 67, y: 87, radii: 25})

	for _, i := range figures {
		i.draw()
		var c, ok = i.(circle)
		if ok {
			fmt.Printf("found a circle with perimeter (%f)\n", c.Perimeter())
		}
	}
}
