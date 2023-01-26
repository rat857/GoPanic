package main

import (
	"gopanic/jiexi"
	"gopanic/osdo"
	"gopanic/yanzheng"
	"os"
)

func main() {
	osdo.Title()
	/*//输出参数
	outinfo := flag.String("o", "poc.yaml", "output info")
	flag.Parse()*/
	//初始化生成poc.yaml
	_, errPoc := os.Stat("poc.yaml")
	if errPoc != nil {
		osdo.WriteYaml(jiexi.PocModeYaml(), "poc.yaml")
	}

	//初始化生成pocMode.yaml
	_, errPocmode := os.Stat("pocMode.yaml")
	if errPocmode != nil {
		osdo.WriteYaml(jiexi.ModeYaml(), "pocMode.yaml")
	}

	//初始化生成config.yaml
	_, errConf := os.Stat("config.yaml")
	if errConf != nil {
		osdo.WriteYaml(jiexi.ModeConfig(), "config.yaml")
	}

	osdo.WriteListTxt(yanzheng.Important(), "good.txt")
}
