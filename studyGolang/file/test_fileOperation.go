package main

import (
	"fmt"
	"path/filepath"
	"path"
)
func main() {
	/*
	文件操作：
	1.路径：
		相对路径：relative
			test.txt
			相对于当前工程
		绝对路径：absolute
			/Users/klayhu/Desktop/vscode_project/Go/studyGolang/file/test.txt

		.当前目录
		..上一层
	2.创建文件夹，如果文件夹存在，创建失败
		os.MkDir()，创建一层
		os.MkDirAll()，可以创建多层

	3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
		os.Create()，创建文件

	4.打开文件：让当前的程序，和指定的文件之间建立一个连接
		os.Open(filename)
		os.OpenFile(filename,mode,perm)

	5.关闭文件：程序和文件之间的链接断开。
		file.Close()

	5.删除文件或目录：慎用，慎用，再慎用
		os.Remove()，删除文件和空目录
		os.RemoveAll()，删除所有
	 */
	// 路径
	fileName1:="/Users/klayhu/Desktop/vscode_project/Go/studyGolang/file/test.txt" // 不同机器不同
	fileName2:="test2.txt"
	fmt.Println(filepath.IsAbs(fileName1)) //判断是否是绝对路径 true
	fmt.Println(filepath.IsAbs(fileName2)) //false
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2)) //

	fmt.Println("获取父目录：",path.Join(fileName1,".."))

	//2.创建目录
	/*
	 err := os.Mkdir("/Users/klayhu/Desktop/vscode_project/Go/studyGolang/file/filefirst",os.ModePerm)
	 if err != nil{
	 	fmt.Println("err:",err)
	 	return
	 }
	 fmt.Println("文件夹创建成功。。")
	 err :=os.MkdirAll("/Users/klayhu/Desktop/vscode_project/Go/studyGolang/file/filefirst/filesecond",os.ModePerm)
	 if err != nil{
	 	fmt.Println("err:",err)
	 	return
	 }
	 fmt.Println("多层文件夹创建成功")
	 */

	//3.创建文件:Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	/*
	file1,err :=os.Create("/Users/klayhu/Desktop/vscode_project/Go/studyGolang/file/filefirst/test1.txt")
	if err != nil{
		fmt.Println("err：",err)
		return
	}
	fmt.Println(file1)

	file2,err := os.Create(fileName2)//创建相对路径的文件，是以当前工程为参照的
	if err != nil{
		fmt.Println("err :",err)
		return
	}
	fmt.Println(file2)
	*/

	//4.打开文件：
	/*
	 //file3 ,err := os.Open(fileName1) //只读的
	 //if err != nil{
	 //	fmt.Println("err:",err)
	 //	return
	 //}
	 //fmt.Println(file3)
	/*
	第一个参数：文件名称
	第二个参数：文件的打开方式
		const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	)
	第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	 //file4,err := os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
	 //if err != nil{
	 //	fmt.Println("err:",err)
	 //	return
	 //}
	 //fmt.Println(file4)
	*/

	 //5.关闭文件
	 //file4.Close()

	 //6.删除文件或文件夹：
	 //删除文件
	//err :=  os.Remove("/Users/ruby/Documents/pro/a/aa.txt")
	//if err != nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("删除文件成功。。")
}
