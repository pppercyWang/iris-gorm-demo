/*
@Time : 2019/5/17 10:26 
@Author : Lukebryan
@File : excelutil.go
@Software: GoLand
*/
package utils

//读取Excel
//func ReadExcel(path string, fileName string) ([][]string) {
//	var data [][]string
//	xlsx, err := excelize.OpenFile(path + fileName)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//	// Get value from cell by given worksheet name and axis.
//	xlsx.GetCellValue("Sheet1", "B2")
//	// Get all the rows in the Sheet1.
//	rows := xlsx.GetRows("Sheet1")
//	for rindex, row := range rows {
//		if rindex < 5 {
//			continue
//		}
//		var rdata []string
//		for cindex, colCell := range row {
//			if cindex < 1 {
//				continue
//			}
//			rdata = append(rdata, colCell)
//		}
//		data = append(data, rdata)
//	}
//	return data
//}
