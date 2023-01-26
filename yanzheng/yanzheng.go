package yanzheng

import (
	"crypto/tls"
	"flag"
	"fmt"
	"gopanic/osdo"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// yaml配对请求
func Important() []string {
	//指定FILEorURL
	pocYaml := flag.String("v", "pocMode.yaml", "指定一个pocYAML文件")
	UrlString := flag.String("u", "nul", "指定一个url")
	UrlFile := flag.String("f", "nul", "指定一个存放了url的文件")
	flag.Parse()
	var urls = make([]string, 0)
	if (*UrlString) != "nul" && (*UrlFile) == "nul" {
		urls = append(urls, *UrlString)
	} else if *UrlFile != "nul" && (*UrlString) == "nul" {
		urls = osdo.ReadTxtList(*UrlFile)
	} else {
		fmt.Println("error")
	}

	pocinfo := osdo.ReadYamlReq(*pocYaml)
	config := osdo.ReadYamlConf("config.yaml")
	var bodList = make([]string, 0)
	if pocinfo.Type == "POST" || pocinfo.Type == "GET" {
		var wg sync.WaitGroup
		var tr *http.Transport

		proxList := strings.Split(config.Proxy, "://")

		if config.Proxy != "" { //判断proxy是否有值
			tr = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},                           //防止https证书错误的问题
				Proxy:           http.ProxyURL(&url.URL{Scheme: proxList[0], Host: proxList[1]}), //代理
			}
		} else {
			tr = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //防止https证书错误的问题
			}
		}
		client := &http.Client{Transport: tr, Timeout: time.Second * time.Duration(config.TimeOut)} //超时时间
		for i := 0; i < len(urls); i++ {
			wg.Add(1)
			go func(i int) {
				fmt.Println("\033[32m第", i, "个")
				defer wg.Done()
				data := pocinfo.Body
				boby := strings.NewReader(data)
				req, _ := http.NewRequest(pocinfo.Type, urls[i]+pocinfo.Resource, boby)
				for key, vule := range pocinfo.Head {
					req.Header.Set(key, vule)
				}
				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				defer resp.Body.Close()
				bodyByte, _ := io.ReadAll(resp.Body)
				bodyStr := string(bodyByte)
				if strings.Contains(bodyStr, pocinfo.Poc) {
					bodList = append(bodList, urls[i])
					fmt.Printf("\033[31m%s\n", "have the vule")
				}
			}(i)
		}
		wg.Wait()
	} else {
		fmt.Println("oh sorry!,暂时不支持该类型")
	}
	return bodList
}
