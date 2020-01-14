package main

import (
	"Lv/internal/bootstrap"
	"bytes"
	"fmt"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"strings"
	"time"
)

func main(){
	err :=ExampleNewPDFGenerator()
	if err != nil {
		fmt.Println(err.Error())
	} else{
		fmt.Println("success")
	}
}

func ExampleNewPDFGenerator()(err error) {

	// Create new PDF generator
	var pdfg *wkhtmltopdf.PDFGenerator
	pdfg, err = wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		bootstrap.WriteLog(
			"ERROR",
			"CreateAuditRecordListPDFWithHTML: 建立PDF產生器失敗, "+err.Error(),
		)
		return
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	/**
	 * 產生HTML
	 */
	var tpl *template.Template
	tpl, err = template.New("security.html").ParseFiles(
		"app/security.html",
	)
	if err != nil {
		bootstrap.WriteLog(
			"ERROR",
			"CreateAuditRecordListPDFWithHTML: 建立新模板失敗, "+err.Error(),
		)
		return
	}

	renderData := []string{}
	writer := new(bytes.Buffer)
	err = tpl.Execute(writer, renderData)
	if err != nil {
		bootstrap.WriteLog(
			"ERROR",
			"CreateAuditRecordListPDFWithHTML: 產生模板內容失敗, "+err.Error(),
		)
		return
	}

	// Create from local file
	page := wkhtmltopdf.NewPageReader(writer)

	page.FooterRight.Set("[page]")
	page.MinimumFontSize.Set(12)
	page.FooterFontSize.Set(12)
	page.Zoom.Set(0.95)
	pdfg.AddPage(page)


	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	version := 1
	pdfName := fmt.Sprintf(
		"審查紀錄表_%s_V%d.pdf",
		time.Now().Format("20060102"),
		version,
	)

	err = pdfg.WriteFile(pdfName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

	return
}

func createPage(html string) *wkhtmltopdf.PageReader{
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	//pdfg.AddPage(page)


	// Set options for this page
	page.FooterRight.Set("[page]")
	page.MinimumFontSize.Set(12)
	page.FooterFontSize.Set(12)
	page.Zoom.Set(0.95)

	return page
}