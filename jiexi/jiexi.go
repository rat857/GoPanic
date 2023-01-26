package jiexi

// 请求包的结构体
type Req struct {
	Type     string            //GET/POST
	Resource string            //Eg: /admin/login.php
	Head     map[string]string //"User-Agent"="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	Body     string            //Eg: username:admin&password:123
	Poc      string            //requests里有什么说明成功 Eg:login success
	Fofa     string
	Shodan   string
	Link     []string //相关链接地址
	Beizhu   string   //备注
}
type Conf struct {
	TimeOut int    //超时时间
	Proxy   string //代理
}

// 初始化
func PocModeYaml() Req {
	//生成poc
	var pocYaml Req
	return pocYaml
}

// 生成一个模板文件
func ModeYaml() Req {
	var modoYaml Req
	modoYaml.Type = "POST"
	modoYaml.Resource = "/login/verify"
	modoYaml.Head = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36", "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"}
	modoYaml.Body = "username=admin&password=123"
	modoYaml.Poc = "login success"
	modoYaml.Fofa = "nps admin login"
	modoYaml.Link = []string{"https://0x20h.com/p/eada.html"}
	modoYaml.Beizhu = "这是nps的默认密码爆破模块admin/123"
	return modoYaml
}

// 生成一个config文件
func ModeConfig() Conf {
	var modeConf Conf
	modeConf.TimeOut = 10
	modeConf.Proxy = "http://127.0.0.1:8083"
	return modeConf
}
