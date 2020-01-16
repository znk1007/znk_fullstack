package eighth

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
 * @Author: your name
 * @Date: 2020-01-12 21:33:47
 * @LastEditTime : 2020-01-13 21:46:28
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demos/eighth/homework08.go
 */

//Coder 攻城狮信息
type Coder struct {
	Name      string //名称
	Education string //学历
	School    string // 学校
	Major     string // 专业
	Years     string // 工作年限
	Job       string // 职位
	Salary    string // 薪水
	Language  string // 编程语言
}

//CoderInfo 攻城狮信息数组
type CoderInfo struct {
	coders  []Coder
	coder   chan Coder
	quit    chan bool
	srcPath string
	dstPath string
}

//CreateCoderInfo 创建CrateCoderInfo对象
func CreateCoderInfo(srcPath, dstPath string) CoderInfo {
	info := CoderInfo{
		coders:  []Coder{},
		coder:   make(chan Coder),
		quit:    make(chan bool),
		srcPath: srcPath,
		dstPath: dstPath,
	}
	go info.write()
	return info
}

func (info CoderInfo) write() {
	for {
		select {
		case coder := <-info.coder:
			info.coders = append(info.coders, coder)
		case <-info.quit:
			info.save()
			return
		}
	}
}

func (info CoderInfo) axisA(idx int) string {
	return fmt.Sprintf("A%d", idx)
}

func (info CoderInfo) axisB(idx int) string {
	return fmt.Sprintf("B%d", idx)
}

func (info CoderInfo) axisC(idx int) string {
	return fmt.Sprintf("C%d", idx)
}

func (info CoderInfo) axisD(idx int) string {
	return fmt.Sprintf("D%d", idx)
}

func (info CoderInfo) axisE(idx int) string {
	return fmt.Sprintf("E%d", idx)
}

func (info CoderInfo) axisF(idx int) string {
	return fmt.Sprintf("G%d", idx)
}

func (info CoderInfo) axisG(idx int) string {
	return fmt.Sprintf("G%d", idx)
}

func (info CoderInfo) axisH(idx int) string {
	return fmt.Sprintf("H%d", idx)
}

//ReadAndSave 读取xls文件数据，匹配字符串，保存到另外的xls文件
func (info CoderInfo) ReadAndSave(matchStr string) {
	if len(info.srcPath) == 0 {
		if quit := <-info.quit; quit == false {
			info.quit <- true
		}
		return
	}
	f, err := excelize.OpenFile(info.srcPath)
	if err != nil {
		fmt.Println("open file: ", err.Error())
		go func() {
			info.quit <- true
		}()
		return
	}
	rows, err := f.Rows("工作表 1 - 2019-12-17 资深Go语言工程师实战课")
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			go func() {
				info.quit <- true
			}()
		}
		var match bool
		for _, colCell := range row {
			if strings.Contains(colCell, matchStr) {
				match = true
			}
		}
		if match && len(row) > 7 {
			coder := Coder{
				Name:      row[0],
				Education: row[1],
				School:    row[2],
				Major:     row[3],
				Years:     row[4],
				Job:       row[5],
				Salary:    row[6],
				Language:  row[7],
			}
			info.coder <- coder
		}
	}
	info.quit <- true
}

//Save 保存数据
func (info CoderInfo) save() {
	f := excelize.NewFile()
	sheetName := "工作表 1 - 2019-12-17 资深Go语言工程师实战课"
	index := f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "微信昵称")
	f.SetCellValue(sheetName, "B1", "最高学历")
	f.SetCellValue(sheetName, "C1", "毕业学校")
	f.SetCellValue(sheetName, "D1", "所处行业")
	f.SetCellValue(sheetName, "E1", "工作年限")
	f.SetCellValue(sheetName, "F1", "目前职位")
	f.SetCellValue(sheetName, "G1", "目前月薪（单位：K）")
	f.SetCellValue(sheetName, "H1", "编程语言基础")
	for idx, coder := range info.coders {
		f.SetCellValue(sheetName, info.axisA(idx+2), coder.Name)
		f.SetCellValue(sheetName, info.axisB(idx+2), coder.Education)
		f.SetCellValue(sheetName, info.axisC(idx+2), coder.School)
		f.SetCellValue(sheetName, info.axisD(idx+2), coder.Major)
		f.SetCellValue(sheetName, info.axisE(idx+2), coder.Years)
		f.SetCellValue(sheetName, info.axisF(idx+2), coder.Job)
		f.SetCellValue(sheetName, info.axisG(idx+2), coder.Salary)
		f.SetCellValue(sheetName, info.axisH(idx+2), coder.Language)

	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(info.dstPath); err != nil {
		fmt.Println("save err: ", err.Error())
	}

}
