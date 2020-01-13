package eighth

/*
 * @Author: your name
 * @Date: 2020-01-12 21:33:47
 * @LastEditTime : 2020-01-12 22:14:48
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demos/eighth/homework08.go
 */

//Coders 攻城狮信息数组
type Coders struct {
	coders []Coder
}

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
