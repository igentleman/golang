package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")    //找寻文件的名字
	vp.SetConfigType("yaml")      // 找寻文件的类型
	vp.AddConfigPath("./configs") //从当前目录下的哪个文件夹找寻，
	err := vp.ReadInConfig()      //读取配置文件
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
