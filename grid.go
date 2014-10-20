package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"math"
	"os"
)

func grid(canvas *pdf.Canvas, side, stretch pdf.Unit) {
	rows := pdf.Unit(6)
	delta := side * (stretch - 1) / 2
	for i := pdf.Unit(0); i < rows; i++ {
		x1 := i * side / rows
		x2 := (i + 1) * side / rows
		y1 := side/2 + delta*i/rows
		y2 := side/2 + delta*(i+1)/rows
		path := pdf.Path{}
		path.Move(pdf.Point{x1, y1})
		path.Line(pdf.Point{x2, y2})
		path.Line(pdf.Point{x2, -y2})
		path.Line(pdf.Point{x1, -y1})
		path.Close()
		canvas.Stroke(&path)
	}
	path := pdf.Path{}
	for i := pdf.Unit(0); i < rows/2; i++ {
		y1 := (i / rows) * side
		y2 := y1 * stretch
		path.Move(pdf.Point{0, y1})
		path.Line(pdf.Point{side, y2})
		if i != 0 {
			path.Move(pdf.Point{0, -y1})
			path.Line(pdf.Point{side, -y2})
		}
	}
	canvas.Stroke(&path)
}

func main() {
	width := pdf.USLetterWidth
	height := pdf.USLetterHeight
	square := 5 * pdf.Inch
	doc := pdf.New()
	canvas := doc.NewPage(width, height)
	canvas.Push()
	canvas.Translate((width-square)/2, height/2)
	grid(canvas, square, 1)
	canvas.Pop()
	canvas.Push()
	canvas.Translate(width/2, (height+square)/2)
	canvas.Rotate(math.Pi / 2)
	canvas.Scale(float32((height-square)/2/square), 1)
	grid(canvas, square, width/square)
	canvas.Pop()
	canvas.Push()
	canvas.Translate(width/2, (height-square)/2)
	canvas.Rotate(-math.Pi / 2)
	canvas.Scale(float32((height-square)/2/square), 1)
	grid(canvas, square, width/square)
	canvas.Pop()
	canvas.Push()
	canvas.Translate((width-square)/2, height/2)
	canvas.Rotate(math.Pi)
	canvas.Scale(float32((width-square)/2/square), 1)
	grid(canvas, square, height/square)
	canvas.Pop()
	canvas.Push()
	canvas.Translate((width+square)/2, height/2)
	canvas.Scale(float32((width-square)/2/square), 1)
	grid(canvas, square, height/square)
	canvas.Pop()
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
