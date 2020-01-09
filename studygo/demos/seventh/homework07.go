package seventh

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

/*翻译对象*/
type Translation struct {
	Src      string
	Target   string
	File string
}

/*翻译*/
func (trans Translation) Translate(src string) {
	trans.Src = src
	trans.Target = ""
	_, name, _, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fmt.Println("name = ", name)
	thePath := path.Join(path.Dir(name), trans.File)
	file, err := os.Open(thePath)
	if err != nil {
		fmt.Println("open err: ", err)
		return
	}
	defer file.Close()
	for {
		buf := make([]byte, 1024)
		_, err := file.Read(buf)
		if err == io.EOF || err != nil {
			fmt.Println("file err: ", err)
			return
		}
		str := string(buf)
		//fmt.Println(str)
		strs := strings.Split(str, "#")
		for _, val := range strs {
			//fmt.Println("str idx: ", i)
			fmt.Println("str val: ", val)
			if val == "#" + src {
				trans.Target = val
				return
			}
		}
	}
}
