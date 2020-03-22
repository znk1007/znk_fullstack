package eleventh

/*
#include <stdio.h> 
extern void CSnake();
void createGame(){
	printf("创建贪吃蛇游戏\n");
}
void startGame() {
	printf("开始贪吃蛇游戏\n");
}
void runGame(){
	printf("执行贪吃蛇游戏\n");
}
void endGame(char *s) {
	printf("结束贪吃蛇游戏 %s\n",s);
}
*/
import "C"
//SnakeGame 贪吃蛇游戏
func SnakeGame()  {
	
	C.createGame()
	cStr := C.CString("自爆了\n")
	C.endGame(cStr)
}
//ExportSnakeGame 暴露贪吃蛇
func ExportSnakeGame(str *C.char) {
	C.CSnake()
	fmt.Println("ExportSnake!")
    fmt.Println(C.GoString(str))
}
