package config

type Server struct {
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite    Sqlite    `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Qiniu     Qiniu     `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Log       Log       `mapstructure:"log" json:"log" yaml:"log"`
	Mongo     Mongo     `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Dayu      Dayu      `mapstructure:"dayu" json:"dayu" yaml:"dayu"`
	XiaoETong XiaoETong `mapstructure:"xiaoETong" json:"xiaoETong" yaml:"xiaoetong"`
	WeiXin    WeiXin    `mapstructure:"wx" json:"wx" yaml:"wx"`
}

type XiaoETong struct {
	AppId        string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	GrantType    string `mapstructure:"grant_type" json:"grant_type" yaml:"grant_type"`
	ClientSecret string `mapstructure:"client_secret" json:"secret_key" yaml:"client_secret"`
	ClientId     string `mapstructure:"client_id" json:"client_id" yaml:"client_id"`
	URL          string `mapstructure:"url" json:"url" yaml:"url"`
}

type WeiXin struct {
	Url         string `mapstructure:"url" json:"url" yaml:"url"`
	SecretKey   string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	Appkey      string `mapstructure:"appkey" json:"appkey" yaml:"appkey"`
	RedirectUrl string `mapstructure:"redirectUrl" json:"redirectUrl" yaml:"redirectUrl"`
}

type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	EnabledMongo  bool   `mapstructure:"enabled-mongo" json:"enabledMongo" yaml:"enabled-mongo"`
	CertURL       string `mapstructure:"cert-url" json:"certURL" yaml:"cert-url"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}
type Qiniu struct {
	AccessKey string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath   string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

type Sqlite struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	LogMode  bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Mongo struct {
	DbName   string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
}

type Dayu struct {
	Appkey    string `mapstructure:"appkey" json:"appkey" yaml:"appkey"`
	SecretKey string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
	SignName  string `mapstructure:"signName" json:"signName" yaml:"signName"`
}
