package config

type Server struct {
	//JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	//Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	//Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	//Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	//Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Postgresql Postgresql `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	Sqlite3    Sqlite3    `mapstructure:"sqlite3" json:"sqlite3" yaml:"sqlite3"`
	Oracle     Oracle     `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Dameng     Dameng     `mapstructure:"dameng" json:"dameng" yaml:"dameng"`
	// oss
	//Local Local `mapstructure:"local" json:"local" yaml:"local"`
	//Qiniu Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
