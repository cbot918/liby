package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	// 测试的 URL
	strURL := "https://www.ucg.ac.me/skladiste/blog_44233/objava_64433/fajlovi/Computer%20Networking%20_%20A%20Top%20Down%20Approach,%207th,%20converted.pdf"

	resp, err := http.Head(strURL)
	if err != nil {
		fmt.Println("resp, err := http.Head(strURL)  报错: strURL = ", strURL)
		log.Fatalln(err)
	}

	// fmt.Printf("%#v\n", resp)
	fileLength := int(resp.ContentLength)

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", 0, fileLength))
	// fmt.Printf("%#v", req)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient.Do(req)", "error")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 创建文件
	filename := path.Base(strURL)
	flags := os.O_CREATE | os.O_WRONLY
	f, err := os.OpenFile(filename, flags, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		log.Fatal("err")
	}
	defer f.Close()

	// 写入数据
	buf := make([]byte, 16*1024)
	_, err = io.CopyBuffer(f, resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println("io.EOF")
			return
		}
		fmt.Println(err)
		log.Fatal(err)
	}

}
