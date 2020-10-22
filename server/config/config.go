package config

var (
	hasInit = false
)

//系统环境
type ModeType string

const (
	Debug   ModeType = "debug"   //开发环境
	Test    ModeType = "test"    //测试环境
	Release ModeType = "release" //线上环境
)

//数据库类型
type DBType int32

const (
	SysDB DBType = 0 //管理后台数据库
	SerDB DBType = 1 //业务数据库
)

//缓存类型
type CacheType int32

const (
	SysCache CacheType = 0 //管理后台缓存
	SerCache CacheType = 1 //业务缓存
)

type Server struct {
	Casbin   *Casbin    `mapstructure:"casbin" json:"casbin" `
	JWT      *JWT       `mapstructure:"jwt" json:"jwt" `
	Database *Database  `mapstructure:"database" json:"database" ` // 数据库配置项
	Qiniu    *Qiniu     `mapstructure:"qiniu" json:"qiniu" `
	Redis    []*Redis   `mapstructure:"redis" json:"redis" `
	System   *System    `mapstructure:"system" json:"system" `
	Captcha  *Captcha   `mapstructure:"captcha" json:"captcha" `
	Logger   *Logger    `mapstructure:"logger" json:"logger" `  // Log配置项
	Service  []*Service `mapstructure:"service" json:"service"` //代购服务器配置
	Gen      *Gen       `mapstructure:"gen" json:"gen"`
}

type Gen struct {
	FrontPath string `mapstructure:"front-path" json:"front-path" `
	RootPath  string `mapstructure:"root-path" json:"root-path" `
}

type Logger struct {
	Path       string `mapstructure:"path" json:"path" `
	Level      string `mapstructure:"level" json:"level" `
	Stdout     bool   `mapstructure:"stdout" json:"stdout" `
	EnabledBUS bool   `mapstructure:"enabledbus" json:"enabledbus" `
	EnabledREQ bool   `mapstructure:"enabledreq" json:"enabledreq" `
	EnabledDB  bool   `mapstructure:"enableddb" json:"enableddb" `
}

type Service struct {
	DB      string `mapstructure:"db" json:"db" `
	Name    string `mapstructure:"name" json:"name" `
	Redis   string `mapstructure:"redis" json:"redis" `
	BaseURL string `mapstructure:"base_url" json:"base_url" `
}

type System struct {
	UseMultipoint bool     `mapstructure:"use-multipoint" json:"useMultipoint" `
	Addr          string   `mapstructure:"addr" json:"addr" `
	Mode          ModeType `mapstructure:"mode" json:"mode" `
	Stdout        bool     `mapstructure:"stdout" json:"stdout" `
	Static        string   `mapstructure:"static" json:"static" `
	LogRoute      bool     `mapstructure:"logroute" json:"logroute" `
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" `
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" `
}

type DBConnect struct {
	URL  string `mapstructure:"url" json:"url" `
	Name string `mapstructure:"name" json:"name" `
	Type DBType `mapstructure:"type" json:"type" `
}

type Database struct {
	Driver        string       `mapstructure:"driver" json:"driver" `
	Stdout        bool         `mapstructure:"stdout" json:"stdout" `
	AutoMigrate   bool         `mapstructure:"auto-migrate" `
	SingularTable bool         `mapstructure:"singular-table" `
	Connects      []*DBConnect `mapstructure:"connect" json:"connect" `
}

type Redis struct {
	URL  string    `mapstructure:"url" json:"url" `
	Name string    `mapstructure:"name" json:"name" `
	Type CacheType `mapstructure:"type" json:"type" `
}

type Qiniu struct {
	AccessKey string `mapstructure:"access-key" json:"accessKey" `
	SecretKey string `mapstructure:"secret-key" json:"secretKey" `
	Bucket    string `mapstructure:"bucket" json:"bucket" `
	ImgPath   string `mapstructure:"img-path" json:"imgPath" `
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" `
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" `
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" `
}

type Sqlite struct {
	Username string `mapstructure:"username" json:"username" `
	Password string `mapstructure:"password" json:"password" `
	Path     string `mapstructure:"path" json:"path" `
	Config   string `mapstructure:"config" json:"config" `
	LogMode  bool   `mapstructure:"log-mode" json:"logMode" `
}
