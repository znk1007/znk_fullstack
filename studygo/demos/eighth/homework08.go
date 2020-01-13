package eighth

import "github.com/360EntSecGroup-Skylar/excelize"

import "fmt"

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
	coders    []Coder
	writeXLSX bool
	coder     chan Coder
	quit      chan bool
}

//CreateCoderInfo 创建CrateCoderInfo对象
func CreateCoderInfo(writeXLSX bool) CoderInfo {
	return CoderInfo{
		coders:    []Coder{},
		writeXLSX: writeXLSX,
		coder:     make(chan Coder),
		quit:      make(chan bool),
	}
}

func (info CoderInfo) write() {
	// go func() {
	for {
		select {
		case coder := <-info.coder:
			info.coders = append(info.coders, coder)
		case <-info.quit:
			return
		}
	}
	// }()
}

func (info CoderInfo) Read(filePath string) {
	if len(filePath) == 0 {
		info.quit <- true
		return
	}
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("open file: ", err.Error())
		info.quit <- true
		return
	}
	rows, err := f.GetRows("工作表 1 - 2019-12-17 资深Go语言工程师实战课")
	for _, row := range rows {
		// for _, colCell := range row {
		// 	fmt.Println(colCell)
		// }
		fmt.Println(row)
	}
}
