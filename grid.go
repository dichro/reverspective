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
	lt := bounds.Min
	rb := bounds.Max
	rt := pdf.Point{rb.X, lt.Y}
	lb := pdf.Point{lt.X, rb.Y}
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
	}, 0, 0)
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
