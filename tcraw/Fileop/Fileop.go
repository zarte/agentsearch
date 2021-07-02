package Fileop

import (
	"os"
	"fmt"
	"sync"
	"io"
)


var Fileoper *Fileop

type Fileop struct{
	path string
	mu     sync.Mutex // ensures atomic writes; protects the following fields
}



func NewLog(path string) *Fileop {

	//创建日志目录
	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf("crete fail %s", err)
		os.Exit(1)
	} else {
		// fmt.Print("Create Directory OK!")
	}

	//返回日志对象
	Log := &Fileop{
		path : path+"/",
	}
	return Log
}


func (a *Fileop) Adddata(v ...interface{}){
	var f    *os.File
	var err   error

	var filename = "common.data"
	a.mu.Lock()
	defer a.mu.Unlock()

	if checkFileIsExist(a.path+filename) {  //如果文件存在
		f, err = os.OpenFile(a.path+filename, os.O_APPEND,0)  //打开文件
	}else {
		f, err = os.Create(a.path+filename)  //创建文件
	}
	check(err)
	defer f.Close()
	s := fmt.Sprintln(v...)
	_,err = io.WriteString(f, s) //写入文件(字符串)
	check(err)

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Checkexist(path string)  bool{
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return false
}
/**
* 判断文件是否存在  存在返回 true 不存在返回false
*/
func checkFileIsExist(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return false
}
