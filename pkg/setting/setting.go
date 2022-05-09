package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config") //找寻文件的名字
	vp.SetConfigType("yaml")   // 找寻文件的类型
	for _, v := range configs {
		if v != "" {
			// vp.AddConfigPath("./configs") //从当前目录下的哪个文件夹找寻，
			vp.AddConfigPath(v)
		}
	}
	err := vp.ReadInConfig() //读取配置文件
	if err != nil {
		return nil, err
		// return nil, fmt.Errorf("%v", configs)
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
