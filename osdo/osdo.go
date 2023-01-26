package osdo

import (
	"bufio"
	"fmt"
	"gopanic/jiexi"
	"gopkg.in/yaml.v2"
	"os"
)

// 写入Yaml
func WriteYaml(datas interface{}, filename string) { //第一个参数是要写入的数据，第二个是生成的文件的名字
	data, err := yaml.Marshal(datas)
	if err != nil {
		panic(err)
	}
	os.WriteFile(filename, data, 0666)
}

// 读取Yaml为jiexi的Req结构体
func ReadYamlReq(fileName string) jiexi.Req { //第一个参数是文件名，返回值是jiexi.Req
	res, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var data jiexi.Req
	yaml.Unmarshal(res, &data)
	return data
}

// 写入文件（把List按行写入）
func WriteListTxt(resList []string, fileName string) { //第一个参数是写入的url的list,第二个参数生成的文件名
	//创建文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//创建 bufio.Writer
	writer := bufio.NewWriter(file)

	//循环写入文件
	for _, line := range resList {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	//刷新缓存
	writer.Flush()
}

// 按行读取txt文件，并保存在一个List
func ReadTxtList(fileName string) []string { //第一个参数要读取的文件名字
	var urls = make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	a := bufio.NewReader(file)
	for true {
		data, _, error := a.ReadLine()
		if error != nil {
			break
		} else {
			urls = append(urls, string(data))
		}
	}
	return urls
}

// 读取Yaml为jiexi.Conf结构体
func ReadYamlConf(fileName string) jiexi.Conf { //第一个参数是文件名，返回值是jiexi.Conf
	res, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var data jiexi.Conf
	yaml.Unmarshal(res, &data)
	return data
}
func Title() {
	a := ` 
 ██████╗  ██████╗ ██████╗  █████╗ ███╗   ██╗██╗ ██████╗
██╔════╝ ██╔═══██╗██╔══██╗██╔══██╗████╗  ██║██║██╔════╝
██║  ███╗██║   ██║██████╔╝███████║██╔██╗ ██║██║██║     
██║   ██║██║   ██║██╔═══╝ ██╔══██║██║╚██╗██║██║██║     
╚██████╔╝╚██████╔╝██║     ██║  ██║██║ ╚████║██║╚██████╗
 ╚═════╝  ╚═════╝ ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝ ╚═════╝
                                                  @rat857      

`
	fmt.Println(a)
}
