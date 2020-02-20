package fourteenth

import (
	// "crypto/rand"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"os"
	"time"
)

//User 用户模板
type User struct {
	Name string
	Age  int
}

//StringLiteralTemplate 字符串模板
func StringLiteralTemplate() {
	s := "My name is {{.Name}}. I am {{.Age}} year old.\n"
	t, err := template.New("testuser").Parse(s)
	if err != nil {
		log.Fatal("Parse string literal template error: ", err)
	}
	u := User{
		Name: "lianshi",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute string literal template error: ", err)
	}
}

//FileTemplate 文件模板 /Users/huangsam/Desktop/golang/src/
func FileTemplate(fp string) {
	t, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal("Parse file template error: ", err)
	}
	u := User{
		Name: "ls",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute file template error: ", err)
	}
	fmt.Println("")
}

func (u User) String() string {
	return fmt.Sprintf("(name:%s age: %d)", u.Name, u.Age)
}

//DotActionTemplate 点动作
func DotActionTemplate() {
	s := "The user is {{.}}"
	t, err := template.New("test1").Parse(s)
	if err != nil {
		log.Fatal("Parse error: ", err)
	}
	u := User{
		Name: "lianshi",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
}

//AgeInfo 年龄信息
type AgeInfo struct {
	Age           int
	GreaterThan60 bool
	GreaterThan40 bool
}

//AgeInfoTemplate 年龄信息模板
func AgeInfoTemplate(fp string) {
	t, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal("Parse file template error: ", err)
	}
	rand.Seed(time.Now().Unix())
	age := rand.Intn(100)
	info := AgeInfo{
		Age:           age,
		GreaterThan60: age > 60,
		GreaterThan40: age > 40,
	}
	err = t.Execute(os.Stdout, info)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
}

//Item ...
type Item struct {
	Name  string
	Price int
}

//ItemTemplate Item模板
func ItemTemplate(file string) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal("Parse error: ", err)
		return
	}
	items := []Item{
		{"iPhone", 699},
		{"iPad", 799},
		{"iWatch", 199},
		{"MacBook", 999},
	}
	err = t.Execute(os.Stdout, items)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
}

//Pet 宠物
type Pet struct {
	Name  string
	Age   int
	Owner User
}

//PetTemplate 宠物模板
func PetTemplate(file string) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal("Parse error: ", err)
	}
	p := Pet{
		Name: "Orange",
		Age:  2,
		Owner: User{
			Name: "ls",
			Age:  18,
		},
	}
	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal("Execute error: ", err)
	}
	fmt.Println("")
}

//NestTemplate 包含动作模板
func NestTemplate(file1 string, file2 string) {
	t, err := template.ParseFiles(file1, file2)
	if err != nil {
		log.Fatal("nest parse err: ", err)
	}
	err = t.Execute(os.Stdout, "test data")
	if err != nil {
		log.Fatal("nest execute error: ", err)
	}
}
