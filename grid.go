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

type face struct {
	tx, ty  pdf.Unit
	rot     float32
	sx, sy  pdf.Unit
	stretch pdf.Unit
}

func faces(width, height, square pdf.Unit) []face {
	return []face{
		{
			(width - square) / 2, height / 2,
			0,
			1, 1,
			1,
		},
		{
			width / 2, (height + square) / 2,
			math.Pi / 2,
			(height - square) / 2 / square, 1,
			width / square,
		},
		{
			width / 2, (height - square) / 2,
			-math.Pi / 2,
			(height - square) / 2 / square, 1,
			width / square,
		},
		{
			(width - square) / 2, height / 2,
			math.Pi,
			(width - square) / 2 / square, 1,
			height / square,
		},
		{
			(width + square) / 2, height / 2,
			0,
			(width - square) / 2 / square, 1,
			height / square,
		},
	}
}

func project(width, height, square, stretch, squash pdf.Unit) []face {
	return []face{
		{
			(width - square) / 2, height / 2,
			0,
			1, 1,
			1,
		},
		{
			width / 2, (height + square) / 2,
			math.Pi / 2,
			squash, 1,
			stretch,
		},
		{
			width / 2, (height - square) / 2,
			-math.Pi / 2,
			squash, 1,
			stretch,
		},
		{
			(width - square) / 2, height / 2,
			math.Pi,
			squash, 1,
			stretch,
		},
		{
			(width + square) / 2, height / 2,
			0,
			squash, 1,
			stretch,
		},
	}
}

func main() {
	width := pdf.USLetterWidth
	height := pdf.USLetterHeight
	square := 4 * pdf.Inch
	doc := pdf.New()
	canvas := doc.NewPage(width, height)
	for _, face := range project(width, height, square, 1.6, 0.5) {
		canvas.Push()
		canvas.Translate(face.tx, face.ty)
		canvas.Rotate(face.rot)
		canvas.Scale(float32(face.sx), float32(face.sy))
		grid(canvas, square, face.stretch)
		canvas.Pop()
	}
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
