package main

import (
	"bitbucket.org/zombiezen/gopdf/pdf"
	"fmt"
	"math"
	"os"
)

func (f *face) grid(canvas *pdf.Canvas, side pdf.Unit) {
	canvas.Push()
	canvas.Translate(f.tx, f.ty)
	canvas.Rotate(f.rot)
	canvas.Scale(float32(f.sx), 1)
	rows := pdf.Unit(6)
	delta := side * (f.stretch - 1) / 2
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
		canvas.SetColor(f.grayFill, f.grayFill, f.grayFill)
		canvas.SetStrokeColor(f.grayStroke, f.grayStroke, f.grayStroke)
		canvas.FillStroke(&path)
	}
	path := pdf.Path{}
	for i := pdf.Unit(0); i < rows/2; i++ {
		y1 := (i / rows) * side
		y2 := y1 * f.stretch
		path.Move(pdf.Point{0, y1})
		path.Line(pdf.Point{side, y2})
		if i != 0 {
			path.Move(pdf.Point{0, -y1})
			path.Line(pdf.Point{side, -y2})
		}
	}
	canvas.Stroke(&path)
	canvas.Pop()
}

type face struct {
	tx, ty  pdf.Unit
	rot     float32
	sx      pdf.Unit
	stretch pdf.Unit
	grayFill, grayStroke float32
}

func faces(width, height, square pdf.Unit, colour float32) []face {
	return []face{
		{
			(width - square) / 2, height / 2,
			0,
			1,
			1,
			colour, 0,
		},
		{
			width / 2, (height + square) / 2,
			math.Pi / 2,
			(height - square) / 2 / square,
			width / square,
			colour, 0,
		},
		{
			width / 2, (height - square) / 2,
			-math.Pi / 2,
			(height - square) / 2 / square,
			width / square,
			colour, 0,
		},
		{
			(width - square) / 2, height / 2,
			math.Pi,
			(width - square) / 2 / square,
			height / square,
			colour, 0,
		},
		{
			(width + square) / 2, height / 2,
			0,
			(width - square) / 2 / square,
			height / square,
			colour, 0,
		},
	}
}

func project(width, height, square, stretch pdf.Unit) []face {
	narrowest := width
	if height < width {
		narrowest = height
	}
	squash := (narrowest - square) / 2 / square
	return []face{
		{
			(width - square) / 2, height / 2,
			0,
			1,
			1,
			1, 0.9,
		},
		{
			width / 2, (height + square) / 2,
			math.Pi / 2,
			squash,
			stretch,
			1, 0.9,
		},
		{
			width / 2, (height - square) / 2,
			-math.Pi / 2,
			squash,
			stretch,
			0.5, 0,
		},
		{
			(width - square) / 2, height / 2,
			math.Pi,
			squash,
			stretch,
			0.8, 0,
		},
		{
			(width + square) / 2, height / 2,
			0,
			squash,
			stretch,
			0.8, 0,
		},
	}
}

func main() {
	width := pdf.USLetterWidth
	height := pdf.USLetterHeight
	square := 4 * pdf.Inch
	doc := pdf.New()
	canvas := doc.NewPage(width, height)
	for _, face := range project(width, height, square, 1.6) {
		face.grid(canvas, square)
	}
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
