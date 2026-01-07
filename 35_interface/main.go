package main

import (
	"fmt"
	"math"
)

// interface di golang itu semacam kontrak function
// dia memuat kumpulan nama function  untuk digunakan oleh struck
// contoh :

type geometry interface{
	area() float64
	perimeter() float64
}

type Rect struct{
	width, height float64
}

type Circle struct {
	radius float64
}

func (rectangle Rect) area() float64{
	return rectangle.width * rectangle.height
}
func (rectangle Rect) perimeter() float64{
	return 2*(rectangle.width + rectangle.height)
}

func (cirlce Circle) area() float64{
	return math.Pi * cirlce.radius * cirlce.radius
}

func (Circle Circle) perimeter() float64{
	return 2 * math.Pi * Circle.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perimeter())
}




func detectCircle(g geometry) {
    if c, ok := g.(Circle); ok {
        fmt.Println("circle with radius", c.radius)
    }else{
		fmt.Println("not circle")
	}
}

func main (){
	r := Rect{width: 3, height:4}
	c := Circle{radius: 5}

//	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

}