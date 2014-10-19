package main

import (
	"os"
	"fmt"
        "bitbucket.org/zombiezen/gopdf/pdf"
)

type grid struct {
	canvas *pdf.Canvas
}

func (g *grid) face(radius pdf.Unit, xShrink, yShrink pdf.Unit) {
	lb := pdf.Point{-radius, -radius}
	rt := pdf.Point{radius, radius}
	rb := pdf.Point{rt.X, lb.Y}
	lt := pdf.Point{lb.X, rt.Y}
	delta := yShrink / 2
	if yShrink >= 0 {
		rt.Y *= 1 - delta
		rb.Y *= 1 - delta
	} else {
		lt.Y *= 1 + delta
		lb.Y *= 1 + delta
	}
	delta = xShrink / 2
	if xShrink <= 0 {
		lb.X *= 1 + delta
		rb.X *= 1 + delta
	} else {
		lt.X *= 1 - delta
		rt.X *= 1 - delta
	}
	path := pdf.Path{}
	path.Move(lt)
	path.Line(rt)
	path.Line(rb)
	path.Line(lb)
	path.Close()
	g.canvas.Stroke(&path)
}

func main() {
	width := pdf.USLetterWidth
	height := pdf.USLetterHeight
	square := 5 * pdf.Inch
	//xStart, yStart := (width - square) / 2, (height - square) / 2
	doc := pdf.New()
	canvas := doc.NewPage(width, height)
	g := &grid{canvas: canvas}
	radius := square / 2
	canvas.Push()
	canvas.Translate(width / 2, height / 2)
	g.face(radius, 0, 0)
	canvas.Pop()
	canvas.Push()
	//canvas.Scale(float32((width - square) / (2 * square)), float32(height / square))
	canvas.Translate(width / 2, height / 2)
	g.face(radius, 0, height-square)
	canvas.Pop()
//	g.face(pdf.Rectangle{
//		Min: pdf.Point{0, middle.Max.Y},
//		Max: pdf.Point{width, height},
//	}, square - width, 0)
//	g.face(pdf.Rectangle{
//		Min: pdf.Point{middle.Max.X, 0},
//		Max: pdf.Point{width, height},
//	}, 0, square-height)
//	g.face(pdf.Rectangle{
//		Min: pdf.Point{0, 0},
//		Max: pdf.Point{width, yStart},
//	}, width-square, 0)
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
