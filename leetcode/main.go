package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	//打包
	zipFile, err := os.Create("work1.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	imageURLs := []string{
		"https://picx.zhimg.com/70/v2-3b4fc7e3a1195a081d0259246c38debc_1440w.avis?source=172ae18b&biz_tag=Post",
		"https://pic2.zhimg.com/v2-df7fd62e3fffcd9c14c9f82db55a9a3d_r.jpg",
	}
	fileSlice := []string{}
	for i := 1; i <= len(imageURLs); i++ {
		filename := fmt.Sprintf("image%d.jpg", i)
		fileSlice = append(fileSlice, filename)
		response, err := http.Get(imageURLs[i-1])
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		zipFile, err := zipWriter.Create(filepath.Base(filename))
		if err != nil {
			panic(err)
		}

		// 将文件内容写入ZIP文件
		_, err = io.Copy(zipFile, response.Body)
		if err != nil {
			panic(err)
		}
	}
}
