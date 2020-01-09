package seventh

import (
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
	result   chan Result
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
		result:   make(chan Result),
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
			fmt.Println("middleStr: ===", middleStr, len(strs))
			
			if len(strs) >= 2 {
				fmt.Println("strs 0:", strs[0])
				fmt.Println("strs 1:", strs[1])
				trans.result <- Result{
					Src:    strs[0],
					Target: strs[1],
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
func (trans *Translation) Result() Result {
	//result := <-trans.result
	//fmt.Println(result)
	//return result
	go func() {
		for {
			select {
			case result := <-trans.result:
				fmt.Println("final result: ", result)
				//return result
			case <-trans.quit:
				return
				//return Result{
				//	Src:    "",
				//	Target: "",
				//}
			}
		}
	}()
	return Result{}
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
				//fmt.Println("find val: ", val)
				//fmt.Println("find src: ", src)
				return
			}
		}
	}
}
