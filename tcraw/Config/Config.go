package Config

import (
    goconfig "github.com/zarte/comutil/Goconfig"
    "fmt"
)

type config struct {
    Debug         bool
    Dsn           string
    RetryTimes    int
    DialTimeout   int
    ConnTimeout   int
    RedirectTimes int
    PJexepath     string
    RedisHost     string
    SeleniumPath  string
    Seleniumport  int
    Ptype  string
    Pcountry  string
    Panonymity string
    Path string
}

var ConfigValue = new(config)

func Load(configFile string) {
    config, err := goconfig.LoadConfigFile(configFile)
    if err != nil {
       fmt.Println(err)
       return
    }

    debug, _ := config.Bool(goconfig.DEFAULT_SECTION, "debug")
    ConfigValue.Debug = bool(debug)
    ConfigValue.Dsn, _ = config.GetValue(goconfig.DEFAULT_SECTION, "dsn")
    ConfigValue.RetryTimes, _ = config.Int(goconfig.DEFAULT_SECTION, "RetryTimes")
    ConfigValue.DialTimeout, _ = config.Int(goconfig.DEFAULT_SECTION, "DialTimeout")
    ConfigValue.ConnTimeout, _ = config.Int(goconfig.DEFAULT_SECTION, "ConnTimeout")
    ConfigValue.RedirectTimes, _ = config.Int(goconfig.DEFAULT_SECTION, "RedirectTimes")

    ConfigValue.PJexepath, _ = config.GetValue(goconfig.DEFAULT_SECTION, "PJexepath")
    ConfigValue.RedisHost, _ = config.GetValue(goconfig.DEFAULT_SECTION, "RedisHost")

    ConfigValue.SeleniumPath, _ = config.GetValue(goconfig.DEFAULT_SECTION, "SeleniumPath")
    ConfigValue.Seleniumport, _ = config.Int(goconfig.DEFAULT_SECTION, "Seleniumport")


    ConfigValue.Ptype, _ = config.GetValue(goconfig.DEFAULT_SECTION, "Ptype")
    ConfigValue.Pcountry, _ = config.GetValue(goconfig.DEFAULT_SECTION, "Pcountry")
    ConfigValue.Panonymity, _ = config.GetValue(goconfig.DEFAULT_SECTION, "Panonymity")

    if ConfigValue.RedirectTimes < 1 {
        ConfigValue.RedirectTimes = 5
    }

    if ConfigValue.DialTimeout < 1 {
        ConfigValue.DialTimeout = 2
    }

    if ConfigValue.ConnTimeout < 1 {
        ConfigValue.ConnTimeout = 2
    }

    if ConfigValue.RedirectTimes < 1 {
        ConfigValue.RedirectTimes = 5
    }

}
