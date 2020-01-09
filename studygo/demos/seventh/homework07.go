package seventh

import (
	"fmt"
	"io"
	"os"
)

/*翻译对象*/
type Translation struct {
	Src    string
	Target string
}

/*翻译*/
func (trans Translation) Translate(src string) {
	trans.Src = src
	trans.Target = ""
	file, err := os.Open("/dict.txt")
	if err != nil {
		fmt.Println("open err: ", err)
		return
	}
	for {
		buf := make([]byte, 1024)
		_, err := file.Read(buf)
		if err == io.EOF || err != nil {
			fmt.Println("file err: ", err)
			return
		}
		str := string(buf)
		fmt.Println(str)
	}
}
