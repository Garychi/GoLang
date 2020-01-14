package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

const (
	fileName = "MyXLSXFile.xlsx"
)

func main() {
	createExcel()
	//readExcel()
}

func createExcel(){
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var titleRow *xlsx.Row
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	options := GetSecurityOption()
	title := getTitleList(options)
	rawData := getSecurityDataToList(options)

	titleRow = sheet.AddRow()
	for i:= range title{
		cell = titleRow.AddCell()
		cell.Value = title[i]
	}

	for i := range rawData {
		row = sheet.AddRow()
		tmp := rawData[i]

		datSize := len(tmp)
		for j := 0; j < datSize; j++ {
			cell = row.AddCell()
			cell.Value = tmp[j]
		}
	}


	err = file.Save(fileName)
	if err != nil {
		fmt.Println("has error")
		fmt.Printf(err.Error())
	}
}



func getTitleList(options []SecurityOption)(title []string){
	title = append(title, "項次")
	title = append(title, "工地編號")
	title = append(title, "工地名稱")
	for i:= range options{
		title = append(title, options[i].Title)
	}

	return
}

func GetSecurityOption()(options []SecurityOption){
	options = append(options, SecurityOption{ID:1,Title:"類別一"})
	options = append(options, SecurityOption{ID:2,Title:"類別二"})
	options = append(options, SecurityOption{ID:3,Title:"類別三"})
	options = append(options, SecurityOption{ID:4,Title:"類別四"})
	return
}

func getSecurityDataToList(options []SecurityOption)(data [][]string){
	rawData :=[]SecurityExcel{}
	rawData = append(rawData, SecurityExcel{ID:1,SiteCode:"9487",SiteName:"台康生技",SecurityOptionId:1})
	rawData = append(rawData, SecurityExcel{ID:2,SiteCode:"9405",SiteName:"長榮航勤",SecurityOptionId:2})
	rawData = append(rawData, SecurityExcel{ID:3,SiteCode:"9476",SiteName:"巨大總部",SecurityOptionId:3})
	rawData = append(rawData, SecurityExcel{ID:4,SiteCode:"9486",SiteName:"友勁科技",SecurityOptionId:4})

	optionsSize := len(options)

	for i:= range rawData{
		d :=[]string{}
		d = append(d, strconv.FormatInt(rawData[i].ID, 10))
		d = append(d, rawData[i].SiteCode)
		d = append(d, rawData[i].SiteName)

		// 動態類別
		for j := 0; j < optionsSize; j++ {
			if rawData[i].SecurityOptionId == options[j].ID {
				d = append(d, "✓")
			} else {
				d = append(d, "")
			}
		}

		data = append(data, d)

	}

	return
}

type SecurityExcel struct {
	ID					int64
	SiteCode 			string
	SiteName			string
	SecurityOptionId	int64
}

//func readExcel() {
//	excelFileName := fileName
//	xlFile, err := xlsx.OpenFile(excelFileName)
//	if err != nil {
//		fmt.Printf(err.Error())
//	}
//	for _, sheet := range xlFile.Sheets {
//		for _, row := range sheet.Rows {
//			for _, cell := range row.Cells {
//				text := cell.String()
//				fmt.Printf("%s\n", text)
//			}
//		}
//	}
//}