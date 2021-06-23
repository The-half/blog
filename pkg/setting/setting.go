package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()

	//viper代码默认为config
	//vp.SetConfigFile("configs")

	//设置配置类型
	vp.SetConfigType("yaml")
	//添加配置文件路径
	vp.AddConfigPath("configs/")
	//fmt.Print(vp)
	//读取配置文件信息
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}