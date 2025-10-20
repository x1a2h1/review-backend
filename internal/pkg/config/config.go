package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

// using viper read or write config file
var cfg *viper.Viper

var Cfg *Config

func init() {
	cfg = viper.New()
	cfg.AddConfigPath("./")
	cfg.SetConfigName("config")
	cfg.SetConfigType("toml")
	if err := cfg.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件失败，err:%v \n", err)
		if !os.IsNotExist(err) {
			if file, err := os.OpenFile("./config.toml", os.O_CREATE|os.O_WRONLY, 0644); err != nil {
				fmt.Printf("创建配置文件失败，err:%v \n", err)
			} else {
				file.Close()
			}
		}
	}
	// 开始合并
	merge()
	read()

	cfg.WatchConfig()
	cfg.OnConfigChange(func(in fsnotify.Event) {
		read()
	})
}

func merge() {
	out := make(map[string]any)
	config := &mapstructure.DecoderConfig{
		TagName: "mapstructure",
		Result:  &out}
	dec, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Println("创建编码器失败", err)
		return
	}
	if err := dec.Decode(&Config{}); err != nil {
		fmt.Println("编码文件失败", err)
		return
	}
	for key, value := range out {
		cfg.SetDefault(key, value)
	}
	if err := cfg.WriteConfig(); err != nil {
		fmt.Println("写入配置文件失败", err)
	}
}
func read() {
	content := &Config{}
	err := cfg.Unmarshal(&content)
	if err != nil {
		fmt.Println("解析配置文件失败", err)
		return
	}
	Cfg = content
}
func Set(key string, value any) {
	cfg.Set(key, value)
	if err := cfg.SafeWriteConfig(); err != nil {
		fmt.Println("写入配置文件失败", err)
	}
}
