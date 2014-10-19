package main

import (
	"os"
	"fmt"
        "bitbucket.org/zombiezen/gopdf/pdf"
)

type grid struct {
	canvas *pdf.Canvas
}

func (g *grid) face(bounds pdf.Rectangle, xShrink, yShrink pdf.Unit) {
	lb := bounds.Min
	rt := bounds.Max
	rb := pdf.Point{rt.X, lb.Y}
	lt := pdf.Point{lb.X, rt.Y}
	delta := yShrink / 2
	if yShrink > 0 {
		rt.Y -= delta
		rb.Y += delta
	} else {
		lt.Y += delta
		lb.Y -= delta
	}
	delta = xShrink / 2
	if xShrink < 0 {
		lb.X -= delta
		rb.X += delta
	} else {
		lt.X += delta
		rt.X -= delta
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
	xStart, yStart := (width - square) / 2, (height - square) / 2
	doc := pdf.New()
	canvas := doc.NewPage(width, height)
	g := &grid{canvas: canvas}
	middle := pdf.Rectangle{
		Min: pdf.Point{xStart, yStart},
		Max: pdf.Point{xStart+square, yStart+square},
	}
	g.face(middle, 0, 0)
	g.face(pdf.Rectangle{
		Min: pdf.Point{0, 0},
		Max: pdf.Point{xStart, height},
	}, 0, height-square)
	g.face(pdf.Rectangle{
		Min: pdf.Point{0, middle.Max.Y},
		Max: pdf.Point{width, height},
	}, square - width, 0)
	g.face(pdf.Rectangle{
		Min: pdf.Point{middle.Max.X, 0},
		Max: pdf.Point{width, height},
	}, 0, square-height)
	g.face(pdf.Rectangle{
		Min: pdf.Point{0, 0},
		Max: pdf.Point{width, yStart},
	}, width-square, 0)
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
