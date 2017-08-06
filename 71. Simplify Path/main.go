package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "/a/./b/./../c/"
	fmt.Println(simplifyPath(testString))
}

func simplifyPath(path string) string {
	var pathRecord [] string
	nowPath := path
	//1 先按照'/'将字符串转成切片
	s := strings.FieldsFunc(nowPath, split)

	//2 遍历切片,按不同的字符进行处理
	for _, value := range s {
		if value == "." {
			//2.1 如果只有一个'.',则直接跳过
			continue
		}else if value == ".." {
			//2.2 如果有两个'.',并且当前不是/路径,则到上一级路径
			if len(pathRecord) < 1 {
				continue
			}
			pathRecord = pathRecord[0:len(pathRecord) - 1]
		}else{
			//2.3 如果是其他字符,则认为是一级目录,进入目录
			pathRecord = append(pathRecord, value)
		}
	}
	//3 将切片按'/'连接,并在最开始拼上'/'
	return "/" + strings.Join(pathRecord, "/")
}

func split(s rune) bool {
	return s == '/';
}