package config

type MysqlConfig struct{
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
	User string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
type ServerConfig struct{
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
}