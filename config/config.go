package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
type ConsulCOnfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
type ServerConfig struct {
	//Host string `mapstructure:"host" json:"host"`
	//Port int `mapstructure:"port" json:"port"`
	//最后都会放到配置中心nacos去配置，这里就不配置这么多了
	Name       string          `mapstructure:"name" json:"name"` //给consul用的，服务发现通过这个name来拿到服务
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulCOnfig `mapstructure:"consul" json:"consul"`
}