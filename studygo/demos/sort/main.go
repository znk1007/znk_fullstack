/*
 * @Author: your name
 * @Date: 2019-12-19 22:27:34
 * @LastEditTime : 2019-12-19 22:45:39
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demos/sort/main.go
 */
package main

/*
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install github.com/josharian/impl
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install github.com/golang/lint/golint
go install github.com/cweill/gotests/gotests
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install github.com/josharian/impl
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install github.com/cweill/gotests/gotests
*/
 
func main() {
	
}


func swap(lhs *int, rhs *int)  {
	var temp = *lhs;
	*lhs = *rhs;
	*rhs = temp;
}