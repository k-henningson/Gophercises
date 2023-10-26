package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)

	pdf.AddPage()

	// Text layout and style
	pdf.MoveTo(0, 0)
	pdf.SetFont("arial", "B", 30)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(40, lineHt, "Hello, world")
	pdf.MoveTo(0, lineHt*2.0)

	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*1.5, "This is some text. If it is too long it will be word wrapped. If there is a new line it will\n be wrapped as well.", gofpdf.BorderNone, gofpdf.AlignRight, false)

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}
