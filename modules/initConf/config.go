package initConf

import (
	"github.com/Unknwon/goconfig"
)

func InitConf() (conf *goconfig.ConfigFile, err error) {

	conf, err = goconfig.LoadConfigFile("conf/app.conf")
	if err != nil {
		return nil, err
	}
	return conf, nil
}
