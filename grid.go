package main

import (
	"os"
	"fmt"
        "bitbucket.org/zombiezen/gopdf/pdf"
)

type grid struct {
	canvas *pdf.Canvas
}

func (g *grid) face(side, stretch pdf.Unit) {
	rows := pdf.Unit(6)
	for i := pdf.Unit(0); i < rows; i++ {
		x1 := i * side / rows;
		x2 := (i+1) * side / rows;
		y1 := side / 2 * (1 + stretch * i / rows)
		y2 := side / 2 * (1 + stretch * (i+1) / rows)
		path := pdf.Path{}
		path.Move(pdf.Point{x1, y1})
		path.Line(pdf.Point{x2, y2})
		path.Line(pdf.Point{x2, -y2})
		path.Line(pdf.Point{x1, -y1})
		path.Close()
		g.canvas.Stroke(&path)
	}
	path := pdf.Path{}
	for i := pdf.Unit(0); i < rows / 2; i++ {
		y1 := (i / rows) * side
		y2 := y1 * (1 + stretch)
		path.Move(pdf.Point{0, y1})
		path.Line(pdf.Point{side, y2})
		if i != 0 {
			path.Move(pdf.Point{0, -y1})
			path.Line(pdf.Point{side, -y2})
		}
	}
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
	//radius := square / 2
	canvas.Push()
	canvas.Translate(0, height / 2)
	g.face(square, 0)
	canvas.Pop()
	//canvas.Push()
	//canvas.Scale(float32((width - square) / (2 * square)), float32(height / square))
	//canvas.Translate(width / 2, height / 2)
	//g.face(radius, 0, height-square)
	//canvas.Pop()
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
