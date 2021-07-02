package main

import (
    "agentsearch/tcraw/Config"
    "agentsearch/tcraw/Fileop"
    "agentsearch/tcraw/Logger"
    "agentsearch/tcraw/Worker"
    "log"
    _ "net/http/pprof"
    "os"
    "time"
)

var logfile *os.File

func initLog() {
    if logfile != nil {
        logfile.Close()
        logfile = nil
    }

    path := os.Args[1] + "logs/" + time.Now().Format("2006/01/")
    os.MkdirAll(path, 0666)
    file := os.Args[1] + "logs/" + time.Now().Format("2006/01/02")
    var err error
    logfile, err = os.OpenFile(file + ".log", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
    if err != nil {
        log.Printf("%s\r\n", err.Error())
        Logger.Logger = nil
        logfile = nil
    } else {
        Logger.Logger = log.New(logfile, "\r\n", log.Ltime)
		Fileop.Fileoper = Fileop.NewLog(os.Args[1])
    }
}

//运行 go run main.go ini文件绝对路径
func main() {
    defer func() {
        if logfile != nil {
            logfile.Close()
        }
    }()

    Config.Load(os.Args[1] + "config.ini")
    initLog()
    Config.ConfigValue.Path = os.Args[1]
    for {
        //每次运行初始化
        startTime := time.Now().Unix()
        Worker.ThreadRun()
        useTime := int(time.Now().Unix() - startTime)
        Logger.ToFile("end used time[", useTime, "]")
        os.Exit(0)
    }
}


