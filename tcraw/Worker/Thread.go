package Worker

import (
	"agentsearch/tcraw/Config"
	"agentsearch/tcraw/Fileop"
	"agentsearch/tcraw/Logger"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var RunPageNums int = 0
var SuccessRunNums int = 0
var SuccessChannel = make(chan bool)
var RunPageNumsChannel = make(chan int)

var Totalpagenums int = 0


var chthreadnum = make(chan int)
var totalchthread int

var datereg *regexp.Regexp
var yesterday int64


func ThreadRun() {
    testproxy()
}

type ProxyIp struct {
    Port          int             `json:"port"`
    Host        string          `json:"host"`
    Type        string          `json:"type"`
    Country        string          `json:"country"`
    From        string          `json:"from"`
	Anonymity   string 			 `json:"anonymity"`
}

type Proxitem struct {
    Url    string
    Type   string
    Port   string
    Country string
	Anonymity string
}

var Proxlist chan Proxitem

var tProxitem map[string]interface{}

func testproxy()  {

    Proxlist = make(chan Proxitem,20)
    chthreadnum = make(chan int)

   totalnum := 0
   var proxyIp ProxyIp

	//生成等待协程
	for w := 1; w <= 10; w++ {
		totalchthread++
		go testproxychild(Proxlist)
	}

    if Fileop.Checkexist(Config.ConfigValue.Path+"proxlist.data") {
		f, err := os.OpenFile("./proxlist.data", os.O_RDONLY,0)  //打开文件
		if err != nil {
			fmt.Println(err)
		}else{
			defer f.Close()
			reader := bufio.NewReader(f)
			for {
				line,_,err :=reader.ReadLine()
				if err!=nil {
					fmt.Println(err)
					break
				}else{
					item := strings.Fields(string(line))
					if len(item) < 4{
						continue
					}
					if Config.ConfigValue.Ptype != "" && Config.ConfigValue.Ptype != item[3]{
						continue
					}
					if Config.ConfigValue.Pcountry != "" && Config.ConfigValue.Pcountry != item[2]{
						continue
					}



					var info Proxitem
					info.Url = item[0]
					info.Port = item[1]
					info.Type = item[3]
					info.Country = item[2]

					if len(item)>4{
						if Config.ConfigValue.Panonymity != "" && Config.ConfigValue.Panonymity != item[4]{
							continue
						}
						info.Anonymity = item[4]
					}


					totalnum++
					Proxlist <- info
				}
			}
		}
	}else{
		//获取代理列表
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/fate0/proxylist/master/proxy.list", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("user-agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
		var resp *http.Response
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		//将请求结果的byte数组放到responseBytes中
		reader := bufio.NewReader(resp.Body)

		for {
			line,_,err :=reader.ReadLine()
			if err!=nil {
				fmt.Println(err)
				break
			}else{
				err :=json.Unmarshal(line, &proxyIp)
			//	var result map[string]interface{}
				//err :=json.Unmarshal(line, &result)
				//s1 := int(result["port"].(float64))
				//fmt.Println(result["port"],"---",s1)
				if err !=nil {
					fmt.Println(err)
				}else{
					port := strconv.Itoa(proxyIp.Port)
					if proxyIp.Host == ""  || port ==""{
						continue
					}
					if Config.ConfigValue.Ptype != "" && Config.ConfigValue.Ptype != proxyIp.Type{
						continue
					}
					if Config.ConfigValue.Pcountry != "" && Config.ConfigValue.Pcountry != proxyIp.Country{
						continue
					}
					if Config.ConfigValue.Panonymity != "" && Config.ConfigValue.Panonymity != proxyIp.Anonymity{
						continue
					}

					var info Proxitem
					info.Url = proxyIp.Host
					info.Port = port
					info.Type = proxyIp.Type
					info.Country = proxyIp.Country
					info.Anonymity = proxyIp.Anonymity
					totalnum++
					Proxlist <- info
				}
			}
		}
	}


    close(Proxlist)


    //等待线程完成
    //等待协程结束
    var completenum int
    completenum = 0
    for {
        _ =<- chthreadnum
        completenum ++
        if completenum >= totalchthread{
            break
        }
    }
    fmt.Println("totalnum is:",totalnum)
}

func testproxychild(list <- chan Proxitem)  {
    defer func(){
        chthreadnum <- 1
    }()
    for item := range list{
            _, err := net.Dial("tcp",item.Url+":"+item.Port )
            if err == nil {
                //端口开放
				Logger.Debug(item.Url,"open")
                Fileop.Fileoper.Adddata(item.Url,item.Port ,item.Country,item.Type,item.Anonymity)
            }else{
				Logger.Debug(item.Url,"close")
            }
            /*
			   if proxyIp["host"] == ""  || proxyIp["port"] ==""{
				   continue
			   }
			   host := proxyIp["host"].(string)
			   port := strconv.FormatFloat(proxyIp["port"].(float64),'E',-1,64)
			   url := host+":"+ port
			   _, err = net.Dial("tcp",url )
			   if err == nil {
				   //端口开放
				   fmt.Println(proxyIp["host"],proxyIp["port"],"open")
				   Logger.ToFile(proxyIp["host"],proxyIp["port"],proxyIp["country"],proxyIp["type"])
			   }else{
				  // fmt.Println(err)
				   fmt.Println(proxyIp["port"],url,"close")
			   }
			   */
    }
}


