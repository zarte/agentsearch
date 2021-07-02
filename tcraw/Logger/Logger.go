package Logger

import (
    "log"
    "agentsearch/tcraw/Config"
)

var Logger *log.Logger

func Debug(v ...interface{}) {
    if Config.ConfigValue.Debug {
        log.Println(v...)
    }
}

func ToFile(v ...interface{}) {
    if Logger != nil {
        Logger.Println(v...)
        if Config.ConfigValue.Debug {
            log.Println(v...)
        }
    } else {
        log.Println(v...)
    }
}