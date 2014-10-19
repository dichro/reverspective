package main

import (
	"os"
	"fmt"
        "bitbucket.org/zombiezen/gopdf/pdf"
)

func main() {
	width := pdf.USLetterWidth
	height := pdf.USLetterHeight
	square := 5 * pdf.Inch
	lt := pdf.Point{X: (width - square) / 2, Y: (height - square) / 2}
	rt := lt
	rt.X += square
	rb := rt
	rb.Y += square
	lb := lt
	lb.Y += square
	path := pdf.Path{}
	path.Move(lt)
	path.Line(rt)
	path.Line(rb)
	path.Line(lb)
	path.Close()

	doc := pdf.New()
	canvas := doc.NewPage(width, height)
	canvas.Stroke(&path)
	canvas.Close()
	err := doc.Encode(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
