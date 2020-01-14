package util

import (
	"github.com/signintech/gopdf"
)

func aa() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	_ = pdf.SetFont("auto", "", 16)
	//err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
	//if err != nil {
	//	log.Print(err.Error())
	//	return
	//}

	//err = pdf.SetFont("wts11", "", 14)
	//if err != nil {
	//	log.Print(err.Error())
	//	return
	//}

	//rec :=gopdf.Rect{W:30,H:20}
	//var a = new(gopdf.Rect)
	//a.W=30
	//a.H =10
	//a.UnitsToPoints(10)
	//a.PointsToUnits(10)


	//pdf.Cell(a, "您好")

	//pdf.WritePdf("hello.pdf")
	//fmt.Println("done")
}
