package main

import (
<<<<<<< HEAD
	"fmt"
=======
	"flag"
	"fmt"
	"image/color"
	"time"
>>>>>>> fb39ca26dd04ddf99154102bea155b65e21f5087

	"github.com/jung-kurt/gofpdf"
)

<<<<<<< HEAD
const (
	bannerHt = 94.0
	xIndent  = 40.0
	taxRate  = 0.09
)

type LineItem struct {
	UnitName       string
	PricePerUnit   int
	UnitsPurchased int
}

func main() {
	//////////////////////////VIDEO 4/////////////////////////////
	lineItems := []LineItem{
		{
			UnitName:       "2x6 Lumber - 8'",
			PricePerUnit:   375, // in cents
			UnitsPurchased: 220,
		}, {
			UnitName:       "Drywall Sheet",
			PricePerUnit:   822, // in cents
			UnitsPurchased: 50,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   1455, // in cents
			UnitsPurchased: 3,
		}, {
			UnitName:       "This is a line item with a very long description to test that our word wrapping is implemented and working as intended",
			PricePerUnit:   3211, // in cents
			UnitsPurchased: 3,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   5, // in cents
			UnitsPurchased: 3300,
		}, {
			UnitName:       "Paint",
			PricePerUnit:   332, // in cents
			UnitsPurchased: 44,
		},
	}
	subtotal := 0
	for _, li := range lineItems {
		subtotal += li.PricePerUnit * li.UnitsPurchased
	}
	tax := int(float64(subtotal) * taxRate)
	total := subtotal + tax
	totalUSD := toUSD(total)

	//////////////////////////VIDEO 1/////////////////////////////
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	/////////////////////////////VIDEO 1//////////////////////////////////////////////
	// // Basic Text Stuff
	// pdf.MoveTo(0, 0)
	// pdf.SetFont("arial", "B", 30)
	// _, lineHt := pdf.GetFontSize()
	// pdf.SetTextColor(255, 0, 0)
	// pdf.Text(0, lineHt, "Hello, world!")
	// pdf.MoveTo(0, lineHt*2.0)

	// pdf.SetFont("times", "", 18)
	// pdf.SetTextColor(100, 100, 100)
	// _, lineHt = pdf.GetFontSize()
	// pdf.MultiCell(0, lineHt, "Here is some text. If it is too long it will be word wrapped automatically. If there is a new line it will be\nwrapped as well (unlike other ways of writing text in gofpdf).", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// // Basic Shapes
	// pdf.SetFillColor(0, 255, 0)
	// pdf.SetDrawColor(0, 0, 255)
	// pdf.Rect(10, 100, 100, 100, "FD")
	// pdf.SetFillColor(100, 200, 200)
	// pdf.Polygon([]gofpdf.PointType{
	// 	{110, 250},
	// 	{160, 300},
	// 	{110, 350},
	// 	{60, 300},
	// 	{70, 230},
	// }, "F")

	// pdf.ImageOptions("images/gopher.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
	// 	ReadDpi: true,
	// }, 0, "")

	///////////////////////////VIDEO 2 ///////////////////////
	// top banner
	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * 0.8},
	}, "F")

	//banner logo
	pdf.ImageOptions("images/gopher.png", 248.0, 0+(bannerHt-(bannerHt/1.5))/2.0, 0, bannerHt/1.5, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	//bottom banner
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.25)},
		{w, h - (bannerHt * 0.1)},
		{w, h},
	}, "F")

	// Banner - INVOICE
	pdf.SetFont("arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/3.1, "INVOICE")

	// Banner - Phone, email, domain
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-2.0*124.0, (bannerHt-(lineHt*4.5))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "(123) 456-7890\njon@calhoun.io\nGophercises.com", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	// Banner - Address
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-124.0, (bannerHt-(lineHt*4.5))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "123 Fake St\nSome Town, PA\n12345", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	///////////////////////VIDEO 3	/////////////////////////

	// Summary - Billed To, Invoice #, Date of Issue (from function)
	_, sy := summaryBlock(pdf, xIndent, bannerHt+lineHt*2.0, "Billed To", "Client Name", "123 Client Address", "City, State, Country", "Postal Code")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*2.0, "Invoice Number", "0000000123")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*6.25, "Date of Issue", "05/29/2018")

	// Summary - Invoice Total
	x, y := w-xIndent-124.0, bannerHt+lineHt*2.25
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.CellFormat(124.0, lineHt, "Invoice Total", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x+2.0, y+lineHt*1.5
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 48)
	_, lineHt = pdf.GetFontSize()
	alpha := 58
	pdf.SetTextColor(72+alpha, 42+alpha, 55+alpha)
	totalUSD = "$1234.56"
	pdf.CellFormat(124.0, lineHt, totalUSD, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x-2.0, y+lineHt*1.25

	if sy > y {
		y = sy
	}

	x, y = xIndent-20.0, y+30.0
	pdf.Rect(x, y, w-(xIndent*2.0)+40.0, 3.0, "F")

	///////////////////////VIDEO 4	/////////////////////////

	// Line Items - headers
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	x, y = xIndent-2.0, y+lineHt
	pdf.MoveTo(x, y)
	pdf.CellFormat(w/2.65+1.5, lineHt, "Description", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignLeft, false, 0, "")
	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100.0, lineHt, "Price Per Unit", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, "Quantity", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, "Amount", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	// Line Items - real data using function and struct
	y = y + lineHt
	for _, li := range lineItems {
		x, y = lineItem(pdf, x, y, li)
	}

	// Subtotal etc
	x, y = w/1.75, y+lineHt*2.25
	x, y = trailerLine(pdf, x, y, "Subtotal", subtotal)
	x, y = trailerLine(pdf, x, y, "Tax", tax)
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(x+20.0, y, x+220.0, y)
	y = y + lineHt*0.5
	x, y = trailerLine(pdf, x, y, "Total", total)

	///////////////////////VIDEO 1	/////////////////////////

	// Grid
	//drawGrid(pdf)

	//catch error
	err := pdf.OutputFileAndClose("p4.pdf")
=======
type PDFOption func(*gofpdf.Fpdf)

func FillColor(c color.RGBA) PDFOption {
	return func(pdf *gofpdf.Fpdf) {
		r, g, b := rgb(c)
		pdf.SetFillColor(r, g, b)
	}
}

func rgb(c color.RGBA) (int, int, int) {
	alpha := float64(c.A) / 255.0
	alphaWhite := int(255 * (1.0 - alpha))
	r := int(float64(c.R)*alpha) + alphaWhite
	g := int(float64(c.G)*alpha) + alphaWhite
	b := int(float64(c.B)*alpha) + alphaWhite
	return r, g, b
}

type PDF struct {
	fpdf *gofpdf.Fpdf
	x, y float64
}

func (p *PDF) Move(xDelta, yDelta float64) {
	p.x, p.y = p.x+xDelta, p.y+yDelta
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) MoveAbs(x, y float64) {
	p.x, p.y = x, y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) Text(text string) {
	p.fpdf.Text(p.x, p.y, text)
}

func (p *PDF) Polygon(pts []gofpdf.PointType, opts ...PDFOption) {
	for _, opt := range opts {
		opt(p.fpdf)
	}
	p.fpdf.Polygon(pts, "F")
}

func main() {
	name := flag.String("name", "", "the name of the person who completed the course")
	flag.Parse()

	fpdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := fpdf.GetPageSize()
	fpdf.AddPage()
	pdf := PDF{
		fpdf: fpdf,
	}

	primary := color.RGBA{103, 60, 79, 255}
	secondary := color.RGBA{103, 60, 79, 220}

	// Top and bottom graphics
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{0, h / 9.0},
		{w - (w / 6.0), 0},
	}, FillColor(secondary))
	pdf.Polygon([]gofpdf.PointType{
		{w / 6.0, 0},
		{w, 0},
		{w, h / 9.0},
	}, FillColor(primary))
	pdf.Polygon([]gofpdf.PointType{
		{w, h},
		{w, h - h/8.0},
		{w / 6, h},
	}, FillColor(secondary))
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - h/8.0},
		{w - (w / 6), h},
	}, FillColor(primary))

	fpdf.SetFont("times", "B", 50)
	fpdf.SetTextColor(50, 50, 50)
	pdf.MoveAbs(0, 100)
	_, lineHt := fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "Certificate of Completion", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*2.0)

	fpdf.SetFont("arial", "", 28)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*2.0)

	fpdf.SetFont("times", "B", 42)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, *name, gofpdf.AlignCenter)
	pdf.Move(0, lineHt*1.75)

	fpdf.SetFont("arial", "", 22)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt*1.5, "For successfully completing all twenty programming exercises in the Gophercises programming course for budding Gophers (Go developers)", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*4.5)

	fpdf.ImageOptions("../images/gopher.png", w/2.0-50.0, pdf.y, 100.0, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	pdf.Move(0, 65.0)
	fpdf.SetFillColor(100, 100, 100)
	fpdf.Rect(60.0, pdf.y, 240.0, 1.0, "F")
	fpdf.Rect(490.0, pdf.y, 240.0, 1.0, "F")

	fpdf.SetFont("arial", "", 12)
	pdf.Move(0, lineHt/1.5)
	fpdf.SetTextColor(100, 100, 100)
	pdf.MoveAbs(60.0+105.0, pdf.y)
	pdf.Text("Date")
	pdf.MoveAbs(490.0+60.0, pdf.y)
	pdf.Text("Instructor - Jon Calhoun")
	pdf.MoveAbs(60.0, pdf.y-lineHt/1.5)
	fpdf.SetFont("times", "", 22)
	_, lineHt = fpdf.GetFontSize()
	pdf.Move(0, -lineHt)
	fpdf.SetTextColor(50, 50, 50)
	yr, mo, day := time.Now().Date()
	dateStr := fmt.Sprintf("%d/%d/%d", mo, day, yr)
	fpdf.CellFormat(240.0, lineHt, dateStr, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")
	pdf.MoveAbs(490.0, pdf.y)
	sig, err := gofpdf.SVGBasicFileParse("../images/sig.svg")
	if err != nil {
		panic(err)
	}
	pdf.Move(0, -(sig.Ht*.45 - lineHt))
	fpdf.SVGBasicWrite(&sig, 0.5)

	// fpdf.CellFormat(240.0, lineHt, "Jonathan Calhoun", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")

	// Grid
	// drawGrid(fpdf)
	err = fpdf.OutputFileAndClose("cert.pdf")
>>>>>>> fb39ca26dd04ddf99154102bea155b65e21f5087
	if err != nil {
		panic(err)
	}
}

<<<<<<< HEAD
///////////////////////VIDEO 4	/////////////////////////
func trailerLine(pdf *gofpdf.Fpdf, x, y float64, label string, amount int) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, label, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.SetTextColor(50, 50, 50)
	pdf.CellFormat(119.5, lineHt, toUSD(amount), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	y = y + lineHt*1.5
	return origX, y
}

func toUSD(cents int) string {
	centsStr := fmt.Sprintf("%d", cents%100)
	if len(centsStr) < 2 {
		centsStr = "0" + centsStr
	}
	return fmt.Sprintf("$%d.%s", cents/100, centsStr)
}

func lineItem(pdf *gofpdf.Fpdf, x, y float64, lineItem LineItem) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(50, 50, 50)
	pdf.MoveTo(x, y)
	x, y = xIndent-2.0, y+lineHt*.75
	pdf.MoveTo(x, y)
	pdf.MultiCell(w/2.65+1.5, lineHt, lineItem.UnitName, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	tmp := pdf.SplitLines([]byte(lineItem.UnitName), w/2.65+1.5)
	maxY := y + float64(len(tmp)-1)*lineHt
	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100.0, lineHt, toUSD(lineItem.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, fmt.Sprintf("%d", lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, toUSD(lineItem.PricePerUnit*lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	if maxY > y {
		y = maxY
	}
	y = y + lineHt*1.75
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(xIndent-10.0, y, w-xIndent+10.0, y)
	return origX, y
}

/////////////////VIDEO3//////////////

func summaryBlock(pdf *gofpdf.Fpdf, x, y float64, title string, data ...string) (float64, float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	y = y + lineHt
	pdf.Text(x, y, title)
	y = y + lineHt*.25
	pdf.SetTextColor(50, 50, 50)
	for _, str := range data {
		y = y + lineHt*1.25
		pdf.Text(x, y, str)
	}
	return x, y
}

///////////////VIDEO 1////////////////
=======
>>>>>>> fb39ca26dd04ddf99154102bea155b65e21f5087
func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
<<<<<<< HEAD
	for x := 0.0; x < w; x += (w / 20.0) {
=======
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
>>>>>>> fb39ca26dd04ddf99154102bea155b65e21f5087
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
<<<<<<< HEAD
	for y := 0.0; y < h; y += (w / 20.0) {
=======
	for y := 0.0; y < h; y = y + (w / 20.0) {
		pdf.SetTextColor(80, 80, 80)
>>>>>>> fb39ca26dd04ddf99154102bea155b65e21f5087
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
