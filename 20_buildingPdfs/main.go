package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// TOP MAROON HEADER
	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * .9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * .2)},
		{w, h - (bannerHt * .1)},
		{w, h},
	}, "F")

	// HEADER INVOICE
	pdf.SetFont("Arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/3, "INVOICE")

	//  HEADER PHONE EMAIL & DOMAIN
	pdf.SetFont("Arial", "", 12)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-2.0*124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "(123) 456-7890\nemail@gmail.com\nGophercises.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

	//  HEADER ADDRESS
	pdf.SetFont("Arial", "", 12)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "123 Brown Street\nVancouver, BC\nCanada", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// // Text layout and style
	// pdf.MoveTo(0, 0)
	// pdf.SetFont("arial", "B", 30)
	// _, lineHt := pdf.GetFontSize()
	// pdf.SetTextColor(255, 0, 0)
	// pdf.Text(40, lineHt, "Hello, world")
	// pdf.MoveTo(0, lineHt*2.0)

	// pdf.SetFont("times", "", 18)
	// pdf.SetTextColor(100, 100, 100)
	// _, lineHt = pdf.GetFontSize()
	// pdf.MultiCell(0, lineHt*1.5, "This is some text. If it is too long it will be word wrapped. If there is a new line it will\n be wrapped as well.", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// // Shapes
	// pdf.SetFillColor(0, 255, 0)
	// pdf.SetDrawColor(0, 0, 255)
	// pdf.Rect(10, 100, 100, 100, "FD")
	// pdf.SetFillColor(100, 200, 200)
	// pdf.Polygon([]gofpdf.PointType{
	// 	{110, 250},
	// 	{160, 300},
	// 	{110, 350},
	// 	{60, 300},
	// }, "F")

	// pdf.ImageOptions("images/jump.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
	// 	ReadDpi: true,
	// }, 0, "")

	// Grid
	// drawGrid(pdf)

	err := pdf.OutputFileAndClose("p2.pdf")
	if err != nil {
		panic(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (h / 20.0) {
		if y < bannerHt*.9 {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
