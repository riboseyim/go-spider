package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func read_csv_caipiao(filename string, sep string) []string {
	fmt.Printf("Input file name :%s \n", filename)
	cntb, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	//fmt.Println(ss)
	sz := len(ss)
	data := make([]string, sz)
	for i := 0; i < sz; i++ {
		fmt.Println(ss[i][3])
		data[i] = ss[i][3]
	}
	return data
}

func read_csv_normal(filename string) []string {
	fmt.Printf("Input file name :%s \n", filename)
	cntb, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	//fmt.Println(ss)
	sz := len(ss)
	data := make([]string, sz)
	for i := 0; i < sz; i++ {
		for j := 0; j < len(ss[i]); j++ {
			fmt.Println(ss[i][j])
			data[i] = ss[i][j]
		}

	}
	return data
}

func clear_csv(filename string) {

	os.Remove(filename)

}

func save_csv(filename string, data [][]string, append bool) {

	if !append {
		os.Remove(filename)
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600) //创建文件

	if err != nil {
		panic(err)
	}
	defer f.Close()

	//f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	w.WriteAll(data)      //写入数据
	w.Flush()
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
