package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 准备打包的源文件
	var srcFile = "passwd"
	// 打包后的文件
	var desFile = fmt.Sprintf("%s.tar", srcFile)

	// 需要注意文件的打开即关闭的顺序，因为 defer 是后入先出，所以关闭顺序很重要
	// 第一次写这个示例的时候就没注意，导致写完的 tar 包不完整

	// ###### 第 1 步，先准备好一个 tar.Writer 结构，然后再向里面写入内容。 ######
	// 创建一个文件，用来保存打包后的 passwd.tar 文件
	fw, err := os.Create(desFile)
	ErrPrintln(err)
	defer fw.Close()

	// 通过 fw 创建一个 tar.Writer
	tw := tar.NewWriter(fw)
	// 这里不要忘记关闭，如果不能成功关闭会造成 tar 包不完整
	// 所以这里在关闭的同时进行判断，可以清楚的知道是否成功关闭
	defer func() {
		if err := tw.Close(); err != nil {
			ErrPrintln(err)
		}
	}()

	// ###### 第 2 步，处理文件信息，也就是 tar.Header 相关的 ######
	// tar 包共有两部分内容：文件信息和文件数据
	// 通过 Stat 获取 FileInfo，然后通过 FileInfoHeader 得到 hdr tar.*Header
	fi, err := os.Stat(srcFile)
	ErrPrintln(err)
	hdr, err := tar.FileInfoHeader(fi, "")
	// 将 tar 的文件信息 hdr 写入到 tw
	err = tw.WriteHeader(hdr)
	ErrPrintln(err)

	// 将文件数据写入
	// 打开准备写入的文件
	fr, err := os.Open(srcFile)
	ErrPrintln(err)
	defer fr.Close()

	written, err := io.Copy(tw, fr)
	ErrPrintln(err)

	log.Printf("共写入了 %d 个字符的数据\n", written)
}

// 定义一个用来打印的函数，少写点代码，因为要处理很多次的 err
// 后面其他示例还会继续使用这个函数，就不单独再写，望看到此函数了解
func ErrPrintln(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
