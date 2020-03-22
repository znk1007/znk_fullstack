package seventh

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

type Result struct {
	Src    string
	Target string
}

/*翻译对象*/
type Translation struct {
	middle   chan string
	result   func(string, string)
	quit     chan bool
	filePath string
}

/*初始化翻译对象*/
func CreateTranslation(fileName string) *Translation {
	_, name, _, ok := runtime.Caller(1)
	if !ok {
		return nil
	}
	//fmt.Println("name = ", name)
	thePath := path.Join(path.Dir(name), fileName)
	var trans = &Translation{
		middle:   make(chan string),
		result:   nil,
		quit:     make(chan bool),
		filePath: thePath,
	}
	go trans.run()
	return trans
}

func (trans *Translation) run() {
	for {
		select {
		case middleStr := <-trans.middle:
			strs := strings.Split(middleStr, "\n")
			//fmt.Println("middleStr: ===", middleStr, len(strs))
			if len(strs) >= 2 {
				//fmt.Println("strs 0:", strs[0])
				//fmt.Println("strs 1:", strs[1])
				if trans.result != nil {
					trans.result(strs[0], strs[1])
				}
			} else {
				return
			}
		case <-trans.quit:
			return
		}
	}
}

/*翻译结果*/
func (trans *Translation) GetResult(result func(src string, result string)) {
	trans.result = result
}

/*翻译*/
func (trans *Translation) Translate(src string) {
	if trans == nil {
		return
	}
	file, err := os.Open(trans.filePath)
	if err != nil {
		fmt.Println("open err: ", err)
		trans.quit <- true
		return
	}
	defer file.Close()
	for {
		buf := make([]byte, 1024)
		_, err := file.Read(buf)
		if err == io.EOF || err != nil {
			fmt.Println("file err: ", err)
			trans.quit <- true
			return
		}
		str := string(buf)
		//fmt.Println(str)
		strs := strings.Split(str, "#")
		for _, val := range strs {
			//fmt.Println("str val: ", val)
			if strings.HasPrefix(val, src) {
				trans.middle <- val
				//time.Sleep(time.Second)
				return
			}
		}
	}
}

func (trans *Translation) Translate2(src string, result func(src string, target string)) {
	if trans == nil {
		return
	}
	file, err := os.Open(trans.filePath)
	if err != nil {
		fmt.Println("open err: ", err)
		if result != nil {
			result("", "")
		}
		return
	}
	defer file.Close()
	for {
		buf := make([]byte, 1024)
		n, err := file.Read(buf)
		if err != io.EOF && err != nil {
			fmt.Println("file err: ", err)
			if result != nil {
				result("", "")
			}
			return
		}
		if n == 0 {
			trans.quit <- true
			break
		}
		str := string(buf)
		//fmt.Println(str)
		strs := strings.Split(str, "#")
		for _, val := range strs {
			//fmt.Println("str val: ", val)
			if strings.HasPrefix(val, src) {
				if result != nil {
					strs := strings.Split(val, "\n")
					if len(strs) >= 2 {
						if result != nil {
							result(strs[0], strs[1])
						}
					} else {
						if result != nil {
							result("", "")
						}
					}
				}
				return
			}
		}
	}
}

/*拷贝*/
type Copy struct {
	srcPath string
	dstPath string
	bytes   chan []byte
	quit    chan bool
}

/*创建拷贝对象*/
func CreateCopy(srcPath string, dstPath string) Copy {
	c := Copy{
		srcPath: srcPath,
		dstPath: dstPath,
		bytes:   make(chan []byte, 1024),
		quit:    make(chan bool),
	}
	go c.write()
	return c
}

/*写入数据*/
func (c Copy) write() {
	dstFile, err := os.Create(c.dstPath)
	if err != nil {
		fmt.Println("write open err: ", err.Error())
		return
	}
	defer dstFile.Close()
	for {
		select {
		case buf := <-c.bytes:
			dstFile.Write(buf)
		case <-c.quit:
			return
		}
	}
}

func (c Copy) Copy() error {
	if len(c.srcPath) == 0 || len(c.dstPath) == 0 || c.dstPath == c.srcPath {
		c.quit <- true
		return errors.New("源⽂件和⽬的⽂件名字不能相同")
	}
	srcFile, err := os.Open(c.srcPath)
	if err != nil {
		c.quit <- true
		return err
	}
	defer srcFile.Close()
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			c.quit <- true
			break
		}
		if n == 0 {
			c.quit <- true
			fmt.Println("copy successfully")
			break
		}
		c.bytes <- buf[:n]
	}
	return nil
}
