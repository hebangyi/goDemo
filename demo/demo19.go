package demo

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//// 文件操作与文件的读取
// golang 中 os.File 封装了所有文件相关操作, File是一个结构体
func testFile() {

	// 打开文件
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
	}

	// 输出文件
	fmt.Printf("file=%v", file)

	// 关闭文件
	defer file.Close()
}

// 使用reader 带缓冲的方式读取文件
func testReader() {
	// Open opens the named file for reading. If successful, methods on
	// the returned file can be used for reading; the associated file
	// descriptor has mode O_RDONLY.
	file, _ := os.Open("d:/test.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读取到一个换行符就结束
		if err == io.EOF {
			break
		}
		fmt.Sprintf("str = %v", str)
	}

	fmt.Println("文件读取结束...")
}

// 一次性读取文件,适合文件不大的情况
func testReadFile() {
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file error !")
	}
	fmt.Println("content = ", content)         // 读取的byte 文件 []byte
	fmt.Println("content = ", string(content)) // 转换为string 类型
	// 因为我们没有显式Open,因此也不需要Close文件
	// 因为文件Open和Close被封装到 ReadFile 函数内部
}

//
func testOpenFile() {
	filePath := "d:/abc.txt"
	// O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	// O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	// O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// // The remaining values may be or'ed in to control behavior.
	// O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	// O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	// O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	// O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	// O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.

	// FileMode 只在Linux和Unix得文件中有用
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file err= %v\n", err)
		return
	}

	// 准备写入5句 "hello, Gardon"
	str := "hello,Gardon"
	// 写入时,使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer 是带缓存,因此在调用WriterString方法时,其实
	// 内容是先写入到缓存的,所以需要调用Flush方法,将缓存的数据真正写入到文件中
	// 否则文件中会丢失数据
	writer.Flush()
}

func testCopyFile() {
	file1Path := "d:/abc.txt"
	file2Path := "e:/kkk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		// 说明读取文件有错误
		fmt.Printf("read file err = %v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Println("写文件发生错误!")
		return
	}

	srcFile, err := os.Open(file1Path)
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(file2Path, os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	io.Copy(writer, reader)

}

// 判断文件是否存在
func testFileExist() {
	path := "d:/abc.txt"
	_, err := os.Stat(path)
	if err == nil {
		// 文件目录存在
		fmt.Println("文件或者目录存在")
		return
	}
	if os.IsNotExist(err) {
		fmt.Println("文件或者目录不存在")
		return
	}
	fmt.Println("发生未知错误,无法判断")
	return
}
