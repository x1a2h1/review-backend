package config

type Config struct {
	Server struct {
		Port  int  `mapstructure:"port"`
		Debug bool `mapstructure:"debug"`
	} `mapstructure:"server"`
	SMS struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		From     string `mapstructure:"from"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		UseSSL   bool   `mapstructure:"useSSL"`
	} `mapstructure:"sms"`
	Mysql struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		ConnPool struct {
			MaxOpenConns    int `mapstructure:"maxOpenConns"`
			MaxIdleConns    int `mapstructure:"maxIdleConns"`
			ConnMaxLifetime int `mapstructure:"connMaxLifetime"` // 单位：秒
			ConnMaxIdleTime int `mapstructure:"connMaxIdleTime"` // 单位：秒
		} `mapstructure:"connPool"`
	} `mapstructure:"mysql"`
	Storage struct {
		Platform string   `mapstructure:"platform"`
		USS      struct{} `mapstructure:"uss"`
		KODO     struct {
			AccessKey string `mapstructure:"access_key"`
			SecretKey string `mapstructure:"secret_key"`
			Bucket    string `mapstructure:"bucket"`
			Server    string `mapstructure:"server"`
		} `mapstructure:"kodo"`
	} `mapstructure:"storage"`
}
